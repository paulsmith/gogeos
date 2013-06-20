package geos

// Created mechanically from C API header - DO NOT EDIT

/*
#include <geos_c.h>
*/
import "C"

import (
	"unsafe"
)

func cinitGEOS(notice_function C.GEOSMessageHandler, error_function C.GEOSMessageHandler) C.GEOSContextHandle_t {
	return C.initGEOS_r(notice_function, error_function)
}

func cfinishGEOS() {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.finishGEOS_r(handle)
}

func cGEOSversion() *C.char {
	return C.GEOSversion()
}

func cGEOSGeomFromWKT(wkt *C.char) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromWKT_r(handle, wkt)
}

func cGEOSGeomToWKT(g *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToWKT_r(handle, g)
}

func cGEOS_getWKBOutputDims() C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_getWKBOutputDims_r(handle)
}

func cGEOS_setWKBOutputDims(newDims C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_setWKBOutputDims_r(handle, newDims)
}

func cGEOS_getWKBByteOrder() C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_getWKBByteOrder_r(handle)
}

func cGEOS_setWKBByteOrder(byteOrder C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_setWKBByteOrder_r(handle, byteOrder)
}

func cGEOSGeomFromWKB_buf(wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromWKB_buf_r(handle, wkb, size)
}

func cGEOSGeomToWKB_buf(g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToWKB_buf_r(handle, g, size)
}

func cGEOSGeomFromHEX_buf(hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromHEX_buf_r(handle, hex, size)
}

func cGEOSGeomToHEX_buf(g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToHEX_buf_r(handle, g, size)
}

func cGEOSCoordSeq_create(size C.uint, dims C.uint) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_create_r(handle, size, dims)
}

func cGEOSCoordSeq_clone(s *C.GEOSCoordSequence) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_clone_r(handle, s)
}

func cGEOSCoordSeq_destroy(s *C.GEOSCoordSequence) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSCoordSeq_destroy_r(handle, s)
}

func cGEOSCoordSeq_setX(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setX_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setY(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setY_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setZ(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setZ_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setOrdinate(s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setOrdinate_r(handle, s, idx, dim, val)
}

func cGEOSCoordSeq_getX(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getX_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getY(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getY_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getZ(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getZ_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getOrdinate(s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getOrdinate_r(handle, s, idx, dim, val)
}

func cGEOSCoordSeq_getSize(s *C.GEOSCoordSequence, size *C.uint) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getSize_r(handle, s, size)
}

func cGEOSCoordSeq_getDimensions(s *C.GEOSCoordSequence, dims *C.uint) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getDimensions_r(handle, s, dims)
}

func cGEOSProject(g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSProject_r(handle, g, p)
}

func cGEOSInterpolate(g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSInterpolate_r(handle, g, d)
}

func cGEOSProjectNormalized(g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSProjectNormalized_r(handle, g, p)
}

func cGEOSInterpolateNormalized(g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSInterpolateNormalized_r(handle, g, d)
}

func cGEOSBufferParams_create() *C.GEOSBufferParams {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_create_r(handle)
}

func cGEOSBufferParams_destroy(parms *C.GEOSBufferParams) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSBufferParams_destroy_r(handle, parms)
}

func cGEOSBufferParams_setEndCapStyle(p *C.GEOSBufferParams, style C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setEndCapStyle_r(handle, p, style)
}

func cGEOSBufferParams_setJoinStyle(p *C.GEOSBufferParams, joinStyle C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setJoinStyle_r(handle, p, joinStyle)
}

func cGEOSBufferParams_setMitreLimit(p *C.GEOSBufferParams, mitreLimit C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setMitreLimit_r(handle, p, mitreLimit)
}

func cGEOSBufferParams_setQuadrantSegments(p *C.GEOSBufferParams, quadSegs C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setQuadrantSegments_r(handle, p, quadSegs)
}

func cGEOSBufferParams_setSingleSided(p *C.GEOSBufferParams, singleSided C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setSingleSided_r(handle, p, singleSided)
}

func cGEOSBufferWithParams(g1 *C.GEOSGeometry, p *C.GEOSBufferParams, width C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferWithParams_r(handle, g1, p, width)
}

func cGEOSBuffer(g1 *C.GEOSGeometry, width C.double, quadsegs C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBuffer_r(handle, g1, width, quadsegs)
}

func cGEOSBufferWithStyle(g1 *C.GEOSGeometry, width C.double, quadsegs C.int, endCapStyle C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferWithStyle_r(handle, g1, width, quadsegs, endCapStyle, joinStyle, mitreLimit)
}

func cGEOSSingleSidedBuffer(g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double, leftSide C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSingleSidedBuffer_r(handle, g1, width, quadsegs, joinStyle, mitreLimit, leftSide)
}

func cGEOSOffsetCurve(g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSOffsetCurve_r(handle, g1, width, quadsegs, joinStyle, mitreLimit)
}

func cGEOSGeom_createPoint(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createPoint_r(handle, s)
}

func cGEOSGeom_createEmptyPoint() *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyPoint_r(handle)
}

func cGEOSGeom_createLinearRing(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createLinearRing_r(handle, s)
}

func cGEOSGeom_createLineString(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createLineString_r(handle, s)
}

func cGEOSGeom_createEmptyLineString() *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyLineString_r(handle)
}

func cGEOSGeom_createEmptyPolygon() *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyPolygon_r(handle)
}

func cGEOSGeom_createPolygon(shell *C.GEOSGeometry, holes **C.GEOSGeometry, nholes C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createPolygon_r(handle, shell, holes, nholes)
}

func cGEOSGeom_createCollection(_type C.int, geoms **C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createCollection_r(handle, _type, geoms, ngeoms)
}

func cGEOSGeom_createEmptyCollection(_type C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyCollection_r(handle, _type)
}

func cGEOSGeom_clone(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_clone_r(handle, g)
}

func cGEOSGeom_destroy(g *C.GEOSGeometry) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSGeom_destroy_r(handle, g)
}

func cGEOSEnvelope(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSEnvelope_r(handle, g1)
}

func cGEOSIntersection(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSIntersection_r(handle, g1, g2)
}

func cGEOSConvexHull(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSConvexHull_r(handle, g1)
}

func cGEOSDifference(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDifference_r(handle, g1, g2)
}

func cGEOSSymDifference(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSymDifference_r(handle, g1, g2)
}

func cGEOSBoundary(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBoundary_r(handle, g1)
}

func cGEOSUnion(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnion_r(handle, g1, g2)
}

func cGEOSUnaryUnion(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnaryUnion_r(handle, g)
}

func cGEOSUnionCascaded(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnionCascaded_r(handle, g1)
}

func cGEOSPointOnSurface(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPointOnSurface_r(handle, g1)
}

func cGEOSGetCentroid(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetCentroid_r(handle, g)
}

func cGEOSPolygonize(geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonize_r(handle, &geoms[0], ngeoms)
}

func cGEOSPolygonizer_getCutEdges(geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonizer_getCutEdges_r(handle, &geoms[0], ngeoms)
}

func cGEOSPolygonize_full(input *C.GEOSGeometry, cuts **C.GEOSGeometry, dangles **C.GEOSGeometry, invalidRings **C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonize_full_r(handle, input, cuts, dangles, invalidRings)
}

func cGEOSLineMerge(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSLineMerge_r(handle, g)
}

func cGEOSSimplify(g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSimplify_r(handle, g1, tolerance)
}

func cGEOSTopologyPreserveSimplify(g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSTopologyPreserveSimplify_r(handle, g1, tolerance)
}

func cGEOSGeom_extractUniquePoints(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_extractUniquePoints_r(handle, g)
}

func cGEOSSharedPaths(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSharedPaths_r(handle, g1, g2)
}

func cGEOSSnap(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSnap_r(handle, g1, g2, tolerance)
}

func cGEOSDisjoint(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDisjoint_r(handle, g1, g2)
}

func cGEOSTouches(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSTouches_r(handle, g1, g2)
}

func cGEOSIntersects(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSIntersects_r(handle, g1, g2)
}

func cGEOSCrosses(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCrosses_r(handle, g1, g2)
}

func cGEOSWithin(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWithin_r(handle, g1, g2)
}

func cGEOSContains(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSContains_r(handle, g1, g2)
}

func cGEOSOverlaps(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSOverlaps_r(handle, g1, g2)
}

func cGEOSEquals(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSEquals_r(handle, g1, g2)
}

func cGEOSEqualsExact(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSEqualsExact_r(handle, g1, g2, tolerance)
}

func cGEOSCovers(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCovers_r(handle, g1, g2)
}

func cGEOSCoveredBy(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoveredBy_r(handle, g1, g2)
}

func cGEOSPrepare(g *C.GEOSGeometry) *C.GEOSPreparedGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPrepare_r(handle, g)
}

func cGEOSPreparedGeom_destroy(g *C.GEOSPreparedGeometry) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSPreparedGeom_destroy_r(handle, g)
}

func cGEOSPreparedContains(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedContains_r(handle, pg1, g2)
}

func cGEOSPreparedContainsProperly(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedContainsProperly_r(handle, pg1, g2)
}

func cGEOSPreparedCoveredBy(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedCoveredBy_r(handle, pg1, g2)
}

func cGEOSPreparedCovers(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedCovers_r(handle, pg1, g2)
}

func cGEOSPreparedCrosses(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedCrosses_r(handle, pg1, g2)
}

func cGEOSPreparedDisjoint(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedDisjoint_r(handle, pg1, g2)
}

func cGEOSPreparedIntersects(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedIntersects_r(handle, pg1, g2)
}

func cGEOSPreparedOverlaps(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedOverlaps_r(handle, pg1, g2)
}

func cGEOSPreparedTouches(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedTouches_r(handle, pg1, g2)
}

func cGEOSPreparedWithin(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedWithin_r(handle, pg1, g2)
}

func cGEOSSTRtree_create(nodeCapacity C.size_t) *C.GEOSSTRtree {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSTRtree_create_r(handle, nodeCapacity)
}

func cGEOSSTRtree_insert(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_insert_r(handle, tree, g, unsafe.Pointer(item))
}

func cGEOSSTRtree_query(tree *C.GEOSSTRtree, g *C.GEOSGeometry, callback C.GEOSQueryCallback, userdata *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_query_r(handle, tree, g, callback, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_iterate(tree *C.GEOSSTRtree, callback C.GEOSQueryCallback, userdata *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_iterate_r(handle, tree, callback, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_remove(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSTRtree_remove_r(handle, tree, g, unsafe.Pointer(item))
}

func cGEOSSTRtree_destroy(tree *C.GEOSSTRtree) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_destroy_r(handle, tree)
}

func cGEOSisEmpty(g1 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisEmpty_r(handle, g1)
}

func cGEOSisSimple(g1 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisSimple_r(handle, g1)
}

func cGEOSisRing(g1 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisRing_r(handle, g1)
}

func cGEOSHasZ(g1 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSHasZ_r(handle, g1)
}

func cGEOSisClosed(g1 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisClosed_r(handle, g1)
}

func cGEOSRelatePattern(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, pat *C.char) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelatePattern_r(handle, g1, g2, pat)
}

func cGEOSRelate(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelate_r(handle, g1, g2)
}

func cGEOSRelatePatternMatch(mat *C.char, pat *C.char) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelatePatternMatch_r(handle, mat, pat)
}

func cGEOSRelateBoundaryNodeRule(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, bnr C.int) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelateBoundaryNodeRule_r(handle, g1, g2, bnr)
}

func cGEOSisValid(g1 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisValid_r(handle, g1)
}

func cGEOSisValidReason(g1 *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisValidReason_r(handle, g1)
}

func cGEOSisValidDetail(g *C.GEOSGeometry, flags C.int, reason **C.char, location **C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisValidDetail_r(handle, g, flags, reason, location)
}

func cGEOSGeomType(g1 *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomType_r(handle, g1)
}

func cGEOSGeomTypeId(g1 *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomTypeId_r(handle, g1)
}

func cGEOSGetSRID(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetSRID_r(handle, g)
}

func cGEOSSetSRID(g *C.GEOSGeometry, SRID C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSetSRID_r(handle, g, SRID)
}

func cGEOSGetNumGeometries(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetNumGeometries_r(handle, g)
}

func cGEOSGetGeometryN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetGeometryN_r(handle, g, n)
}

func cGEOSNormalize(g1 *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSNormalize_r(handle, g1)
}

func cGEOSGetNumInteriorRings(g1 *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetNumInteriorRings_r(handle, g1)
}

func cGEOSGeomGetNumPoints(g1 *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetNumPoints_r(handle, g1)
}

func cGEOSGeomGetX(g1 *C.GEOSGeometry, x *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetX_r(handle, g1, x)
}

func cGEOSGeomGetY(g1 *C.GEOSGeometry, y *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetY_r(handle, g1, y)
}

func cGEOSGetInteriorRingN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetInteriorRingN_r(handle, g, n)
}

func cGEOSGetExteriorRing(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetExteriorRing_r(handle, g)
}

func cGEOSGetNumCoordinates(g1 *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetNumCoordinates_r(handle, g1)
}

func cGEOSGeom_getCoordSeq(g *C.GEOSGeometry) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getCoordSeq_r(handle, g)
}

func cGEOSGeom_getDimensions(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getDimensions_r(handle, g)
}

func cGEOSGeom_getCoordinateDimension(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getCoordinateDimension_r(handle, g)
}

func cGEOSGeomGetPointN(g1 *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetPointN_r(handle, g1, n)
}

func cGEOSGeomGetStartPoint(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetStartPoint_r(handle, g1)
}

func cGEOSGeomGetEndPoint(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetEndPoint_r(handle, g1)
}

func cGEOSArea(g1 *C.GEOSGeometry, area *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSArea_r(handle, g1, area)
}

func cGEOSLength(g1 *C.GEOSGeometry, length *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSLength_r(handle, g1, length)
}

func cGEOSDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDistance_r(handle, g1, g2, dist)
}

func cGEOSHausdorffDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSHausdorffDistance_r(handle, g1, g2, dist)
}

func cGEOSHausdorffDistanceDensify(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, densifyFrac C.double, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSHausdorffDistanceDensify_r(handle, g1, g2, densifyFrac, dist)
}

func cGEOSGeomGetLength(g1 *C.GEOSGeometry, length *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetLength_r(handle, g1, length)
}

func cGEOSOrientationIndex(Ax C.double, Ay C.double, Bx C.double, By C.double, Px C.double, Py C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSOrientationIndex_r(handle, Ax, Ay, Bx, By, Px, Py)
}

func cGEOSWKTReader_create() *C.GEOSWKTReader {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTReader_create_r(handle)
}

func cGEOSWKTReader_destroy(reader *C.GEOSWKTReader) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTReader_destroy_r(handle, reader)
}

func cGEOSWKTReader_read(reader *C.GEOSWKTReader, wkt *C.char) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTReader_read_r(handle, reader, wkt)
}

func cGEOSWKTWriter_create() *C.GEOSWKTWriter {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTWriter_create_r(handle)
}

func cGEOSWKTWriter_destroy(writer *C.GEOSWKTWriter) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_destroy_r(handle, writer)
}

func cGEOSWKTWriter_write(reader *C.GEOSWKTWriter, g *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTWriter_write_r(handle, reader, g)
}

func cGEOSWKTWriter_setTrim(writer *C.GEOSWKTWriter, trim C.char) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setTrim_r(handle, writer, trim)
}

func cGEOSWKTWriter_setRoundingPrecision(writer *C.GEOSWKTWriter, precision C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setRoundingPrecision_r(handle, writer, precision)
}

func cGEOSWKTWriter_setOutputDimension(writer *C.GEOSWKTWriter, dim C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setOutputDimension_r(handle, writer, dim)
}

func cGEOSWKTWriter_getOutputDimension(writer *C.GEOSWKTWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTWriter_getOutputDimension_r(handle, writer)
}

func cGEOSWKTWriter_setOld3D(writer *C.GEOSWKTWriter, useOld3D C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setOld3D_r(handle, writer, useOld3D)
}

func cGEOSWKBReader_create() *C.GEOSWKBReader {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBReader_create_r(handle)
}

func cGEOSWKBReader_destroy(reader *C.GEOSWKBReader) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBReader_destroy_r(handle, reader)
}

func cGEOSWKBReader_read(reader *C.GEOSWKBReader, wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBReader_read_r(handle, reader, wkb, size)
}

func cGEOSWKBReader_readHEX(reader *C.GEOSWKBReader, hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBReader_readHEX_r(handle, reader, hex, size)
}

func cGEOSWKBWriter_create() *C.GEOSWKBWriter {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_create_r(handle)
}

func cGEOSWKBWriter_destroy(writer *C.GEOSWKBWriter) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_destroy_r(handle, writer)
}

func cGEOSWKBWriter_write(writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_write_r(handle, writer, g, size)
}

func cGEOSWKBWriter_writeHEX(writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_writeHEX_r(handle, writer, g, size)
}

func cGEOSWKBWriter_getOutputDimension(writer *C.GEOSWKBWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getOutputDimension_r(handle, writer)
}

func cGEOSWKBWriter_setOutputDimension(writer *C.GEOSWKBWriter, newDimension C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setOutputDimension_r(handle, writer, newDimension)
}

func cGEOSWKBWriter_getByteOrder(writer *C.GEOSWKBWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getByteOrder_r(handle, writer)
}

func cGEOSWKBWriter_setByteOrder(writer *C.GEOSWKBWriter, byteOrder C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setByteOrder_r(handle, writer, byteOrder)
}

func cGEOSWKBWriter_getIncludeSRID(writer *C.GEOSWKBWriter) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getIncludeSRID_r(handle, writer)
}

func cGEOSWKBWriter_setIncludeSRID(writer *C.GEOSWKBWriter, writeSRID C.char) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setIncludeSRID_r(handle, writer, writeSRID)
}

func cGEOSFree(buffer *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSFree_r(handle, unsafe.Pointer(buffer))
}
