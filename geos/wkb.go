package geos

/*
#include "geos.h"
*/
import "C"

import (
	"encoding/hex"
	"errors"
	"runtime"
	"unsafe"
)

var (
	ErrWKBRead  = errors.New("geos: reading WKB")
	ErrWKBWrite = errors.New("geos: writing WKB")
)

type WKBDecoder struct {
	r *C.GEOSWKBReader
}

func NewWKBDecoder() *WKBDecoder {
	r := cGEOSWKBReader_create_r(handle)
	d := &WKBDecoder{r}
	runtime.SetFinalizer(d, (*WKBDecoder).destroy)
	return d
}

func (d *WKBDecoder) destroy() {
	// XXX: mutex
	cGEOSWKBReader_destroy_r(handle, d.r)
	d.r = nil
}

func (d *WKBDecoder) Decode(wkb []byte) (*Geometry, error) {
	var cwkb []C.uchar
	for i := range wkb {
		cwkb = append(cwkb, C.uchar(wkb[i]))
	}
	g := cGEOSWKBReader_read_r(handle, d.r, &cwkb[0], C.size_t(len(wkb)))
	if g == nil {
		return nil, ErrWKBRead
	}
	return GeomFromPtr(g), nil
}

func (d *WKBDecoder) DecodeHex(wkbHex string) (*Geometry, error) {
	wkb, err := hex.DecodeString(wkbHex)
	if err != nil {
		return nil, err
	}
	return d.Decode(wkb)
}

type WKBWriter struct {
	w *C.GEOSWKBWriter
}

func NewWKBWriter() *WKBWriter {
	w := C.GEOSWKBWriter_create_r(handle)
	if w == nil {
		return nil
	}
	writer := &WKBWriter{w}
	runtime.SetFinalizer(writer, (*WKBWriter).destroy)
	return writer
}

var DefaultWKBWriter *WKBWriter

func writeWkb(w *WKBWriter, g *Geometry, fn func(C.GEOSContextHandle_t, *C.GEOSWKBWriter, *C.GEOSGeometry, *C.size_t) *C.uchar) ([]byte, error) {
	var size C.size_t
	bytes := fn(handle, w.w, g.g, &size)
	if bytes == nil {
		return nil, ErrWKBWrite
	}
	ptr := unsafe.Pointer(bytes)
	defer C.free(ptr)
	l := int(size)
	var out []byte
	for i := 0; i < l; i++ {
		el := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(C.uchar(0))*uintptr(i))
		out = append(out, byte(*(*C.uchar)(el)))
	}
	return out, nil
}

func (w *WKBWriter) Write(g *Geometry) ([]byte, error) {
	return writeWkb(w, g, cGEOSWKBWriter_write_r)
}

func (w *WKBWriter) WriteHex(g *Geometry) ([]byte, error) {
	return writeWkb(w, g, cGEOSWKBWriter_writeHEX_r)
}

func (w *WKBWriter) destroy() {
	// XXX: mutex
	C.GEOSWKBWriter_destroy_r(handle, w.w)
	w.w = nil
}

func init() {
	DefaultWKBWriter = NewWKBWriter()
}
