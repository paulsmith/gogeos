package geos

/*
#include "geos.h"
*/
import "C"

import (
	"errors"
	"runtime"
)

type PGeometry struct {
	p *C.GEOSPreparedGeometry
}

func PrepareGeometry(g *Geometry) *PGeometry {
	ptr := C.GEOSPrepare_r(handle, g.g)
	p := &PGeometry{ptr}
	runtime.SetFinalizer(p, (*PGeometry).destroy)
	return p
}

func (p *PGeometry) destroy() {
	C.GEOSPreparedGeom_destroy_r(handle, p.p)
	p.p = nil
}

// Prepared geometry binary predicates

func (p *PGeometry) Contains(other *PGeometry) (bool, error) {
	return p.predicate("contains", cpreparedcontains, other)
}

func (p *PGeometry) ContainsP(other *PGeometry) (bool, error) {
	return p.predicate("contains", cpreparedcontainsproperly, other)
}

func (p *PGeometry) CoveredBy(other *PGeometry) (bool, error) {
	return p.predicate("covered by", cpreparedcoveredby, other)
}

func (p *PGeometry) Covers(other *PGeometry) (bool, error) {
	return p.predicate("covers", cpreparedcovers, other)
}

func (p *PGeometry) Crosses(other *PGeometry) (bool, error) {
	return p.predicate("crosses", cpreparedcrosses, other)
}

func (p *PGeometry) Disjoint(other *PGeometry) (bool, error) {
	return p.predicate("disjoint", cprepareddisjoint, other)
}

func (p *PGeometry) Intersects(other *PGeometry) (bool, error) {
	return p.predicate("intersects", cpreparedintersects, other)
}

func (p *PGeometry) Overlaps(other *PGeometry) (bool, error) {
	return p.predicate("overlaps", cpreparedoverlaps, other)
}

func (p *PGeometry) Touches(other *PGeometry) (bool, error) {
	return p.predicate("touches", cpreparedtouches, other)
}

func (p *PGeometry) Within(other *PGeometry) (bool, error) {
	return p.predicate("within", cpreparedwithin, other)
}

func (p *PGeometry) predicate(name string, fn func(C.GEOSContextHandle_t, *C.GEOSPreparedGeometry, *C.GEOSPreparedGeometry) C.char, other *PGeometry) (bool, error) {
	i := fn(handle, p.p, other.p)
	if i == 2 {
		return false, errors.New("geos: prepared " + name)
	}
	return i == 1, nil
}
