package geos

/*
#include "geos.h"
*/
import "C"

import (
	"runtime"
)

type CoordSeq struct {
	c *C.GEOSCoordSequence
}

func NewCoordSeq(size, dims int) *CoordSeq {
	p := C.GEOSCoordSeq_create_r(handle, C.uint(size), C.uint(dims))
	if p == nil {
		return nil
	}
	return coordSeqFromPtr(p)
}

func coordSeqFromPtr(c *C.GEOSCoordSequence) *CoordSeq {
	cs := &CoordSeq{c}
	runtime.SetFinalizer(cs, (*CoordSeq).destroy)
	return cs
}

func (c *CoordSeq) Clone() (*CoordSeq, error) {
	p := C.GEOSCoordSeq_clone_r(handle, c.c)
	if p == nil {
		return nil, Error()
	}
	return coordSeqFromPtr(p), nil
}

func (c *CoordSeq) setX(idx int, val float64) error {
	i := C.GEOSCoordSeq_setX_r(handle, c.c, C.uint(idx), C.double(val))
	if i == 0 {
		return Error()
	}
	return nil
}

func (c *CoordSeq) setY(idx int, val float64) error {
	i := C.GEOSCoordSeq_setY_r(handle, c.c, C.uint(idx), C.double(val))
	if i == 0 {
		return Error()
	}
	return nil
}

func (c *CoordSeq) setZ(idx int, val float64) error {
	i := C.GEOSCoordSeq_setZ_r(handle, c.c, C.uint(idx), C.double(val))
	if i == 0 {
		return Error()
	}
	return nil
}

func (c *CoordSeq) GetX(idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getX_r(handle, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *CoordSeq) GetY(idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getY_r(handle, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *CoordSeq) GetZ(idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getZ_r(handle, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *CoordSeq) Size() int {
	var val C.uint
	i := C.GEOSCoordSeq_getSize_r(handle, c.c, &val)
	if i == 0 { // C API exception
		return 0
	}
	return int(val)
}

func (c *CoordSeq) Dims() int {
	var val C.uint
	i := C.GEOSCoordSeq_getDimensions_r(handle, c.c, &val)
	if i == 0 { // C API exception
		return 0
	}
	return int(val)
}

func (c *CoordSeq) destroy() {
	// XXX: mutex
	C.GEOSCoordSeq_destroy_r(handle, c.c)
	c.c = nil
}
