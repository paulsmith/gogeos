package geos

/*
#include "geos.h"
*/
import "C"

import (
	"runtime"
)

type Geometry struct {
	g *C.GEOSGeometry
}

func geomFromPtr(p *C.GEOSGeometry) *Geometry {
	g := &Geometry{p}
	runtime.SetFinalizer(g, (*Geometry).destroy)
	return g
}

// FromWKT is a factory function that returns a geometry decoded from a
// Well-Known Text (WKT) string
func FromWKT(wkt string) (*Geometry, error) {
	decoder := NewWKTDecoder()
	return decoder.Decode(wkt)
}

func (g *Geometry) destroy() {
	C.GEOSGeom_destroy_r(handle, g.g)
	g.g = nil
}

func (g *Geometry) ToWKT() (string, error) {
	encoder := NewWKTEncoder()
	return encoder.Encode(g)
}

func (g *Geometry) String() string {
	str, err := g.ToWKT()
	if err != nil {
		return "" // XXX: better to panic?
	}
	return str
}

// Linearref functions

// Return distance of point 'p' projected on this geometry from origin.
// This must be a lineal geometry */
func (g *Geometry) Project(p *Geometry) float64 {
	// XXX: error if wrong geometry types
	return float64(C.GEOSProject_r(handle, g.g, p.g))
}

// Return closest point to given distance within geometry
// This geometry must be a LineString */
func (g *Geometry) Interpolate(dist float64) (*Geometry, error) {
	p := C.GEOSInterpolate_r(handle, g.g, C.double(dist))
	// XXX: test for exception
	return geomFromPtr(p), nil
}

// Buffer functions
// XXX: buffer w number of segments, endcap, join, mitre limit

func (g *Geometry) Buffer(d float64) (*Geometry, error) {
	const quadsegs = 8
	return geomFromC("Buffer", cGEOSBuffer_r(handle, g.g, C.double(d), quadsegs))
}

// Geometry Constructors

func NewPoint(coords ...Coord) (*Geometry, error) {
	if len(coords) == 0 {
		return emptyGeom("EmptyPoint", cGEOSGeom_createEmptyPoint_r)
	}
	cs, err := coordSeqFromSlice(coords)
	if err != nil {
		return nil, err
	}
	return geomFromCoordSeq(cs, "NewPoint", cGEOSGeom_createPoint_r)
}

func NewLinearRing(coords ...Coord) (*Geometry, error) {
	cs, err := coordSeqFromSlice(coords)
	if err != nil {
		return nil, err
	}
	return geomFromCoordSeq(cs, "NewLinearRing", cGEOSGeom_createLinearRing_r)
}

func NewLineString(coords ...Coord) (*Geometry, error) {
	cs, err := coordSeqFromSlice(coords)
	if err != nil {
		return nil, err
	}
	return geomFromCoordSeq(cs, "NewLineString", cGEOSGeom_createLineString_r)
}

func EmptyPolygon() (*Geometry, error) {
	return emptyGeom("EmptyPoint", cGEOSGeom_createEmptyPolygon_r)
}

func NewPolygon(shell []Coord, holes ...[]Coord) (*Geometry, error) {
	ext, err := NewLinearRing(shell...)
	if err != nil {
		return nil, err
	}
	var ints []*Geometry
	for i := range holes {
		g, err := NewLinearRing(holes[i]...)
		if err != nil {
			return nil, err
		}
		ints = append(ints, g)
	}
	return PolygonFromGeom(ext, ints...)
}

func PolygonFromGeom(shell *Geometry, holes ...*Geometry) (*Geometry, error) {
	var ptrHoles **C.GEOSGeometry
	// build c array of geom ptrs
	var holeCPtrs []*C.GEOSGeometry
	for i := range holes {
		holeCPtrs = append(holeCPtrs, holes[i].g)
	}
	if len(holeCPtrs) > 0 {
		ptrHoles = &holeCPtrs[0]
	}
	return geomFromC("NewPolygon", cGEOSGeom_createPolygon_r(handle, shell.g, ptrHoles, C.uint(len(holeCPtrs))))
}

func NewCollection(_type GeometryType, geoms ...*Geometry) (*Geometry, error) {
	if len(geoms) == 0 {
		return geomFromC("EmptyCollection", cGEOSGeom_createEmptyCollection_r(handle, C.int(_type)))
	}
	var ptrGeoms **C.GEOSGeometry
	// build c array of geom ptrs
	var geomCPtrs []*C.GEOSGeometry
	for i := range geoms {
		geomCPtrs = append(geomCPtrs, geoms[i].g)
	}
	ptrGeoms = &geomCPtrs[0]
	return geomFromC("NewCollection", cGEOSGeom_createCollection_r(handle, C.int(_type), ptrGeoms, C.uint(len(geomCPtrs))))
}

func (g *Geometry) Clone() (*Geometry, error) {
	return g.unaryTopo("Clone", cGEOSGeom_clone_r)
}

// Unary topology functions

func (g *Geometry) Envelope() (*Geometry, error) {
	return g.unaryTopo("Envelope", cGEOSEnvelope_r)
}

func (g *Geometry) ConvexHull() (*Geometry, error) {
	return g.unaryTopo("ConvexHull", cGEOSConvexHull_r)
}

func (g *Geometry) Boundary() (*Geometry, error) {
	return g.unaryTopo("Boundary", cGEOSBoundary_r)
}

func (g *Geometry) UnaryUnion() (*Geometry, error) {
	return g.unaryTopo("UnaryUnion", cGEOSUnaryUnion_r)
}

func (g *Geometry) PointOnSurface() (*Geometry, error) {
	return g.unaryTopo("PointOnSurface", cGEOSPointOnSurface_r)
}

func (g *Geometry) Centroid() (*Geometry, error) {
	return g.unaryTopo("Centroid", cGEOSGetCentroid_r)
}

func (g *Geometry) LineMerge() (*Geometry, error) {
	return g.unaryTopo("LineMerge", cGEOSLineMerge_r)
}

// Returns a geometry simplified by amount given by tolerance.
// May not preserve topology -- see SimplifyP().
func (g *Geometry) Simplify(tolerance float64) (*Geometry, error) {
	return g.simplify("simplify", cGEOSSimplify_r, tolerance)
}

// Returns a geometry simplified by amount given by tolerance.
// Unlike Simplify(), SimplifyP() guarantees it will preserve topology.
func (g *Geometry) SimplifyP(tolerance float64) (*Geometry, error) {
	return g.simplify("simplify", cGEOSTopologyPreserveSimplify_r, tolerance)
}

// Return all distinct vertices of input geometry as a MULTIPOINT.
func (g *Geometry) UniquePoints() (*Geometry, error) {
	return g.unaryTopo("UniquePoints", cGEOSGeom_extractUniquePoints_r)
}

// Find paths shared between the two given lineal geometries.
// Returns a GEOMETRYCOLLECTION having two elements:
// - first element is a MULTILINESTRING containing shared paths
//   having the _same_ direction on both inputs
// - second element is a MULTILINESTRING containing shared paths
//   having the _opposite_ direction on the two inputs
func (g *Geometry) SharedPaths(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("SharedPaths", cGEOSSharedPaths_r, other)
}

// Snap first geometry on to second with given tolerance.
func (g *Geometry) Snap(other *Geometry, tolerance float64) (*Geometry, error) {
	return geomFromC("Snap", cGEOSSnap_r(handle, g.g, other.g, C.double(tolerance)))
}

func (g *Geometry) Prepare() *PGeometry {
	return PrepareGeometry(g)
}

// Binary topology functions

func (g *Geometry) Intersection(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("Intersection", cGEOSIntersection_r, other)
}

func (g *Geometry) Difference(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("Difference", cGEOSDifference_r, other)
}

func (g *Geometry) SymDifference(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("SymDifference", cGEOSSymDifference_r, other)
}

func (g *Geometry) Union(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("Union", cGEOSUnion_r, other)
}

// Binary predicate functions

func (g *Geometry) Disjoint(other *Geometry) (bool, error) {
	return g.binaryPred("Disjoint", cGEOSDisjoint_r, other)
}

func (g *Geometry) Touches(other *Geometry) (bool, error) {
	return g.binaryPred("Touches", cGEOSTouches_r, other)
}

func (g *Geometry) Intersects(other *Geometry) (bool, error) {
	return g.binaryPred("Intersects", cGEOSIntersects_r, other)
}

func (g *Geometry) Crosses(other *Geometry) (bool, error) {
	return g.binaryPred("Crosses", cGEOSCrosses_r, other)
}

func (g *Geometry) Within(other *Geometry) (bool, error) {
	return g.binaryPred("Within", cGEOSWithin_r, other)
}

func (g *Geometry) Contains(other *Geometry) (bool, error) {
	return g.binaryPred("Contains", cGEOSContains_r, other)
}

func (g *Geometry) Overlaps(other *Geometry) (bool, error) {
	return g.binaryPred("Overlaps", cGEOSOverlaps_r, other)
}

func (g *Geometry) Equals(other *Geometry) (bool, error) {
	return g.binaryPred("Equals", cGEOSEquals_r, other)
}

func (g *Geometry) Covers(other *Geometry) (bool, error) {
	return g.binaryPred("Covers", cGEOSCovers_r, other)
}

func (g *Geometry) CoveredBy(other *Geometry) (bool, error) {
	return g.binaryPred("CoveredBy", cGEOSCoveredBy_r, other)
}

func (g *Geometry) EqualsExact(other *Geometry, tolerance float64) (bool, error) {
	return boolFromC("EqualsExact", cGEOSEqualsExact_r(handle, g.g, other.g, C.double(tolerance)))
}

// Unary predicate functions

func (g *Geometry) IsEmpty() (bool, error) {
	return g.unaryPred("IsEmpty", cGEOSisEmpty_r)
}

func (g *Geometry) IsSimple() (bool, error) {
	return g.unaryPred("IsSimple", cGEOSisSimple_r)
}

func (g *Geometry) IsRing() (bool, error) {
	return g.unaryPred("IsRing", cGEOSisRing_r)
}

func (g *Geometry) HasZ() (bool, error) {
	return g.unaryPred("HasZ", cGEOSHasZ_r)
}

func (g *Geometry) IsClosed() (bool, error) {
	return g.unaryPred("IsClosed", cGEOSisClosed_r)
}

// Geometry info functions

func (g *Geometry) Type() (GeometryType, error) {
	i := C.GEOSGeomTypeId_r(handle, g.g)
	if i == -1 {
		// XXX: error
		return -1, Error()
	}
	return cGeomTypeIds[i], nil
}

func (g *Geometry) SRID() (int, error) {
	return intFromC("SRID", C.GEOSGetSRID_r(handle, g.g), 0)
}

func (g *Geometry) SetSRID(srid int) {
	C.GEOSSetSRID_r(handle, g.g, C.int(srid))
}

func (g *Geometry) NGeometry() (int, error) {
	return intFromC("NGeometry", cGEOSGetNumGeometries_r(handle, g.g), -1)
}

// XXX: method to return a slice of geometries

func (g *Geometry) Geometry(n int) (*Geometry, error) {
	return geomFromC("Geometry", cGEOSGetGeometryN_r(handle, g.g, C.int(n)))
}

// Modifies geometry in-place, clone first if this is not wanted/safe
func (g *Geometry) Normalize() error {
	_, err := intFromC("Normalize", cGEOSNormalize_r(handle, g.g), -1)
	return err
}

func (g *Geometry) NPoint() (int, error) {
	return intFromC("NPoint", cGEOSGeomGetNumPoints_r(handle, g.g), -1)
}

type float64Getter func(C.GEOSContextHandle_t, *C.GEOSGeometry, *C.double) C.int

// Geometry must be a Point
func (g *Geometry) X() (float64, error) {
	return g.float64FromC("X", cGEOSGeomGetX_r, -1)
}

// Geometry must be a Point
func (g *Geometry) Y() (float64, error) {
	return g.float64FromC("Y", cGEOSGeomGetY_r, -1)
}

// Geometry must be a Polygon
func (g *Geometry) Holes() ([]*Geometry, error) {
	n, err := intFromC("NInteriorRing", cGEOSGetNumInteriorRings_r(handle, g.g), -1)
	if err != nil {
		return nil, err
	}
	holes := make([]*Geometry, n)
	for i := 0; i < n; i++ {
		ring, err := geomFromC("InteriorRing", cGEOSGetInteriorRingN_r(handle, g.g, C.int(i)))
		if err != nil {
			return nil, err
		}
		holes[i] = ring
	}
	return holes, nil
}

// Geometry must be a Polygon
func (g *Geometry) Shell() (*Geometry, error) {
	return geomFromC("ExteriorRing", cGEOSGetExteriorRing_r(handle, g.g))
}

func (g *Geometry) NCoordinate() (int, error) {
	return intFromC("NCoordinate", cGEOSGetNumCoordinates_r(handle, g.g), -1)
}

// Geometry must be a LineString, LinearRing, or Point
func (g *Geometry) CoordSeq() (*CoordSeq, error) {
	c := C.GEOSGeom_getCoordSeq_r(handle, g.g)
	if c == nil {
		return nil, Error()
	}
	return coordSeqFromPtr(c), nil
}

// Coords returns a slice of Coord, a sequence of coordinates underlying the
// point, linestring, or linear ring.
func (g *Geometry) Coords() ([]Coord, error) {
	c := C.GEOSGeom_getCoordSeq_r(handle, g.g)
	if c == nil {
		return nil, Error()
	}
	cs := coordSeqFromPtr(c)
	return coordSlice(cs)
}

func (g *Geometry) Dimension() int {
	return int(cGEOSGeom_getDimensions_r(handle, g.g))
}

// Return 2 or 3
func (g *Geometry) CoordDimension() int {
	return int(C.GEOSGeom_getCoordinateDimension_r(handle, g.g))
}

// Geometry must be LineString
func (g *Geometry) Point(n int) (*Geometry, error) {
	return geomFromC("Point", cGEOSGeomGetPointN_r(handle, g.g, C.int(n)))
}

// Geometry must be LineString
func (g *Geometry) StartPoint() (*Geometry, error) {
	return geomFromC("StartPoint", cGEOSGeomGetStartPoint_r(handle, g.g))
}

// Geometry must be LineString
func (g *Geometry) EndPoint() (*Geometry, error) {
	return geomFromC("EndPoint", C.GEOSGeomGetEndPoint_r(handle, g.g))
}

// Misc functions

func (g *Geometry) Area() (float64, error) {
	return g.float64FromC("Area", cGEOSArea_r, 0)
}

func (g *Geometry) Length() (float64, error) {
	return g.float64FromC("Length", cGEOSLength_r, 0)
}

func (g *Geometry) Distance(other *Geometry) (float64, error) {
	return g.binaryFloat("Distance", cGEOSDistance_r, other)
}

func (g *Geometry) HausdorffDistance(other *Geometry) (float64, error) {
	return g.binaryFloat("HausdorffDistance", cGEOSHausdorffDistance_r, other)
}

func (g *Geometry) HausdorffDistanceDensify(other *Geometry, densifyFrac float64) (float64, error) {
	var d C.double
	return float64FromC("HausdorffDistanceDensify", cGEOSHausdorffDistanceDensify_r(handle, g.g, other.g, C.double(densifyFrac), &d), d)
}

// various wrappers around C API

type unaryTopo func(C.GEOSContextHandle_t, *C.GEOSGeometry) *C.GEOSGeometry
type unaryPred func(C.GEOSContextHandle_t, *C.GEOSGeometry) C.char

func (g *Geometry) unaryTopo(name string, cfn unaryTopo) (*Geometry, error) {
	return geomFromC(name, cfn(handle, g.g))
}

func (g *Geometry) unaryPred(name string, cfn unaryPred) (bool, error) {
	return boolFromC(name, cfn(handle, g.g))
}

type binaryTopo func(C.GEOSContextHandle_t, *C.GEOSGeometry, *C.GEOSGeometry) *C.GEOSGeometry
type binaryPred func(C.GEOSContextHandle_t, *C.GEOSGeometry, *C.GEOSGeometry) C.char

func (g *Geometry) binaryTopo(name string, cfn binaryTopo, other *Geometry) (*Geometry, error) {
	return geomFromC(name, cfn(handle, g.g, other.g))
}

func (g *Geometry) binaryPred(name string, cfn binaryPred, other *Geometry) (bool, error) {
	return boolFromC(name, cfn(handle, g.g, other.g))
}

func geomFromCoordSeq(cs *CoordSeq, name string, cfn func(C.GEOSContextHandle_t, *C.GEOSCoordSequence) *C.GEOSGeometry) (*Geometry, error) {
	return geomFromC(name, cfn(handle, cs.c))
}

func emptyGeom(name string, cfn func(C.GEOSContextHandle_t) *C.GEOSGeometry) (*Geometry, error) {
	return geomFromC(name, cfn(handle))
}

func geomFromC(name string, ptr *C.GEOSGeometry) (*Geometry, error) {
	if ptr == nil {
		return nil, Error()
	}
	return geomFromPtr(ptr), nil
}

func boolFromC(name string, c C.char) (bool, error) {
	if c == 2 {
		return false, Error()
	}
	return c == 1, nil
}

func intFromC(name string, i C.int, exception C.int) (int, error) {
	if i == exception {
		return 0, Error()
	}
	return int(i), nil
}

func (g *Geometry) float64FromC(name string, cfn float64Getter, exception C.int) (float64, error) {
	var d C.double
	i := cfn(handle, g.g, &d)
	if i == exception {
		return 0.0, Error()
	}
	return float64(d), nil
}

func float64FromC(name string, rv C.int, d C.double) (float64, error) {
	if rv == 0 {
		return 0.0, Error()
	}
	return float64(d), nil
}

type binaryFloatGetter func(C.GEOSContextHandle_t, *C.GEOSGeometry, *C.GEOSGeometry, *C.double) C.int

func (g *Geometry) binaryFloat(name string, cfn binaryFloatGetter, other *Geometry) (float64, error) {
	var d C.double
	return float64FromC(name, cfn(handle, g.g, other.g, &d), d)
}

func (g *Geometry) simplify(name string, cfn func(C.GEOSContextHandle_t, *C.GEOSGeometry, C.double) *C.GEOSGeometry, d float64) (*Geometry, error) {
	return geomFromC(name, cfn(handle, g.g, C.double(d)))
}
