package geos

// Created mechanically from C API header - DO NOT EDIT

/*
#include <geos_c.h>
*/
import "C"

import (
	"unsafe"
)

func cinitGEOS_r(notice_function C.GEOSMessageHandler, error_function C.GEOSMessageHandler) C.GEOSContextHandle_t {
	return C.initGEOS_r(notice_function, error_function)
}

func cfinishGEOS_r(handle C.GEOSContextHandle_t) {
	C.finishGEOS_r(handle)
}

func cGEOSversion() *C.char {
	return C.GEOSversion()
}

func cGEOSGeomFromWKT_r(handle C.GEOSContextHandle_t, wkt *C.char) *C.GEOSGeometry {
	return C.GEOSGeomFromWKT_r(handle, wkt)
}

func cGEOSGeomToWKT_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.char {
	return C.GEOSGeomToWKT_r(handle, g)
}

func cGEOS_getWKBOutputDims_r(handle C.GEOSContextHandle_t) C.int {
	return C.GEOS_getWKBOutputDims_r(handle)
}

func cGEOS_setWKBOutputDims_r(handle C.GEOSContextHandle_t, newDims C.int) C.int {
	return C.GEOS_setWKBOutputDims_r(handle, newDims)
}

func cGEOS_getWKBByteOrder_r(handle C.GEOSContextHandle_t) C.int {
	return C.GEOS_getWKBByteOrder_r(handle)
}

func cGEOS_setWKBByteOrder_r(handle C.GEOSContextHandle_t, byteOrder C.int) C.int {
	return C.GEOS_setWKBByteOrder_r(handle, byteOrder)
}

func cGEOSGeomFromWKB_buf_r(handle C.GEOSContextHandle_t, wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	return C.GEOSGeomFromWKB_buf_r(handle, wkb, size)
}

func cGEOSGeomToWKB_buf_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	return C.GEOSGeomToWKB_buf_r(handle, g, size)
}

func cGEOSGeomFromHEX_buf_r(handle C.GEOSContextHandle_t, hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	return C.GEOSGeomFromHEX_buf_r(handle, hex, size)
}

func cGEOSGeomToHEX_buf_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	return C.GEOSGeomToHEX_buf_r(handle, g, size)
}

func cGEOSCoordSeq_create_r(handle C.GEOSContextHandle_t, size C.uint, dims C.uint) *C.GEOSCoordSequence {
	return C.GEOSCoordSeq_create_r(handle, size, dims)
}

func cGEOSCoordSeq_clone_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence) *C.GEOSCoordSequence {
	return C.GEOSCoordSeq_clone_r(handle, s)
}

func cGEOSCoordSeq_destroy_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence) {
	C.GEOSCoordSeq_destroy_r(handle, s)
}

func cGEOSCoordSeq_setX_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	return C.GEOSCoordSeq_setX_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setY_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	return C.GEOSCoordSeq_setY_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setZ_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	return C.GEOSCoordSeq_setZ_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setOrdinate_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val C.double) C.int {
	return C.GEOSCoordSeq_setOrdinate_r(handle, s, idx, dim, val)
}

func cGEOSCoordSeq_getX_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	return C.GEOSCoordSeq_getX_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getY_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	return C.GEOSCoordSeq_getY_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getZ_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	return C.GEOSCoordSeq_getZ_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getOrdinate_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val *C.double) C.int {
	return C.GEOSCoordSeq_getOrdinate_r(handle, s, idx, dim, val)
}

func cGEOSCoordSeq_getSize_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, size *C.uint) C.int {
	return C.GEOSCoordSeq_getSize_r(handle, s, size)
}

func cGEOSCoordSeq_getDimensions_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence, dims *C.uint) C.int {
	return C.GEOSCoordSeq_getDimensions_r(handle, s, dims)
}

func cGEOSProject_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	return C.GEOSProject_r(handle, g, p)
}

func cGEOSInterpolate_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	return C.GEOSInterpolate_r(handle, g, d)
}

func cGEOSProjectNormalized_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	return C.GEOSProjectNormalized_r(handle, g, p)
}

func cGEOSInterpolateNormalized_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	return C.GEOSInterpolateNormalized_r(handle, g, d)
}

func cGEOSBufferParams_create_r(handle C.GEOSContextHandle_t) *C.GEOSBufferParams {
	return C.GEOSBufferParams_create_r(handle)
}

func cGEOSBufferParams_destroy_r(handle C.GEOSContextHandle_t, parms *C.GEOSBufferParams) {
	C.GEOSBufferParams_destroy_r(handle, parms)
}

func cGEOSBufferParams_setEndCapStyle_r(handle C.GEOSContextHandle_t, p *C.GEOSBufferParams, style C.int) C.int {
	return C.GEOSBufferParams_setEndCapStyle_r(handle, p, style)
}

func cGEOSBufferParams_setJoinStyle_r(handle C.GEOSContextHandle_t, p *C.GEOSBufferParams, joinStyle C.int) C.int {
	return C.GEOSBufferParams_setJoinStyle_r(handle, p, joinStyle)
}

func cGEOSBufferParams_setMitreLimit_r(handle C.GEOSContextHandle_t, p *C.GEOSBufferParams, mitreLimit C.double) C.int {
	return C.GEOSBufferParams_setMitreLimit_r(handle, p, mitreLimit)
}

func cGEOSBufferParams_setQuadrantSegments_r(handle C.GEOSContextHandle_t, p *C.GEOSBufferParams, quadSegs C.int) C.int {
	return C.GEOSBufferParams_setQuadrantSegments_r(handle, p, quadSegs)
}

func cGEOSBufferParams_setSingleSided_r(handle C.GEOSContextHandle_t, p *C.GEOSBufferParams, singleSided C.int) C.int {
	return C.GEOSBufferParams_setSingleSided_r(handle, p, singleSided)
}

func cGEOSBufferWithParams_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, p *C.GEOSBufferParams, width C.double) *C.GEOSGeometry {
	return C.GEOSBufferWithParams_r(handle, g1, p, width)
}

func cGEOSBuffer_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, width C.double, quadsegs C.int) *C.GEOSGeometry {
	return C.GEOSBuffer_r(handle, g1, width, quadsegs)
}

func cGEOSBufferWithStyle_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, width C.double, quadsegs C.int, endCapStyle C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	return C.GEOSBufferWithStyle_r(handle, g1, width, quadsegs, endCapStyle, joinStyle, mitreLimit)
}

func cGEOSSingleSidedBuffer_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double, leftSide C.int) *C.GEOSGeometry {
	return C.GEOSSingleSidedBuffer_r(handle, g1, width, quadsegs, joinStyle, mitreLimit, leftSide)
}

func cGEOSOffsetCurve_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	return C.GEOSOffsetCurve_r(handle, g1, width, quadsegs, joinStyle, mitreLimit)
}

func cGEOSGeom_createPoint_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence) *C.GEOSGeometry {
	return C.GEOSGeom_createPoint_r(handle, s)
}

func cGEOSGeom_createEmptyPoint_r(handle C.GEOSContextHandle_t) *C.GEOSGeometry {
	return C.GEOSGeom_createEmptyPoint_r(handle)
}

func cGEOSGeom_createLinearRing_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence) *C.GEOSGeometry {
	return C.GEOSGeom_createLinearRing_r(handle, s)
}

func cGEOSGeom_createLineString_r(handle C.GEOSContextHandle_t, s *C.GEOSCoordSequence) *C.GEOSGeometry {
	return C.GEOSGeom_createLineString_r(handle, s)
}

func cGEOSGeom_createEmptyLineString_r(handle C.GEOSContextHandle_t) *C.GEOSGeometry {
	return C.GEOSGeom_createEmptyLineString_r(handle)
}

func cGEOSGeom_createEmptyPolygon_r(handle C.GEOSContextHandle_t) *C.GEOSGeometry {
	return C.GEOSGeom_createEmptyPolygon_r(handle)
}

func cGEOSGeom_createPolygon_r(handle C.GEOSContextHandle_t, shell *C.GEOSGeometry, holes **C.GEOSGeometry, nholes C.uint) *C.GEOSGeometry {
	return C.GEOSGeom_createPolygon_r(handle, shell, holes, nholes)
}

func cGEOSGeom_createCollection_r(handle C.GEOSContextHandle_t, _type C.int, geoms **C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	return C.GEOSGeom_createCollection_r(handle, _type, geoms, ngeoms)
}

func cGEOSGeom_createEmptyCollection_r(handle C.GEOSContextHandle_t, _type C.int) *C.GEOSGeometry {
	return C.GEOSGeom_createEmptyCollection_r(handle, _type)
}

func cGEOSGeom_clone_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSGeom_clone_r(handle, g)
}

func cGEOSGeom_destroy_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) {
	C.GEOSGeom_destroy_r(handle, g)
}

func cGEOSEnvelope_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSEnvelope_r(handle, g1)
}

func cGEOSIntersection_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSIntersection_r(handle, g1, g2)
}

func cGEOSConvexHull_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSConvexHull_r(handle, g1)
}

func cGEOSDifference_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSDifference_r(handle, g1, g2)
}

func cGEOSSymDifference_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSSymDifference_r(handle, g1, g2)
}

func cGEOSBoundary_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSBoundary_r(handle, g1)
}

func cGEOSUnion_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSUnion_r(handle, g1, g2)
}

func cGEOSUnaryUnion_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSUnaryUnion_r(handle, g)
}

func cGEOSUnionCascaded_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSUnionCascaded_r(handle, g1)
}

func cGEOSPointOnSurface_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSPointOnSurface_r(handle, g1)
}

func cGEOSGetCentroid_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSGetCentroid_r(handle, g)
}

func cGEOSPolygonize_r(handle C.GEOSContextHandle_t, geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	return C.GEOSPolygonize_r(handle, &geoms[0], ngeoms)
}

func cGEOSPolygonizer_getCutEdges_r(handle C.GEOSContextHandle_t, geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	return C.GEOSPolygonizer_getCutEdges_r(handle, &geoms[0], ngeoms)
}

func cGEOSPolygonize_full_r(handle C.GEOSContextHandle_t, input *C.GEOSGeometry, cuts **C.GEOSGeometry, dangles **C.GEOSGeometry, invalidRings **C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSPolygonize_full_r(handle, input, cuts, dangles, invalidRings)
}

func cGEOSLineMerge_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSLineMerge_r(handle, g)
}

func cGEOSSimplify_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	return C.GEOSSimplify_r(handle, g1, tolerance)
}

func cGEOSTopologyPreserveSimplify_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	return C.GEOSTopologyPreserveSimplify_r(handle, g1, tolerance)
}

func cGEOSGeom_extractUniquePoints_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSGeom_extractUniquePoints_r(handle, g)
}

func cGEOSSharedPaths_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSSharedPaths_r(handle, g1, g2)
}

func cGEOSSnap_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	return C.GEOSSnap_r(handle, g1, g2, tolerance)
}

func cGEOSDisjoint_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSDisjoint_r(handle, g1, g2)
}

func cGEOSTouches_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSTouches_r(handle, g1, g2)
}

func cGEOSIntersects_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSIntersects_r(handle, g1, g2)
}

func cGEOSCrosses_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSCrosses_r(handle, g1, g2)
}

func cGEOSWithin_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSWithin_r(handle, g1, g2)
}

func cGEOSContains_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSContains_r(handle, g1, g2)
}

func cGEOSOverlaps_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSOverlaps_r(handle, g1, g2)
}

func cGEOSEquals_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSEquals_r(handle, g1, g2)
}

func cGEOSEqualsExact_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) C.char {
	return C.GEOSEqualsExact_r(handle, g1, g2, tolerance)
}

func cGEOSCovers_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSCovers_r(handle, g1, g2)
}

func cGEOSCoveredBy_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSCoveredBy_r(handle, g1, g2)
}

func cGEOSPrepare_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSPreparedGeometry {
	return C.GEOSPrepare_r(handle, g)
}

func cGEOSPreparedGeom_destroy_r(handle C.GEOSContextHandle_t, g *C.GEOSPreparedGeometry) {
	C.GEOSPreparedGeom_destroy_r(handle, g)
}

func cGEOSPreparedContains_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedContains_r(handle, pg1, g2)
}

func cGEOSPreparedContainsProperly_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedContainsProperly_r(handle, pg1, g2)
}

func cGEOSPreparedCoveredBy_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedCoveredBy_r(handle, pg1, g2)
}

func cGEOSPreparedCovers_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedCovers_r(handle, pg1, g2)
}

func cGEOSPreparedCrosses_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedCrosses_r(handle, pg1, g2)
}

func cGEOSPreparedDisjoint_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedDisjoint_r(handle, pg1, g2)
}

func cGEOSPreparedIntersects_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedIntersects_r(handle, pg1, g2)
}

func cGEOSPreparedOverlaps_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedOverlaps_r(handle, pg1, g2)
}

func cGEOSPreparedTouches_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedTouches_r(handle, pg1, g2)
}

func cGEOSPreparedWithin_r(handle C.GEOSContextHandle_t, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return C.GEOSPreparedWithin_r(handle, pg1, g2)
}

func cGEOSSTRtree_create_r(handle C.GEOSContextHandle_t, nodeCapacity C.size_t) *C.GEOSSTRtree {
	return C.GEOSSTRtree_create_r(handle, nodeCapacity)
}

func cGEOSSTRtree_insert_r(handle C.GEOSContextHandle_t, tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) {
	C.GEOSSTRtree_insert_r(handle, tree, g, unsafe.Pointer(item))
}

func cGEOSSTRtree_query_r(handle C.GEOSContextHandle_t, tree *C.GEOSSTRtree, g *C.GEOSGeometry, callback C.GEOSQueryCallback, userdata *C.void) {
	C.GEOSSTRtree_query_r(handle, tree, g, callback, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_iterate_r(handle C.GEOSContextHandle_t, tree *C.GEOSSTRtree, callback C.GEOSQueryCallback, userdata *C.void) {
	C.GEOSSTRtree_iterate_r(handle, tree, callback, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_remove_r(handle C.GEOSContextHandle_t, tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) C.char {
	return C.GEOSSTRtree_remove_r(handle, tree, g, unsafe.Pointer(item))
}

func cGEOSSTRtree_destroy_r(handle C.GEOSContextHandle_t, tree *C.GEOSSTRtree) {
	C.GEOSSTRtree_destroy_r(handle, tree)
}

func cGEOSisEmpty_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.char {
	return C.GEOSisEmpty_r(handle, g1)
}

func cGEOSisSimple_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.char {
	return C.GEOSisSimple_r(handle, g1)
}

func cGEOSisRing_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.char {
	return C.GEOSisRing_r(handle, g1)
}

func cGEOSHasZ_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.char {
	return C.GEOSHasZ_r(handle, g1)
}

func cGEOSisClosed_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.char {
	return C.GEOSisClosed_r(handle, g1)
}

func cGEOSRelatePattern_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, pat *C.char) C.char {
	return C.GEOSRelatePattern_r(handle, g1, g2, pat)
}

func cGEOSRelate_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.char {
	return C.GEOSRelate_r(handle, g1, g2)
}

func cGEOSRelatePatternMatch_r(handle C.GEOSContextHandle_t, mat *C.char, pat *C.char) C.char {
	return C.GEOSRelatePatternMatch_r(handle, mat, pat)
}

func cGEOSRelateBoundaryNodeRule_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, bnr C.int) *C.char {
	return C.GEOSRelateBoundaryNodeRule_r(handle, g1, g2, bnr)
}

func cGEOSisValid_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.char {
	return C.GEOSisValid_r(handle, g1)
}

func cGEOSisValidReason_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.char {
	return C.GEOSisValidReason_r(handle, g1)
}

func cGEOSisValidDetail_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, flags C.int, reason **C.char, location **C.GEOSGeometry) C.char {
	return C.GEOSisValidDetail_r(handle, g, flags, reason, location)
}

func cGEOSGeomType_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.char {
	return C.GEOSGeomType_r(handle, g1)
}

func cGEOSGeomTypeId_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.int {
	return C.GEOSGeomTypeId_r(handle, g1)
}

func cGEOSGetSRID_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) C.int {
	return C.GEOSGetSRID_r(handle, g)
}

func cGEOSSetSRID_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, SRID C.int) {
	C.GEOSSetSRID_r(handle, g, SRID)
}

func cGEOSGetNumGeometries_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) C.int {
	return C.GEOSGetNumGeometries_r(handle, g)
}

func cGEOSGetGeometryN_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	return C.GEOSGetGeometryN_r(handle, g, n)
}

func cGEOSNormalize_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.int {
	return C.GEOSNormalize_r(handle, g1)
}

func cGEOSGetNumInteriorRings_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.int {
	return C.GEOSGetNumInteriorRings_r(handle, g1)
}

func cGEOSGeomGetNumPoints_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.int {
	return C.GEOSGeomGetNumPoints_r(handle, g1)
}

func cGEOSGeomGetX_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, x *C.double) C.int {
	return C.GEOSGeomGetX_r(handle, g1, x)
}

func cGEOSGeomGetY_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, y *C.double) C.int {
	return C.GEOSGeomGetY_r(handle, g1, y)
}

func cGEOSGetInteriorRingN_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	return C.GEOSGetInteriorRingN_r(handle, g, n)
}

func cGEOSGetExteriorRing_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSGetExteriorRing_r(handle, g)
}

func cGEOSGetNumCoordinates_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) C.int {
	return C.GEOSGetNumCoordinates_r(handle, g1)
}

func cGEOSGeom_getCoordSeq_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) *C.GEOSCoordSequence {
	return C.GEOSGeom_getCoordSeq_r(handle, g)
}

func cGEOSGeom_getDimensions_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) C.int {
	return C.GEOSGeom_getDimensions_r(handle, g)
}

func cGEOSGeom_getCoordinateDimension_r(handle C.GEOSContextHandle_t, g *C.GEOSGeometry) C.int {
	return C.GEOSGeom_getCoordinateDimension_r(handle, g)
}

func cGEOSGeomGetPointN_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	return C.GEOSGeomGetPointN_r(handle, g1, n)
}

func cGEOSGeomGetStartPoint_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSGeomGetStartPoint_r(handle, g1)
}

func cGEOSGeomGetEndPoint_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return C.GEOSGeomGetEndPoint_r(handle, g1)
}

func cGEOSArea_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, area *C.double) C.int {
	return C.GEOSArea_r(handle, g1, area)
}

func cGEOSLength_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, length *C.double) C.int {
	return C.GEOSLength_r(handle, g1, length)
}

func cGEOSDistance_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	return C.GEOSDistance_r(handle, g1, g2, dist)
}

func cGEOSHausdorffDistance_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	return C.GEOSHausdorffDistance_r(handle, g1, g2, dist)
}

func cGEOSHausdorffDistanceDensify_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, densifyFrac C.double, dist *C.double) C.int {
	return C.GEOSHausdorffDistanceDensify_r(handle, g1, g2, densifyFrac, dist)
}

func cGEOSGeomGetLength_r(handle C.GEOSContextHandle_t, g1 *C.GEOSGeometry, length *C.double) C.int {
	return C.GEOSGeomGetLength_r(handle, g1, length)
}

func cGEOSOrientationIndex_r(handle C.GEOSContextHandle_t, Ax C.double, Ay C.double, Bx C.double, By C.double, Px C.double, Py C.double) C.int {
	return C.GEOSOrientationIndex_r(handle, Ax, Ay, Bx, By, Px, Py)
}

func cGEOSWKTReader_create_r(handle C.GEOSContextHandle_t) *C.GEOSWKTReader {
	return C.GEOSWKTReader_create_r(handle)
}

func cGEOSWKTReader_destroy_r(handle C.GEOSContextHandle_t, reader *C.GEOSWKTReader) {
	C.GEOSWKTReader_destroy_r(handle, reader)
}

func cGEOSWKTReader_read_r(handle C.GEOSContextHandle_t, reader *C.GEOSWKTReader, wkt *C.char) *C.GEOSGeometry {
	return C.GEOSWKTReader_read_r(handle, reader, wkt)
}

func cGEOSWKTWriter_create_r(handle C.GEOSContextHandle_t) *C.GEOSWKTWriter {
	return C.GEOSWKTWriter_create_r(handle)
}

func cGEOSWKTWriter_destroy_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKTWriter) {
	C.GEOSWKTWriter_destroy_r(handle, writer)
}

func cGEOSWKTWriter_write_r(handle C.GEOSContextHandle_t, reader *C.GEOSWKTWriter, g *C.GEOSGeometry) *C.char {
	return C.GEOSWKTWriter_write_r(handle, reader, g)
}

func cGEOSWKTWriter_setTrim_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKTWriter, trim C.char) {
	C.GEOSWKTWriter_setTrim_r(handle, writer, trim)
}

func cGEOSWKTWriter_setRoundingPrecision_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKTWriter, precision C.int) {
	C.GEOSWKTWriter_setRoundingPrecision_r(handle, writer, precision)
}

func cGEOSWKTWriter_setOutputDimension_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKTWriter, dim C.int) {
	C.GEOSWKTWriter_setOutputDimension_r(handle, writer, dim)
}

func cGEOSWKTWriter_getOutputDimension_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKTWriter) C.int {
	return C.GEOSWKTWriter_getOutputDimension_r(handle, writer)
}

func cGEOSWKTWriter_setOld3D_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKTWriter, useOld3D C.int) {
	C.GEOSWKTWriter_setOld3D_r(handle, writer, useOld3D)
}

func cGEOSWKBReader_create_r(handle C.GEOSContextHandle_t) *C.GEOSWKBReader {
	return C.GEOSWKBReader_create_r(handle)
}

func cGEOSWKBReader_destroy_r(handle C.GEOSContextHandle_t, reader *C.GEOSWKBReader) {
	C.GEOSWKBReader_destroy_r(handle, reader)
}

func cGEOSWKBReader_read_r(handle C.GEOSContextHandle_t, reader *C.GEOSWKBReader, wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	return C.GEOSWKBReader_read_r(handle, reader, wkb, size)
}

func cGEOSWKBReader_readHEX_r(handle C.GEOSContextHandle_t, reader *C.GEOSWKBReader, hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	return C.GEOSWKBReader_readHEX_r(handle, reader, hex, size)
}

func cGEOSWKBWriter_create_r(handle C.GEOSContextHandle_t) *C.GEOSWKBWriter {
	return C.GEOSWKBWriter_create_r(handle)
}

func cGEOSWKBWriter_destroy_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter) {
	C.GEOSWKBWriter_destroy_r(handle, writer)
}

func cGEOSWKBWriter_write_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	return C.GEOSWKBWriter_write_r(handle, writer, g, size)
}

func cGEOSWKBWriter_writeHEX_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	return C.GEOSWKBWriter_writeHEX_r(handle, writer, g, size)
}

func cGEOSWKBWriter_getOutputDimension_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter) C.int {
	return C.GEOSWKBWriter_getOutputDimension_r(handle, writer)
}

func cGEOSWKBWriter_setOutputDimension_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter, newDimension C.int) {
	C.GEOSWKBWriter_setOutputDimension_r(handle, writer, newDimension)
}

func cGEOSWKBWriter_getByteOrder_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter) C.int {
	return C.GEOSWKBWriter_getByteOrder_r(handle, writer)
}

func cGEOSWKBWriter_setByteOrder_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter, byteOrder C.int) {
	C.GEOSWKBWriter_setByteOrder_r(handle, writer, byteOrder)
}

func cGEOSWKBWriter_getIncludeSRID_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter) C.char {
	return C.GEOSWKBWriter_getIncludeSRID_r(handle, writer)
}

func cGEOSWKBWriter_setIncludeSRID_r(handle C.GEOSContextHandle_t, writer *C.GEOSWKBWriter, writeSRID C.char) {
	C.GEOSWKBWriter_setIncludeSRID_r(handle, writer, writeSRID)
}
