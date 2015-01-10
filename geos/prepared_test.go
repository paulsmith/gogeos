package geos

import (
	"testing"
)

var preparedPredTests = []struct {
	g1, g2 string
	pred   bool
	method func(*PGeometry, *Geometry) (bool, error)
}{
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 3 1, 3 3, 1 3, 1 1))",
		false,
		(*PGeometry).Disjoint,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((3 3, 5 3, 5 5, 3 5, 3 3))",
		true,
		(*PGeometry).Disjoint,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"LINESTRING(1 2, 3 2)",
		true,
		(*PGeometry).Touches,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"LINESTRING(5 2, 6 2)",
		false,
		(*PGeometry).Touches,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 3 1, 3 3, 1 3, 1 1))",
		true,
		(*PGeometry).Intersects,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((3 3, 5 3, 5 5, 3 5, 3 3))",
		false,
		(*PGeometry).Intersects,
	},
	{
		"LINESTRING(0 0, 10 10)",
		"LINESTRING(10 0, 0 10)",
		true,
		(*PGeometry).Crosses,
	},
	{
		"LINESTRING(0 0, 10 10)",
		"LINESTRING(11 0, 11 10)",
		false,
		(*PGeometry).Crosses,
	},
	{
		"LINESTRING(0 0, 10 10)",
		"POLYGON((-5 -5, 5 -5, 5 5, -5 5, -5 -5))",
		true,
		(*PGeometry).Crosses,
	},
	{
		"POINT(3 3)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		true,
		(*PGeometry).Within,
	},
	{
		"POINT(-1 35)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		false,
		(*PGeometry).Within,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(3 3)",
		true,
		(*PGeometry).Contains,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(-1 35)",
		false,
		(*PGeometry).Contains,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 3 1, 3 3, 1 3, 1 1))",
		true,
		(*PGeometry).Overlaps,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((3 3, 5 3, 5 5, 3 5, 3 3))",
		false,
		(*PGeometry).Overlaps,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(3 3)",
		true,
		(*PGeometry).Covers,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(-1 35)",
		false,
		(*PGeometry).Covers,
	},
	{
		"POINT(3 3)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		true,
		(*PGeometry).CoveredBy,
	},
	{
		"POINT(-1 35)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		false,
		(*PGeometry).CoveredBy,
	},
}

func TestPreparedBinaryPred(t *testing.T) {
	for _, test := range preparedPredTests {
		g1 := Must(FromWKT(test.g1)).Prepare()
		g2 := Must(FromWKT(test.g2))
		if actual := mustBool(test.method(g1, g2)); actual != test.pred {
			t.Errorf("%+V(): want %v got %v", test.method, test.pred, actual)
		}
	}
}
