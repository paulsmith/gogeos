package geos

/*
#include "geos.h"
*/
import "C"

type GeometryType C.int

const (
	POINT              GeometryType = C.GEOS_POINT
	LINESTRING         GeometryType = C.GEOS_LINESTRING
	LINEARRING         GeometryType = C.GEOS_LINEARRING
	POLYGON            GeometryType = C.GEOS_POLYGON
	MULTIPOINT         GeometryType = C.GEOS_MULTIPOINT
	MULTILINESTRING    GeometryType = C.GEOS_MULTILINESTRING
	MULTIPOLYGON       GeometryType = C.GEOS_MULTIPOLYGON
	GEOMETRYCOLLECTION GeometryType = C.GEOS_GEOMETRYCOLLECTION
)

var cGeomTypeIds = map[C.int]GeometryType{
	C.GEOS_POINT:              POINT,
	C.GEOS_LINESTRING:         LINESTRING,
	C.GEOS_LINEARRING:         LINEARRING,
	C.GEOS_POLYGON:            POLYGON,
	C.GEOS_MULTIPOINT:         MULTIPOINT,
	C.GEOS_MULTILINESTRING:    MULTILINESTRING,
	C.GEOS_MULTIPOLYGON:       MULTIPOLYGON,
	C.GEOS_GEOMETRYCOLLECTION: GEOMETRYCOLLECTION,
}

var geometryTypes = map[GeometryType]string{
	POINT:              "Point",
	LINESTRING:         "LineString",
	LINEARRING:         "LinearRing",
	POLYGON:            "Polygon",
	MULTIPOINT:         "MultiPoint",
	MULTILINESTRING:    "MultiLineString",
	MULTIPOLYGON:       "MultiPolygon",
	GEOMETRYCOLLECTION: "GeometryCollection",
}

func (t GeometryType) String() string {
	return geometryTypes[t]
}
