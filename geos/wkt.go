package geos

/*
#include "geos.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// Reads the WKT serialization and produces geometries
type WKTReader struct {
	r *C.GEOSWKTReader
}

// Creates a new WKT reader, can be nil if initialization in the C API fails
func NewWKTReader() *WKTReader {
	r := C.GEOSWKTReader_create_r(handle)
	if r == nil {
		return nil
	}
	reader := &WKTReader{r}
	runtime.SetFinalizer(reader, (*WKTReader).destroy)
	return reader
}

// Decode decodes the WKT string and returns a geometry
func (r *WKTReader) Decode(wkt string) (*Geometry, error) {
	cstr := C.CString(wkt)
	defer C.free(unsafe.Pointer(cstr))
	g := C.GEOSWKTReader_read_r(handle, r.r, cstr)
	if g == nil {
		return nil, Error()
	}
	// XXX: GeomFromPtr
	return &Geometry{g}, nil
}

func (r *WKTReader) destroy() {
	// XXX: mutex
	C.GEOSWKTReader_destroy_r(handle, r.r)
	r.r = nil
}

type WKTWriter struct {
	w *C.GEOSWKTWriter
}

func NewWKTWriter() *WKTWriter {
	w := C.GEOSWKTWriter_create_r(handle)
	if w == nil {
		return nil
	}
	writer := &WKTWriter{w}
	runtime.SetFinalizer(writer, (*WKTWriter).destroy)
	return writer
}

var DefaultWKTWriter *WKTWriter

func (w *WKTWriter) Write(g *Geometry) (string, error) {
	// XXX: free?
	cstr := C.GEOSWKTWriter_write_r(handle, w.w, g.g)
	if cstr == nil {
		return "", Error()
	}
	return C.GoString(cstr), nil
}

func (w *WKTWriter) destroy() {
	// XXX: mutex
	C.GEOSWKTWriter_destroy_r(handle, w.w)
	w.w = nil
}

func init() {
	DefaultWKTWriter = NewWKTWriter()
}
