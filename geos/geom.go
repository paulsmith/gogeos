package geos

/*
#include "geos.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// Geometry represents a geometry object, which can be any one of the types of
// the Simple Features Specification of the Open GIS Consortium:
//	Point
//	LineString
//	LinearRing
//	Polygon
//	MultiPoint
//	MultiLineString
//	MultiPolygon
//	GeometryCollection
type Geometry struct {
	g *C.GEOSGeometry
}

// geomFromPtr returns a new Geometry that's been initialized with a C pointer
// to the GEOS C API object.
func geomFromPtr(p *C.GEOSGeometry) *Geometry {
	g := &Geometry{p}
	runtime.SetFinalizer(g, (*Geometry).destroy)
	return g
}

// FromWKT is a factory function that returns a new Geometry decoded from a
// Well-Known Text (WKT) string.
func FromWKT(wkt string) (*Geometry, error) {
	decoder := newWktDecoder()
	return decoder.decode(wkt)
}

// destroy frees the storage associated with the underlying GEOS C API object.
func (g *Geometry) destroy() {
	C.GEOSGeom_destroy_r(handle, g.g)
	g.g = nil
}

// ToWKT returns a string encoding of the geometry, in Well-Known Text (WKT)
// format.
func (g *Geometry) ToWKT() (string, error) {
	encoder := newWktEncoder()
	return encoder.encode(g)
}

// String returns a string encoding of the geometry, in Well-Known Text (WKT)
// format, or the empty string if there is an error creating the encoding.
func (g *Geometry) String() string {
	str, err := g.ToWKT()
	if err != nil {
		return "" // XXX: better to panic?
	}
	return str
}

// Linearref functions

// Return distance of point 'p' projected on this geometry from origin.
// This must be a lineal geometry.
func (g *Geometry) Project(p *Geometry) float64 {
	// XXX: error if wrong geometry types
	return float64(C.GEOSProject_r(handle, g.g, p.g))
}

// Return closest point to given distance within geometry.
// This geometry must be a LineString.
func (g *Geometry) Interpolate(dist float64) (*Geometry, error) {
	p := C.GEOSInterpolate_r(handle, g.g, C.double(dist))
	// XXX: test for exception
	return geomFromPtr(p), nil
}

// XXX: buffer w number of segments, endcap, join, mitre limit

// Buffer computes a new geometry as the dilation (position amount) or erosion
// (negative amount) of the geometry -- a sum or difference, respectively, of
// the geometry with a circle of radius of the absolute value of the buffer
// amount.
func (g *Geometry) Buffer(d float64) (*Geometry, error) {
	const quadsegs = 8
	return geomFromC("Buffer", cGEOSBuffer_r(handle, g.g, C.double(d), quadsegs))
}

// Geometry Constructors

// NewPoint returns a new geometry of type Point, initialized with the given
// coordinate(s). If no coordinates are given, it's an empty geometry (i.e.,
// IsEmpty() == true). It's an
// error if more than one coordinate is given.
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

// NewLinearRing returns a new geometry of type LinearRing, initialized with the
// given coordinates. The number of coordinates must either be zero (none
// given), in which case it's an empty geometry (IsEmpty() == true), or >= 4.
func NewLinearRing(coords ...Coord) (*Geometry, error) {
	cs, err := coordSeqFromSlice(coords)
	if err != nil {
		return nil, err
	}
	return geomFromCoordSeq(cs, "NewLinearRing", cGEOSGeom_createLinearRing_r)
}

// NewLineString returns a new geometry of type LineString, initialized with the
// given coordinates. If no coordinates are given, it's an empty geometry
// (IsEmpty() == true).
func NewLineString(coords ...Coord) (*Geometry, error) {
	cs, err := coordSeqFromSlice(coords)
	if err != nil {
		return nil, err
	}
	return geomFromCoordSeq(cs, "NewLineString", cGEOSGeom_createLineString_r)
}

// EmptyPolygon returns a new geometry of type Polygon that's empty (i.e.,
// IsEmpty() == true).
func EmptyPolygon() (*Geometry, error) {
	return emptyGeom("EmptyPoint", cGEOSGeom_createEmptyPolygon_r)
}

// NewPolygon returns a new geometry of type Polygon, initialized with the given
// shell (exterior ring) and slice of holes (interior rings). The shell and holes
// slice are themselves slices of coordinates. A shell is required, and a
// variadic number of holes (therefore are optional).
//
// To create a new polygon from existing linear ring Geometry objects, use
// PolygonFromGeom.
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

// PolygonFromGeom returns a new geometry of type Polygon, initialized with the
// given shell (exterior ring) and slice of holes (interior rings). The shell
// and slice of holes are geometry objects, and expected to be LinearRings.
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

// NewCollection returns a new geometry that is a collection containing multiple
// geometries given as variadic arguments. The type of the collection (in the
// SFS sense of type -- MultiPoint, MultiLineString, etc.) is determined by the
// first argument. If no geometries are given, the geometry is an empty version
// of the given collection type.
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

// Clone performs a deep copy on the geometry.
func (g *Geometry) Clone() (*Geometry, error) {
	return g.unaryTopo("Clone", cGEOSGeom_clone_r)
}

// Unary topology functions

// Envelope is the bounding box of a geometry, as a polygon.
func (g *Geometry) Envelope() (*Geometry, error) {
	return g.unaryTopo("Envelope", cGEOSEnvelope_r)
}

// ConvexHull computes the smallest convex geometry that contains all the points
// of the geometry.
func (g *Geometry) ConvexHull() (*Geometry, error) {
	return g.unaryTopo("ConvexHull", cGEOSConvexHull_r)
}

// Boundary is the boundary of the geometry.
func (g *Geometry) Boundary() (*Geometry, error) {
	return g.unaryTopo("Boundary", cGEOSBoundary_r)
}

// UnaryUnion computes the union of all the constituent geometries of the
// geometry.
func (g *Geometry) UnaryUnion() (*Geometry, error) {
	return g.unaryTopo("UnaryUnion", cGEOSUnaryUnion_r)
}

// PointOnSurface computes a point geometry guaranteed to be on the surface of
// the geometry.
func (g *Geometry) PointOnSurface() (*Geometry, error) {
	return g.unaryTopo("PointOnSurface", cGEOSPointOnSurface_r)
}

// Centroid is the center point of the geometry.
func (g *Geometry) Centroid() (*Geometry, error) {
	return g.unaryTopo("Centroid", cGEOSGetCentroid_r)
}

// LineMerge will merge together a collection of LineStrings where they touch
// only at their start and end points. The LineStrings must be fully noded. The
// resulting geometry is a new collection.
func (g *Geometry) LineMerge() (*Geometry, error) {
	return g.unaryTopo("LineMerge", cGEOSLineMerge_r)
}

// Simplify returns a geometry simplified by amount given by tolerance.
// May not preserve topology -- see SimplifyP.
func (g *Geometry) Simplify(tolerance float64) (*Geometry, error) {
	return g.simplify("simplify", cGEOSSimplify_r, tolerance)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (g *Geometry) SimplifyP(tolerance float64) (*Geometry, error) {
	return g.simplify("simplify", cGEOSTopologyPreserveSimplify_r, tolerance)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (g *Geometry) UniquePoints() (*Geometry, error) {
	return g.unaryTopo("UniquePoints", cGEOSGeom_extractUniquePoints_r)
}

// SharedPaths finds paths shared between the two given lineal geometries.
// Returns a GeometryCollection having two elements:
//	- first element is a MultiLineString containing shared paths having the _same_ direction on both inputs
//	- second element is a MultiLineString containing shared paths having the _opposite_ direction on the two inputs
func (g *Geometry) SharedPaths(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("SharedPaths", cGEOSSharedPaths_r, other)
}

// Snap returns a new geometry where the geometry is snapped to the given
// geometry by given tolerance.
func (g *Geometry) Snap(other *Geometry, tolerance float64) (*Geometry, error) {
	return geomFromC("Snap", cGEOSSnap_r(handle, g.g, other.g, C.double(tolerance)))
}

// Prepared returns a new prepared geometry from the geometry -- see PGeometry
func (g *Geometry) Prepare() *PGeometry {
	return PrepareGeometry(g)
}

// Binary topology functions

// Intersection returns a new geometry representing the points shared by this
// geometry and the other.
func (g *Geometry) Intersection(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("Intersection", cGEOSIntersection_r, other)
}

// Difference returns a new geometry representing the points making up this
// geometry that do not make up the other.
func (g *Geometry) Difference(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("Difference", cGEOSDifference_r, other)
}

// SymDifference returns a new geometry representing the set combining the
// points in this geometry not in the other, and the points in the other
// geometry and not in this.
func (g *Geometry) SymDifference(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("SymDifference", cGEOSSymDifference_r, other)
}

// Union returns a new geometry representing all points in this geometry and the
// other.
func (g *Geometry) Union(other *Geometry) (*Geometry, error) {
	return g.binaryTopo("Union", cGEOSUnion_r, other)
}

// Binary predicate functions

// Disjoint returns true if the two geometries have no point in common.
func (g *Geometry) Disjoint(other *Geometry) (bool, error) {
	return g.binaryPred("Disjoint", cGEOSDisjoint_r, other)
}

// Touches returns true if the two geometries have at least one point in common,
// but their interiors do not intersect.
func (g *Geometry) Touches(other *Geometry) (bool, error) {
	return g.binaryPred("Touches", cGEOSTouches_r, other)
}

// Intersects returns true if the two geometries have at least one point in
// common.
func (g *Geometry) Intersects(other *Geometry) (bool, error) {
	return g.binaryPred("Intersects", cGEOSIntersects_r, other)
}

// Crosses returns true if the two geometries have some but not all interior
// points in common.
func (g *Geometry) Crosses(other *Geometry) (bool, error) {
	return g.binaryPred("Crosses", cGEOSCrosses_r, other)
}

// Within returns true if every point of this geometry is a point of the other,
// and the interiors of the two geometries have at least one point in common.
func (g *Geometry) Within(other *Geometry) (bool, error) {
	return g.binaryPred("Within", cGEOSWithin_r, other)
}

// Contains returns true if every point of the other is a point of this geometry,
// and the interiors of the two geometries have at least one point in common.
func (g *Geometry) Contains(other *Geometry) (bool, error) {
	return g.binaryPred("Contains", cGEOSContains_r, other)
}

// Overlaps returns true if the geometries have some but not all points in
// common, they have the same dimension, and the intersection of the interiors
// of the two geometries has the same dimension as the geometries themselves.
func (g *Geometry) Overlaps(other *Geometry) (bool, error) {
	return g.binaryPred("Overlaps", cGEOSOverlaps_r, other)
}

// Equals returns true if the two geometries have at least one point in common,
// and no point of either geometry lies in the exterior of the other geometry.
func (g *Geometry) Equals(other *Geometry) (bool, error) {
	return g.binaryPred("Equals", cGEOSEquals_r, other)
}

// Covers returns true if every point of the other geometry is a point of this
// geometry.
func (g *Geometry) Covers(other *Geometry) (bool, error) {
	return g.binaryPred("Covers", cGEOSCovers_r, other)
}

// CoveredBy returns true if every point of this geometry is a point of the
// other geometry.
func (g *Geometry) CoveredBy(other *Geometry) (bool, error) {
	return g.binaryPred("CoveredBy", cGEOSCoveredBy_r, other)
}

// EqualsExact returns true if both geometries are Equal, as evaluated by their
// points being within the given tolerance.
func (g *Geometry) EqualsExact(other *Geometry, tolerance float64) (bool, error) {
	return boolFromC("EqualsExact", cGEOSEqualsExact_r(handle, g.g, other.g, C.double(tolerance)))
}

// Unary predicate functions

// IsEmpty returns true if the set of points of this geometry is empty (i.e.,
// the empty geometry).
func (g *Geometry) IsEmpty() (bool, error) {
	return g.unaryPred("IsEmpty", cGEOSisEmpty_r)
}

// IsSimple returns true iff the only self-intersections are at boundary points.
func (g *Geometry) IsSimple() (bool, error) {
	return g.unaryPred("IsSimple", cGEOSisSimple_r)
}

// IsRing returns true if the lineal geometry has the ring property.
func (g *Geometry) IsRing() (bool, error) {
	return g.unaryPred("IsRing", cGEOSisRing_r)
}

// HasZ returns true if the geometry is 3D.
func (g *Geometry) HasZ() (bool, error) {
	return g.unaryPred("HasZ", cGEOSHasZ_r)
}

// IsClosed returns true if the geometry is closed (i.e., start & end points
// equal).
func (g *Geometry) IsClosed() (bool, error) {
	return g.unaryPred("IsClosed", cGEOSisClosed_r)
}

// Geometry info functions

// Type returns the SFS type of the geometry.
func (g *Geometry) Type() (GeometryType, error) {
	i := C.GEOSGeomTypeId_r(handle, g.g)
	if i == -1 {
		// XXX: error
		return -1, Error()
	}
	return cGeomTypeIds[i], nil
}

// SRID returns the geometry's SRID, if set.
func (g *Geometry) SRID() (int, error) {
	return intFromC("SRID", C.GEOSGetSRID_r(handle, g.g), 0)
}

// SetSRID sets the geometry's SRID.
func (g *Geometry) SetSRID(srid int) {
	C.GEOSSetSRID_r(handle, g.g, C.int(srid))
}

// NGeometry returns the number of component geometries (eg., for
// a collection).
func (g *Geometry) NGeometry() (int, error) {
	return intFromC("NGeometry", cGEOSGetNumGeometries_r(handle, g.g), -1)
}

// XXX: method to return a slice of geometries

// Geometry returns the nth sub-geometry of the geometry (eg., of a collection).
func (g *Geometry) Geometry(n int) (*Geometry, error) {
	return geomFromC("Geometry", cGEOSGetGeometryN_r(handle, g.g, C.int(n)))
}

// Normalize computes the normal form of the geometry.
// Modifies geometry in-place, clone first if this is not wanted/safe.
func (g *Geometry) Normalize() error {
	_, err := intFromC("Normalize", cGEOSNormalize_r(handle, g.g), -1)
	return err
}

// NPoint returns the number of points in the geometry.
func (g *Geometry) NPoint() (int, error) {
	return intFromC("NPoint", cGEOSGeomGetNumPoints_r(handle, g.g), -1)
}

type float64Getter func(C.GEOSContextHandle_t, *C.GEOSGeometry, *C.double) C.int

// X returns the x ordinate of the geometry.
// Geometry must be a Point.
func (g *Geometry) X() (float64, error) {
	return g.float64FromC("X", cGEOSGeomGetX_r, -1)
}

// Y returns the y ordinate of the geometry.
// Geometry must be a Point
func (g *Geometry) Y() (float64, error) {
	return g.float64FromC("Y", cGEOSGeomGetY_r, -1)
}

// Holes returns a slice of geometries (LinearRings) representing the interior
// rings of a polygon (possibly nil).
// Geometry must be a Polygon.
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

// XXX: Holes() returns a [][]Coord?

// Shell returns the exterior ring (a LinearRing) of the geometry.
// Geometry must be a Polygon.
func (g *Geometry) Shell() (*Geometry, error) {
	return geomFromC("ExteriorRing", cGEOSGetExteriorRing_r(handle, g.g))
}

// NCoordinate returns the number of coordinates of the geometry.
func (g *Geometry) NCoordinate() (int, error) {
	return intFromC("NCoordinate", cGEOSGetNumCoordinates_r(handle, g.g), -1)
}

// Geometry must be a LineString, LinearRing, or Point.
func (g *Geometry) coordSeq() (*CoordSeq, error) {
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

// Dimension returns the number of dimensions geometry, eg., 1 for point, 2 for
// linestring.
func (g *Geometry) Dimension() int {
	return int(cGEOSGeom_getDimensions_r(handle, g.g))
}

// CoordDimension returns the number of dimensions of the coordinates of the
// geometry (2 or 3).
func (g *Geometry) CoordDimension() int {
	return int(C.GEOSGeom_getCoordinateDimension_r(handle, g.g))
}

// Point returns the nth point of the geometry.
// Geometry must be LineString.
func (g *Geometry) Point(n int) (*Geometry, error) {
	return geomFromC("Point", cGEOSGeomGetPointN_r(handle, g.g, C.int(n)))
}

// StartPoint returns the 0th point of the geometry.
// Geometry must be LineString.
func (g *Geometry) StartPoint() (*Geometry, error) {
	return geomFromC("StartPoint", cGEOSGeomGetStartPoint_r(handle, g.g))
}

// EndPoint returns the (n-1)th point of the geometry.
// Geometry must be LineString.
func (g *Geometry) EndPoint() (*Geometry, error) {
	return geomFromC("EndPoint", C.GEOSGeomGetEndPoint_r(handle, g.g))
}

// Misc functions

// Area returns the area of the geometry, which must be a areal geometry like
// a polygon or multipolygon.
func (g *Geometry) Area() (float64, error) {
	return g.float64FromC("Area", cGEOSArea_r, 0)
}

// Length returns the length of the geometry, which must be a lineal geometry
// like a linestring or linear ring.
func (g *Geometry) Length() (float64, error) {
	return g.float64FromC("Length", cGEOSLength_r, 0)
}

// Distance returns the Cartesian distance between the two geometries.
func (g *Geometry) Distance(other *Geometry) (float64, error) {
	return g.binaryFloat("Distance", cGEOSDistance_r, other)
}

// HausdorffDistance returns the maximum distance of the geometry to the nearest
// point in the other geometry (i.e., considers the whole shape and position of
// the geometries).
func (g *Geometry) HausdorffDistance(other *Geometry) (float64, error) {
	return g.binaryFloat("HausdorffDistance", cGEOSHausdorffDistance_r, other)
}

func (g *Geometry) HausdorffDistanceDensify(other *Geometry, densifyFrac float64) (float64, error) {
	var d C.double
	return float64FromC("HausdorffDistanceDensify", cGEOSHausdorffDistanceDensify_r(handle, g.g, other.g, C.double(densifyFrac), &d), d)
}

// DE-9IM

// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func (g *Geometry) Relate(other *Geometry) (string, error) {
	cs := cGEOSRelate_r(handle, g.g, other.g)
	if cs == nil {
		return "", Error()
	}
	s := C.GoString(cs)
	C.GEOSFree_r(handle, unsafe.Pointer(cs))
	return s, nil
}

// RelatePat returns true if the DE-9IM matrix equals the intersection matrix of
// the two geometries.
func (g *Geometry) RelatePat(other *Geometry, pat string) (bool, error) {
	cs := C.CString(pat)
	defer C.free(unsafe.Pointer(cs))
	return boolFromC("RelatePat", cGEOSRelatePattern_r(handle, g.g, other.g, cs))
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
