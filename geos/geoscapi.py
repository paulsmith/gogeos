import sys
from string import Template
from collections import namedtuple
from pycparser import c_parser, c_ast, parse_file

Func = namedtuple('Func', ('name', 'type', 'args'))
Arg = namedtuple('Arg', ('name', 'type'))
Type = namedtuple('Type', ('ptr', 'name', 'array'))

class FuncDeclVisitor(c_ast.NodeVisitor):
    def __init__(self):
        self.funcs = []
        self.reset()

    def reset(self):
        self.name = None
        self.ptr = ''
        self.type = None
        self.inargs = False
        self.args = []
        self.argname = None
        self.array = False

    def visit_Typedef(self, node):
        # Prevent func decls in typedefs from being visited
        pass
        
    def visit_FuncDecl(self, node):
        self.visit(node.type)
        if node.args:
            self.inargs = True
            self.visit(node.args)
        self.funcs.append(Func(self.name, self.type, self.args))
        self.reset()

    def visit_PtrDecl(self, node):
        self.ptr += '*'
        self.visit(node.type)

    def visit_TypeDecl(self, node):
        if node.type.__class__.__name__ == 'Struct':
            return
        if self.inargs:
            self.argname = node.declname
        else:
            self.name = node.declname
        self.visit(node.type)

    def visit_ArrayDecl(self, node):
        self.array = True
        self.visit(node.type)

    def visit_IdentifierType(self, node):
        type_ = Type(self.ptr, ' '.join(node.names), self.array)
        if self.inargs:
            self.args.append(Arg(self.argname, type_))
        else:
            self.type = type_
        self.ptr = ''
        self.array = False

def cgo_func_wrappers(filename):
    ast = parse_file(filename, use_cpp=True)
    v = FuncDeclVisitor()
    v.visit(ast)

    funcnames = {}
    threadsafe = []

    for func in v.funcs:
        funcnames[func.name] = func

    for func in v.funcs:
        if not func.name.endswith('_r'):
            if func.name + '_r' in funcnames:
                threadsafe.append(funcnames[func.name + '_r'])
            else:
                threadsafe.append(func)

    print("""
package geos

// Created mechanically from C API header - DO NOT EDIT

/*
#include <geos_c.h>
*/
import "C"

import (
    "unsafe"
)\
""")

    typemap = {
        "unsigned char": "uchar",
        "unsigned int": "uint",
    }

    identmap = {
        "type": "_type",
    }

    for func in threadsafe:
        def gotype(ctype):
            type_ = "C." + typemap.get(ctype.name, ctype.name)
            if ctype.ptr:
                type_ = ctype.ptr + type_
            if ctype.array:
                type_ = '[]' + type_
            return type_
        
        def goident(arg, inbody=True):
            def voidptr(ctype):
                return ctype.ptr and ctype.name == 'void'

            ident = identmap.get(arg.name, arg.name)
            if arg.type.array and inbody:
                ident = '&' + ident + '[0]'
            if voidptr(arg.type) and inbody:
                ident = 'unsafe.Pointer(' + ident + ')'
            return ident

        # Go function signature
        gosig = "func $name($parameters)"
        if func.type.name != "void":
            gosig += " $result"
        gosig += " {"
        t = Template(gosig)
        params = ", ".join([goident(p, inbody=False) + " " + gotype(p.type) for p in func.args if p.type.name != 'GEOSContextHandle_t'])
        result = gotype(func.type)
        func_name = "c" + func.name
        if func_name.endswith('_r'):
            func_name = func_name[:-2]
        print(t.substitute(name=func_name, parameters=params, result=result))

        # Go function body
        gobody = """\
\t${return_stmt}C.$name($args)
}
"""
        if func.name.endswith("_r") and func.name != "initGEOS_r":
            gobody = """\
\t${handle_lock}.Lock()
\tdefer ${handle_lock}.Unlock()
""" + gobody

        t = Template(gobody)
        args = ", ".join([goident(p) for p in func.args])
        return_stmt = 'return ' if func.type.name != 'void' else ''
        print(t.substitute(return_stmt=return_stmt, name=func.name, args=args, handle_lock='handlemu'))

if __name__ == "__main__":
    cgo_func_wrappers(sys.argv[1])
    #from pycparser.c_generator import CGenerator
    #ast = parse_file(sys.argv[1], use_cpp=True)
    #print(CGenerator().visit(ast))
