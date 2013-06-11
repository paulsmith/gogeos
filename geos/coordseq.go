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

func coordSeqFromSlice(coords []Coord) (*CoordSeq, error) {
	// XXX: handle 3-dim
	cs := NewCoordSeq(len(coords), 2)
	for i, c := range coords {
		if err := cs.setX(i, c.X); err != nil {
			return nil, err
		}
		if err := cs.setY(i, c.Y); err != nil {
			return nil, err
		}
	}
	return cs, nil
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

func (c *CoordSeq) x(idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getX_r(handle, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *CoordSeq) y(idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getY_r(handle, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *CoordSeq) z(idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getZ_r(handle, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *CoordSeq) size() (int, error) {
	var val C.uint
	i := C.GEOSCoordSeq_getSize_r(handle, c.c, &val)
	if i == 0 {
		return 0, Error()
	}
	return int(val), nil
}

func (c *CoordSeq) dims() (int, error) {
	var val C.uint
	i := C.GEOSCoordSeq_getDimensions_r(handle, c.c, &val)
	if i == 0 {
		return 0, Error()
	}
	return int(val), nil
}

func (c *CoordSeq) destroy() {
	// XXX: mutex
	C.GEOSCoordSeq_destroy_r(handle, c.c)
	c.c = nil
}
