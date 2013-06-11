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
type wktDecoder struct {
	r *C.GEOSWKTReader
}

// Creates a new WKT decoder, can be nil if initialization in the C API fails
func newWktDecoder() *wktDecoder {
	r := C.GEOSWKTReader_create_r(handle)
	if r == nil {
		return nil
	}
	d := &wktDecoder{r}
	runtime.SetFinalizer(d, (*wktDecoder).destroy)
	return d
}

// decode decodes the WKT string and returns a geometry
func (d *wktDecoder) decode(wkt string) (*Geometry, error) {
	cstr := C.CString(wkt)
	defer C.free(unsafe.Pointer(cstr))
	g := C.GEOSWKTReader_read_r(handle, d.r, cstr)
	if g == nil {
		return nil, Error()
	}
	return geomFromPtr(g), nil
}

func (d *wktDecoder) destroy() {
	// XXX: mutex
	C.GEOSWKTReader_destroy_r(handle, d.r)
	d.r = nil
}

type wktEncoder struct {
	w *C.GEOSWKTWriter
}

func newWktEncoder() *wktEncoder {
	w := C.GEOSWKTWriter_create_r(handle)
	if w == nil {
		return nil
	}
	e := &wktEncoder{w}
	runtime.SetFinalizer(e, (*wktEncoder).destroy)
	return e
}

// Encode returns a string that is the geometry encoded as WKT
func (e *wktEncoder) encode(g *Geometry) (string, error) {
	cstr := C.GEOSWKTWriter_write_r(handle, e.w, g.g)
	if cstr == nil {
		return "", Error()
	}
	return C.GoString(cstr), nil
}

func (e *wktEncoder) destroy() {
	// XXX: mutex
	C.GEOSWKTWriter_destroy_r(handle, e.w)
	e.w = nil
}
