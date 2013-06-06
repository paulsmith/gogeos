package geos

/*
#include "geos.h"
*/
import "C"

// Prepared geometry binary predicates

func cpreparedcontains(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedContains_r(h, g1, g2)
}
func cpreparedcontainsproperly(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedContainsProperly_r(h, g1, g2)
}
func cpreparedcoveredby(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedCoveredBy_r(h, g1, g2)
}
func cpreparedcovers(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedCovers_r(h, g1, g2)
}
func cpreparedcrosses(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedCrosses_r(h, g1, g2)
}
func cprepareddisjoint(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedDisjoint_r(h, g1, g2)
}
func cpreparedintersects(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedIntersects_r(h, g1, g2)
}
func cpreparedoverlaps(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedOverlaps_r(h, g1, g2)
}
func cpreparedtouches(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedTouches_r(h, g1, g2)
}
func cpreparedwithin(h C.GEOSContextHandle_t, g1, g2 *C.GEOSPreparedGeometry) C.char {
	return C.GEOSPreparedWithin_r(h, g1, g2)
}
