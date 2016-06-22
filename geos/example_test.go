package geos_test

import (
	"fmt"

	"github.com/robertogyn19/gogeos/geos"
)

func ExampleGeometry_LineInterpolatePoint() {
	line := geos.Must(geos.FromWKT("LINESTRING(25 50, 100 125, 150 190)"))
	pt := geos.Must(line.LineInterpolatePoint(0.20))
	fmt.Println(pt)
	// Output: POINT (51.5974135047432014 76.5974135047432014)
}
