package geos

/*
#include "geos.h"
*/
import "C"

import (
	"errors"
	"runtime"
)

// PGeometry represents a "prepared geometry", a type of geometry object that is
// optimized for a limited set of operations.
type PGeometry struct {
	p *C.GEOSPreparedGeometry
}

// PrepareGeometry constructs a prepared geometry from a normal geometry object.
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

// Contains computes whether the prepared geometry contains the other prepared
// geometry.
func (p *PGeometry) Contains(other *Geometry) (bool, error) {
	return p.predicate("contains", cGEOSPreparedContains, other)
}

// ContainsP computes whether the prepared geometry properly contains the other
// prepared geometry.
func (p *PGeometry) ContainsP(other *Geometry) (bool, error) {
	return p.predicate("contains", cGEOSPreparedContainsProperly, other)
}

// CoveredBy computes whether the prepared geometry is covered by the other
// prepared geometry.
func (p *PGeometry) CoveredBy(other *Geometry) (bool, error) {
	return p.predicate("covered by", cGEOSPreparedCoveredBy, other)
}

// Covers computes whether the prepared geometry covers the other prepared
// geometry.
func (p *PGeometry) Covers(other *Geometry) (bool, error) {
	return p.predicate("covers", cGEOSPreparedCovers, other)
}

// Crosses computes whether the prepared geometry crosses the other prepared
// geometry.
func (p *PGeometry) Crosses(other *Geometry) (bool, error) {
	return p.predicate("crosses", cGEOSPreparedCrosses, other)
}

// Disjoint computes whether the prepared geometry is disjoint from the other
// prepared geometry.
func (p *PGeometry) Disjoint(other *Geometry) (bool, error) {
	return p.predicate("disjoint", cGEOSPreparedDisjoint, other)
}

// Intersects computes whether the prepared geometry intersects the other
// prepared geometry.
func (p *PGeometry) Intersects(other *Geometry) (bool, error) {
	return p.predicate("intersects", cGEOSPreparedIntersects, other)
}

// Overlaps computes whether the prepared geometry overlaps the other
// prepared geometry.
func (p *PGeometry) Overlaps(other *Geometry) (bool, error) {
	return p.predicate("overlaps", cGEOSPreparedOverlaps, other)
}

// Touches computes whether the prepared geometry touches the other
// prepared geometry.
func (p *PGeometry) Touches(other *Geometry) (bool, error) {
	return p.predicate("touches", cGEOSPreparedTouches, other)
}

// Within computes whether the prepared geometry is within the other
// prepared geometry.
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
