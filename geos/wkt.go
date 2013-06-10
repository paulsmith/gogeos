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
type WKTDecoder struct {
	r *C.GEOSWKTReader
}

// Creates a new WKT decoder, can be nil if initialization in the C API fails
func NewWKTDecoder() *WKTDecoder {
	r := C.GEOSWKTReader_create_r(handle)
	if r == nil {
		return nil
	}
	d := &WKTDecoder{r}
	runtime.SetFinalizer(d, (*WKTDecoder).destroy)
	return d
}

// Decode decodes the WKT string and returns a geometry
func (d *WKTDecoder) Decode(wkt string) (*Geometry, error) {
	cstr := C.CString(wkt)
	defer C.free(unsafe.Pointer(cstr))
	g := C.GEOSWKTReader_read_r(handle, d.r, cstr)
	if g == nil {
		return nil, Error()
	}
	// XXX: GeomFromPtr
	return &Geometry{g}, nil
}

func (d *WKTDecoder) destroy() {
	// XXX: mutex
	C.GEOSWKTReader_destroy_r(handle, d.r)
	d.r = nil
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

// Encode returns a string that is the geometry encoded as WKT
func (w *WKTWriter) Encode(g *Geometry) (string, error) {
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
