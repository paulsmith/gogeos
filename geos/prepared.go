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
	ptr := cGEOSPrepare(g.g)
	p := &PGeometry{ptr}
	runtime.SetFinalizer(p, (*PGeometry).destroy)
	return p
}

func (p *PGeometry) destroy() {
	cGEOSPreparedGeom_destroy(p.p)
	p.p = nil
}

// Prepared geometry binary predicates

func (p *PGeometry) Contains(other *Geometry) (bool, error) {
	return p.predicate("contains", cGEOSPreparedContains, other)
}

func (p *PGeometry) ContainsP(other *Geometry) (bool, error) {
	return p.predicate("contains", cGEOSPreparedContainsProperly, other)
}

func (p *PGeometry) CoveredBy(other *Geometry) (bool, error) {
	return p.predicate("covered by", cGEOSPreparedCoveredBy, other)
}

func (p *PGeometry) Covers(other *Geometry) (bool, error) {
	return p.predicate("covers", cGEOSPreparedCovers, other)
}

func (p *PGeometry) Crosses(other *Geometry) (bool, error) {
	return p.predicate("crosses", cGEOSPreparedCrosses, other)
}

func (p *PGeometry) Disjoint(other *Geometry) (bool, error) {
	return p.predicate("disjoint", cGEOSPreparedDisjoint, other)
}

func (p *PGeometry) Intersects(other *Geometry) (bool, error) {
	return p.predicate("intersects", cGEOSPreparedIntersects, other)
}

func (p *PGeometry) Overlaps(other *Geometry) (bool, error) {
	return p.predicate("overlaps", cGEOSPreparedOverlaps, other)
}

func (p *PGeometry) Touches(other *Geometry) (bool, error) {
	return p.predicate("touches", cGEOSPreparedTouches, other)
}

func (p *PGeometry) Within(other *Geometry) (bool, error) {
	return p.predicate("within", cGEOSPreparedWithin, other)
}

func (p *PGeometry) predicate(name string, fn func(*C.GEOSPreparedGeometry, *C.GEOSGeometry) C.char, other *Geometry) (bool, error) {
	i := fn(p.p, other.g)
	if i == 2 {
		return false, errors.New("geos: prepared " + name)
	}
	return i == 1, nil
}
