package geos

import (
	"bytes"
	"io/ioutil"
	"math"
	"testing"
)

var geomTypeMethodTests = []struct{ wkt, _type string }{
	{"POINT(-117 33)", "Point"},
	{"LINESTRING(10 10, 20 20)", "LineString"},
	{"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))", "Polygon"},
}

func TestGeometryTypeMethod(t *testing.T) {
	for _, test := range geomTypeMethodTests {
		geom := Must(FromWKT(test.wkt))
		actual, err := geom.Type()
		if err != nil {
			panic(err)
		}
		if actual.String() != test._type {
			t.Errorf("Geometry.Type(): want %v, got %v", test._type, actual)
		}
	}
}

var geomTypeTests = []struct {
	_type GeometryType
	out   string
}{
	{POINT, "Point"},
	{MULTILINESTRING, "MultiLineString"},
}

func TestGeometryTypeString(t *testing.T) {
	for _, test := range geomTypeTests {
		if actual := test._type.String(); actual != test.out {
			t.Errorf("GeometryType.String(): want %v, actual %v", test.out, actual)
		}
	}
}

func TestGeometryType(t *testing.T) {
	g := Must(FromWKT("POINT(10 10)"))
	typeID, err := g.Type()
	if err != nil {
		panic(err)
	}
	if typeID != POINT {
		t.Errorf("Geometry.Type(): wanted %v, got %v", POINT, typeID)
	}
}

func TestGeometryProject(t *testing.T) {
	ls := Must(FromWKT("LINESTRING(0 0, 1 1)"))
	pt := Must(FromWKT("POINT(0 1)"))
	expected := 0.7071067811865476
	actual := ls.Project(pt)
	if expected != actual {
		t.Errorf("Geometry.Project(): want %v, got %v", expected, actual)
	}
}

func TestGeometryInterpolate(t *testing.T) {
	g := Must(FromWKT("LINESTRING(0 0, 1 1)"))
	pt := Must(g.Interpolate(0.7071067811865476))
	expected := Must(FromWKT("POINT (0.5 0.5)"))
	ok, err := pt.Equals(expected)
	if err != nil {
		panic(err)
	}
	if !ok {
		t.Errorf("Geometry.Interpolate(): want %v, got %v", expected, pt)
	}
}

func mustEqual(ok bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return ok
}

func TestGeometryBuffer(t *testing.T) {
	g := Must(FromWKT("POINT(0 0)"))
	b := Must(g.Buffer(1.0))
	expected := Must(FromWKT(bufferPoly))
	if !mustEqual(b.EqualsExact(expected, 0.0000001)) {
		t.Errorf("Geometry.Buffer(): want %v, got %v", expected, b)
	}
}

const bufferPoly = `POLYGON ((1.0000000000000000 0.0000000000000000, 0.9807852804032304 -0.1950903220161281, 0.9238795325112870 -0.3826834323650894, 0.8314696123025456 -0.5555702330196017, 0.7071067811865481 -0.7071067811865470, 0.5555702330196032 -0.8314696123025447, 0.3826834323650908 -0.9238795325112863, 0.1950903220161296 -0.9807852804032302, 0.0000000000000016 -1.0000000000000000, -0.1950903220161265 -0.9807852804032308, -0.3826834323650878 -0.9238795325112875, -0.5555702330196004 -0.8314696123025465, -0.7071067811865459 -0.7071067811865492, -0.8314696123025438 -0.5555702330196043, -0.9238795325112857 -0.3826834323650923, -0.9807852804032299 -0.1950903220161312, -1.0000000000000000 -0.0000000000000032, -0.9807852804032311 0.1950903220161249, -0.9238795325112882 0.3826834323650864, -0.8314696123025475 0.5555702330195990, -0.7071067811865505 0.7071067811865446, -0.5555702330196060 0.8314696123025428, -0.3826834323650936 0.9238795325112852, -0.1950903220161322 0.9807852804032297, -0.0000000000000037 1.0000000000000000, 0.1950903220161248 0.9807852804032311, 0.3826834323650867 0.9238795325112881, 0.5555702330195996 0.8314696123025469, 0.7071067811865455 0.7071067811865496, 0.8314696123025438 0.5555702330196044, 0.9238795325112859 0.3826834323650920, 0.9807852804032300 0.1950903220161305, 1.0000000000000000 0.0000000000000000))`

func TestGeometryBufferWithOpts(t *testing.T) {
	g := Must(FromWKT("POINT (0 0)"))
	opts := BufferOpts{QuadSegs: 8, CapStyle: CapRound, JoinStyle: JoinRound, MitreLimit: 5.0}
	b := Must(g.BufferWithOpts(1.0, opts))
	expected := Must(FromWKT(bufferPoly))
	if !mustEqual(b.EqualsExact(expected, 0.000001)) {
		t.Errorf("want %v, got %v", expected, b)
	}
}

func TestOffsetCurve(t *testing.T) {
	g := Must(FromWKT("LINESTRING (0 10, 5 0, 10 10)"))
	opts := BufferOpts{QuadSegs: 8, JoinStyle: JoinRound, MitreLimit: 5.0}
	curve := Must(g.OffsetCurve(1.0, opts))
	expected := Must(FromWKT(offsetCurve))
	if !mustEqual(curve.EqualsExact(expected, 0.000001)) {
		t.Errorf("want %v, got %v", expected, curve)
	}
}

const offsetCurve = `LINESTRING (0.8944271909999159 10.4472135954999583, 5.0000000000000000 2.2360679774997907, 9.1055728090000834 10.4472135954999583)`

func reconstructGeom(g *Geometry) *Geometry {
	typeID, err := g.Type()
	if err != nil {
		panic(err)
	}
	switch typeID {
	case POINT:
		coords := MustCoords(g.Coords())
		return Must(NewPoint(coords...))
	case LINESTRING:
		coords := MustCoords(g.Coords())
		return Must(NewLineString(coords...))
	case LINEARRING:
		coords := MustCoords(g.Coords())
		return Must(NewLinearRing(coords...))
	case POLYGON:
		shell := Must(g.Shell())
		shellCoords := MustCoords(shell.Coords())
		holes, err := g.Holes()
		if err != nil {
			panic(err)
		}
		holesCoords := make([][]Coord, len(holes))
		for i, ring := range holes {
			holesCoords[i] = MustCoords(ring.Coords())
		}
		return Must(NewPolygon(shellCoords, holesCoords...))
	case MULTIPOINT, MULTILINESTRING, MULTIPOLYGON, GEOMETRYCOLLECTION:
		n, err := g.NGeometry()
		if err != nil {
			panic(err)
		}
		var geoms []*Geometry
		for i := 0; i < n; i++ {
			geom := Must(g.Geometry(i))
			geoms = append(geoms, reconstructGeom(geom))
		}
		return Must(NewCollection(typeID, geoms...))
	}
	return nil
}

func TestGeomConstructors(t *testing.T) {
	const filename = `./testdata/test.wkt`
	wkt, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	g1 := Must(FromWKT(string(wkt)))
	g2 := reconstructGeom(g1)
	if !mustEqual(g1.Equals(g2)) {
		t.Errorf("Fine-grained geometry reconstruction failed")
	}
}

func TestArea(t *testing.T) {
	g1 := Must(FromWKT("POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))"))
	expected := 4.0
	area, err := g1.Area()
	if err != nil {
		panic(err)
	}
	if area != expected {
		t.Errorf("Area(): want %v got %v", expected, area)
	}
}

func TestLength(t *testing.T) {
	g1 := Must(FromWKT("LINESTRING(0 0, 1 1)"))
	expected := 1.4142135623730951
	l, err := g1.Length()
	if err != nil {
		panic(err)
	}
	if l != expected {
		t.Errorf("Length(): want %v got %v", expected, l)
	}
}

func TestDistance(t *testing.T) {
	g1 := Must(FromWKT("POINT(0 0)"))
	g2 := Must(FromWKT("POINT(1 1)"))
	expected := math.Sqrt(2)
	d, err := g1.Distance(g2)
	if err != nil {
		panic(err)
	}
	if d != expected {
		t.Errorf("Distance(): want %v got %v", expected, d)
	}
}

func TestLineStringPointFns(t *testing.T) {
	ls := Must(FromWKT(`LINESTRING(-117 35, 42 -12, 55 3, 100 100, 0 715)`))
	g1 := Must(FromWKT("POINT(-117 35)"))
	g2 := Must(FromWKT("POINT(55 3)"))
	g3 := Must(FromWKT("POINT(0 715)"))
	if pt := Must(ls.StartPoint()); !mustEqual(pt.Equals(g1)) {
		t.Errorf("StartPoint(): should equal %v! got %v", g1, pt)
	}
	if pt := Must(ls.EndPoint()); !mustEqual(pt.Equals(g3)) {
		t.Errorf("EndPoint(): should equal %v! got %v", g3, pt)
	}
	if pt := Must(ls.Point(2)); !mustEqual(pt.Equals(g2)) {
		t.Errorf("Point(n): should equal %v! got %v", g2, pt)
	}
}

func TestCoordDim(t *testing.T) {
	tests := []struct {
		wkt string
		dim int
	}{
		{"POINT(1 3)", 2},
		{"POINT(2 4 6)", 3},
	}
	for _, test := range tests {
		if dim := Must(FromWKT(test.wkt)).CoordDimension(); dim != test.dim {
			t.Errorf("CoordDimension(): want %v, got %v", test.dim, dim)
		}
	}
}

func TestDimension(t *testing.T) {
	tests := []struct {
		wkt string
		dim int
	}{
		{"POINT(-117 35)", 0},
		{"LINESTRING(63 -41, 87 -40)", 1},
		{"POLYGON((0 0, 0 1, 1 1, 0 0))", 2},
	}

	for _, test := range tests {
		g := Must(FromWKT(test.wkt))
		if dim := g.Dimension(); dim != test.dim {
			t.Errorf("Dimension(): want %v got %v", test.dim, dim)
		}
	}
}

func TestNCoordinate(t *testing.T) {
	tests := []struct {
		wkt string
		n   int
	}{
		{"POINT EMPTY", 0},
		{"POINT(-117 35)", 1},
		{"LINESTRING(-117 35, 0 0, 1 1)", 3},
		{"POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))", 5},
		{"POLYGON((0 0, 0 2, 2 2, 2 0, 0 0), (1 1, 1 1.5, 1.5 1.5, 1.5 1, 1 1))", 10},
	}
	for _, test := range tests {
		n, err := Must(FromWKT(test.wkt)).NCoordinate()
		if err != nil {
			panic(err)
		}
		if n != test.n {
			t.Errorf("NCoordinate(): want %v, got %v", test.n, n)
		}
	}
}

func TestShell(t *testing.T) {
	g := Must(FromWKT("POLYGON((0 0, 0 2, 2 2, 2 0, 0 0), (1 1, 1 1.5, 1.5 1.5, 1.5 1, 1 1))"))
	ext := Must(g.Shell())
	expected := Must(FromWKT("LINEARRING(0 0, 0 2, 2 2, 2 0, 0 0)"))
	if !mustEqual(ext.Equals(expected)) {
		t.Errorf("Shell(): should equal! got %v", ext)
	}
}

func TestHoles(t *testing.T) {
	poly := Must(FromWKT(`POLYGON((0 0, 0 6, 6 6, 6 0, 0 0),
                                  (1 1, 2 1, 2 2, 1 2, 1 1),
                                  (1 3, 2 3, 2 4, 1 4, 1 3),
                                  (3 2, 4 2, 4 3, 3 3, 3 2))`))
	tests := [][]string{
		{
			"LINEARRING(1 1, 2 1, 2 2, 1 2, 1 1)",
			"LINEARRING(1 3, 2 3, 2 4, 1 4, 1 3)",
			"LINEARRING(3 2, 4 2, 4 3, 3 3, 3 2)",
		},
	}
	for i, holeWkts := range tests {
		holes, err := poly.Holes()
		if err != nil {
			t.Fatalf("#%d: %v", i, err)
		}
		if len(holes) != len(holeWkts) {
			t.Errorf("#%d: want %d holes, got %d", i, len(holeWkts), len(holes))
		}
		for j, wkt := range holeWkts {
			ring := Must(FromWKT(wkt))
			if !mustEqual(holes[j].Equals(ring)) {
				t.Errorf("#%d: want int ring to equal! got %v", i, holes[j])
			}
		}
	}
}

func TestX(t *testing.T) {
	g := Must(FromWKT("POINT(-117 35)"))
	x, err := g.X()
	if err != nil {
		panic(err)
	}
	var expected float64 = -117
	if x != expected {
		t.Errorf("X(): want %v got %v", expected, x)
	}
}

func TestY(t *testing.T) {
	g := Must(FromWKT("POINT(-117 35)"))
	y, err := g.Y()
	if err != nil {
		panic(err)
	}
	var expected float64 = 35
	if y != expected {
		t.Errorf("Y(): want %v got %v", expected, y)
	}
}

func TestNPoint(t *testing.T) {
	g := Must(FromWKT("LINESTRING(0 0, 1 1)"))
	n, err := g.NPoint()
	if err != nil {
		panic(err)
	}
	if n != 2 {
		t.Errorf("NPoint(): want 2 got %v", n)
	}
}

func mustInt(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

func TestNormalize(t *testing.T) {
	g1 := Must(FromWKT("MULTIPOINT((46 27), (61 79), (92 8), (17 7), (33 44))"))
	if err := g1.Normalize(); err != nil {
		t.Errorf("Normalize(): error: %v", err)
	}
	g2 := Must(FromWKT("MULTIPOINT (92 8, 61 79, 46 27, 33 44, 17 7)"))
	if !mustEqual(g1.EqualsExact(g2, 0.0)) {
		t.Errorf("Normalize(): want %v got %v", g2, g1)
	}
}

func TestGeometry(t *testing.T) {
	tests := []struct {
		multi, geom string
		n           int
	}{
		{"MULTIPOINT((0 0), (1 1), (2 2))", "POINT(0 0)", 0},
		{"MULTIPOINT((0 0), (1 1), (2 2))", "POINT(1 1)", 1},
	}
	for _, test := range tests {
		g1 := Must(FromWKT(test.multi))
		g2 := Must(FromWKT(test.geom))
		if g3 := Must(g1.Geometry(test.n)); !mustEqual(g3.Equals(g2)) {
			t.Errorf("Geometry(%v): should equal! got %v", test.n, g3)
		}
	}
}

func TestNGeometry(t *testing.T) {
	tests := []struct {
		wkt string
		n   int
	}{
		{"MULTIPOINT((0 0), (1 1), (2 2))", 3},
	}
	for _, test := range tests {
		g := Must(FromWKT(test.wkt))
		if n := mustInt(g.NGeometry()); n != test.n {
			t.Errorf("NGeometry(): want %v got %v", test.n, n)
		}
	}
}

func TestSRID(t *testing.T) {
	g := Must(FromWKT("POINT(-117 35)"))
	if _, err := g.SRID(); err == nil {
		t.Errorf("SRID(): should be error on unset SRID")
	}
	srid := 4326
	g.SetSRID(srid)
	if actual := mustInt(g.SRID()); actual != srid {
		t.Errorf("SRID(): want %v got %v", srid, actual)
	}
}

func TestIsClosed(t *testing.T) {
	tests := []struct {
		wkt     string
		_closed bool
	}{
		{"LINEARRING(0 0, 1 1, 0 1, 0 0)", true},
		{"LINESTRING(0 0, 1 1, 0 1, 0 0)", true},
		{"LINESTRING(0 0, 1 1, 0 1)", false},
	}
	for _, test := range tests {
		g := Must(FromWKT(test.wkt))
		if _closed := mustBool(g.IsClosed()); _closed != test._closed {
			t.Errorf("IsClosed(): %v - want %v got %v", g, test._closed, _closed)
		}
	}
}

func mustBool(b bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return b
}

func TestHasZ(t *testing.T) {
	tests := []struct {
		wkt string
		z   bool
	}{
		{"POINT(0 0)", false},
		{"POINT(0 0 0)", true},
	}
	for _, test := range tests {
		g := Must(FromWKT(test.wkt))
		if z := mustBool(g.HasZ()); z != test.z {
			t.Errorf("HasZ(): %v - want %v got %v", g, test.z, z)
		}
	}
}

//func TestIsRing ...

//func TestIsSimple ...

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		wkt   string
		empty bool
	}{
		{"POINT EMPTY", true},
		{"POINT(-117 35)", false},
	}
	for _, test := range tests {
		g := Must(FromWKT(test.wkt))
		if empty := mustBool(g.IsEmpty()); empty != test.empty {
			t.Errorf("IsEmpty(): %v - want %v got %v", g, test.empty, empty)
		}
	}
}

var binaryTopoTests = []struct {
	g1, g2, out string
	method      func(*Geometry, *Geometry) (*Geometry, error)
}{
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 1 3, 3 3, 3 1, 1 1))",
		"POLYGON((1 1, 2 1, 2 2, 1 2, 1 1))",
		(*Geometry).Intersection,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 1 3, 3 3, 3 1, 1 1))",
		"POLYGON((2 1, 2 0, 0 0, 0 2, 1 2, 1 1, 2 1))",
		(*Geometry).Difference,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 1 3, 3 3, 3 1, 1 1))",
		"MULTIPOLYGON(((2 1, 2 0, 0 0, 0 2, 1 2, 1 1, 2 1)), ((2 1, 2 2, 1 2, 1 3, 3 3, 3 1, 2 1)))",
		(*Geometry).SymDifference,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 1 3, 3 3, 3 1, 1 1))",
		"POLYGON((2 1, 2 0, 0 0, 0 2, 1 2, 1 3, 3 3, 3 1, 2 1))",
		(*Geometry).Union,
	},
	{
		"LINESTRING(0 1, 1 1, 2 2, 3 3, 4 4, 4 5)",
		"LINESTRING(1 0, 1 1, 4 4, 5 4)",
		"GEOMETRYCOLLECTION (MULTILINESTRING ((1 1, 2 2), (2 2, 3 3), (3 3, 4 4)), MULTILINESTRING EMPTY)",
		(*Geometry).SharedPaths,
	},
}

func TestBinaryTopo(t *testing.T) {
	for _, test := range binaryTopoTests {
		g1 := Must(FromWKT(test.g1))
		g2 := Must(FromWKT(test.g2))
		expected := Must(FromWKT(test.out))
		if actual := Must(test.method(g1, g2)); !mustEqual(expected.Equals(actual)) {
			t.Errorf("%+V(): want %v got %v", test.method, expected, actual)
		}
	}
}

func TestSnap(t *testing.T) {
	tests := []struct {
		g1, g2, out string
		tol         float64
	}{
		{
			"POINT(0.05 0.05)",
			"POINT(0 0)",
			"POINT(0 0)",
			0.1,
		},
	}
	for _, test := range tests {
		g1 := Must(FromWKT(test.g1))
		g2 := Must(FromWKT(test.g2))
		expected := Must(FromWKT(test.out))
		if actual := Must(g1.Snap(g2, test.tol)); !mustEqual(expected.Equals(actual)) {
			t.Errorf("Snap(%v): want %v got %v", test.tol, expected, actual)
		}
	}
}

var unaryTopoTests = []struct {
	g1, out string
	method  func(*Geometry) (*Geometry, error)
}{
	{
		"MULTIPOINT((3 1.5), (3.5 1), (4 1), (5 2), (4 1.5), (3.5 1.5))",
		"POLYGON ((3 1, 5 1, 5 2, 3 2, 3 1))",
		(*Geometry).Envelope,
	},
	{
		"POLYGON((1 1, 3 1, 2 2, 3 3, 1 3, 1 1))",
		"POLYGON ((1 1, 1 3, 3 3, 3 1, 1 1))",
		(*Geometry).ConvexHull,
	},
	/*
	   {
	       "POINT(-117 35)",
	       "GEOMETRYCOLLECTION EMPTY", // XXX can't compare empty geoms for equality
	       (*Geometry).Boundary,
	   },
	*/
	{
		"LINESTRING(0 0, 5 5, 10 0)",
		"MULTIPOINT (0 0, 10 0)",
		(*Geometry).Boundary,
	},
	{
		"POLYGON((1 1, 3 1, 2 2, 3 3, 1 3, 1 1))",
		"LINESTRING (1 1, 3 1, 2 2, 3 3, 1 3, 1 1)",
		(*Geometry).Boundary,
	},
	{
		"MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))",
		"POLYGON ((10 5, 10 0, 0 0, 0 10, 5 10, 5 15, 15 15, 15 5, 10 5))",
		(*Geometry).UnaryUnion,
	},
	{
		"MULTIPOINT((0 0), (1 1), (0 0), (2 2), (-117 35), (2 2))",
		"MULTIPOINT (-117 35, 0 0, 1 1, 2 2)",
		(*Geometry).UnaryUnion,
	},
	{
		"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
		"POINT (0.5 0.5)",
		(*Geometry).PointOnSurface,
	},
	{
		"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
		"POINT (0.5 0.5)",
		(*Geometry).Centroid,
	},
	{
		"MULTILINESTRING((1 5, 3 4, 1 1), (1 1, 2 0, 3 1))",
		"LINESTRING (1 5, 3 4, 1 1, 2 0, 3 1)",
		(*Geometry).LineMerge,
	},
	{
		"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
		"MULTIPOINT (0 0, 1 0, 1 1, 0 1)",
		(*Geometry).UniquePoints,
	},
}

func TestUnaryTopo(t *testing.T) {
	for _, test := range unaryTopoTests {
		g1 := Must(FromWKT(test.g1))
		expected := Must(FromWKT(test.out))
		if actual := Must(test.method(g1)); !mustEqual(actual.EqualsExact(expected, 0.0)) {
			t.Errorf("%+V(): want %v got %v", test.method, expected, actual)
		}
	}
}

func TestSimplifyMethods(t *testing.T) {
	tests := []struct {
		g1, out string
		tol     float64
		method  func(*Geometry, float64) (*Geometry, error)
	}{
		{
			"LINESTRING(0 0, 1 1, 0 2, 1 3, 0 4, 1 5)",
			"LINESTRING (0 0, 1 5)",
			1.0,
			(*Geometry).Simplify,
		},
		{
			"LINESTRING(0 0, 1 1, 0 2, 1 3, 0 4, 1 5)",
			"LINESTRING (0 0, 1 5)",
			1.0,
			(*Geometry).SimplifyP,
		},
		// XXX: geom that would collapse and testing for validity/simplicity
	}
	for _, test := range tests {
		g1 := Must(FromWKT(test.g1))
		expected := Must(FromWKT(test.out))
		if actual := Must(test.method(g1, test.tol)); !mustEqual(actual.EqualsExact(expected, 0.0)) {
			t.Errorf("%+V(): want %v got %v", test.method, expected, actual)
		}
	}
}

var binaryPredTests = []struct {
	g1, g2 string
	pred   bool
	method func(*Geometry, *Geometry) (bool, error)
}{
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 3 1, 3 3, 1 3, 1 1))",
		false,
		(*Geometry).Disjoint,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((3 3, 5 3, 5 5, 3 5, 3 3))",
		true,
		(*Geometry).Disjoint,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"LINESTRING(1 2, 3 2)",
		true,
		(*Geometry).Touches,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"LINESTRING(5 2, 6 2)",
		false,
		(*Geometry).Touches,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 3 1, 3 3, 1 3, 1 1))",
		true,
		(*Geometry).Intersects,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((3 3, 5 3, 5 5, 3 5, 3 3))",
		false,
		(*Geometry).Intersects,
	},
	{
		"LINESTRING(0 0, 10 10)",
		"LINESTRING(10 0, 0 10)",
		true,
		(*Geometry).Crosses,
	},
	{
		"LINESTRING(0 0, 10 10)",
		"LINESTRING(11 0, 11 10)",
		false,
		(*Geometry).Crosses,
	},
	{
		"LINESTRING(0 0, 10 10)",
		"POLYGON((-5 -5, 5 -5, 5 5, -5 5, -5 -5))",
		true,
		(*Geometry).Crosses,
	},
	{
		"POINT(3 3)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		true,
		(*Geometry).Within,
	},
	{
		"POINT(-1 35)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		false,
		(*Geometry).Within,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(3 3)",
		true,
		(*Geometry).Contains,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(-1 35)",
		false,
		(*Geometry).Contains,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((1 1, 3 1, 3 3, 1 3, 1 1))",
		true,
		(*Geometry).Overlaps,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((3 3, 5 3, 5 5, 3 5, 3 3))",
		false,
		(*Geometry).Overlaps,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((0 0, 0 2, 2 2, 2 0, 0 0))",
		true,
		(*Geometry).Equals,
	},
	{
		"POLYGON((0 0, 2 0, 2 2, 0 2, 0 0))",
		"POLYGON((0 0, 0 2, 2 2.1, 2 0, 0 0))",
		false,
		(*Geometry).Equals,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(3 3)",
		true,
		(*Geometry).Covers,
	},
	{
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		"POINT(-1 35)",
		false,
		(*Geometry).Covers,
	},
	{
		"POINT(3 3)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		true,
		(*Geometry).CoveredBy,
	},
	{
		"POINT(-1 35)",
		"POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))",
		false,
		(*Geometry).CoveredBy,
	},
}

func TestBinaryPred(t *testing.T) {
	for _, test := range binaryPredTests {
		g1 := Must(FromWKT(test.g1))
		g2 := Must(FromWKT(test.g2))
		if actual := mustBool(test.method(g1, g2)); actual != test.pred {
			t.Errorf("%+V(): want %v got %v", test.method, test.pred, actual)
		}
	}
}

func TestEqualsExact(t *testing.T) {
	tests := []struct {
		g1, g2 string
		tol    float64
		pred   bool
	}{
		{
			"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
			"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
			0.0,
			true,
		},
		{
			"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
			"POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))",
			0.0,
			false,
		},
		{
			"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
			"POLYGON((0.05 0.05, 1.05 0.05, 1.05 1.05, 0.05 1.05, 0.05 0.05))",
			0.1,
			true,
		},
		{
			"POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))",
			"POLYGON((0.05 0.05, 0.05 1.05, 1.05 1.05, 1.05 0.05, 0.05 0.05))",
			0.1,
			false,
		},
	}
	for _, test := range tests {
		g1 := Must(FromWKT(test.g1))
		g2 := Must(FromWKT(test.g2))
		if actual := mustBool(g1.EqualsExact(g2, test.tol)); actual != test.pred {
			t.Errorf("EqualsExact(): want %v got %v", test.pred, actual)
		}
	}
}

func TestClone(t *testing.T) {
	g1 := Must(FromWKT("POINT(-117 35)"))
	g2 := Must(g1.Clone())
	if !mustEqual(g1.EqualsExact(g2, 0.0)) {
		t.Errorf("Cloned geom should equal! %v != %v", g1, g2)
	}
	if g1.g == g2.g {
		t.Errorf("Cloned geom's C ptrs should not equal!")
	}
}

// Constructors

var basicConstructorTests = []struct {
	coords []Coord
	ctor   func(...Coord) (*Geometry, error)
	err    bool
	empty  bool
}{
	{nil, NewPoint, false, true},
	{[]Coord{NewCoord(-117, 35)}, NewPoint, false, false},
	{[]Coord{NewCoord(-117, 35), NewCoord(0, 0)}, NewPoint, true, false},
	{nil, NewLineString, false, true},
	{[]Coord{NewCoord(0, 0), NewCoord(10, 10), NewCoord(20, 20)}, NewLineString, false, false},
	{nil, NewLinearRing, false, true},
	{[]Coord{NewCoord(0, 0), NewCoord(10, 10), NewCoord(10, 0), NewCoord(0, 0)}, NewLinearRing, false, false},
	{[]Coord{NewCoord(0, 0)}, NewLinearRing, true, false},
	{[]Coord{NewCoord(0, 0), NewCoord(10, 10), NewCoord(0, 0)}, NewLinearRing, true, false},
}

func TestConstructors(t *testing.T) {
	for i, test := range basicConstructorTests {
		geom, err := test.ctor(test.coords...)
		if err != nil {
			if !test.err {
				t.Errorf("#%d: ctor: want no error, got: %v", i, err)
			}
			continue
		}
		empty, err := geom.IsEmpty()
		if err != nil {
			t.Errorf("#%d: empty error: %v", i, err)
			continue
		}
		if empty != test.empty {
			t.Errorf("#%d: empty: want %v, got %v", i, test.empty, empty)
		}
	}
}

var polygonConstructorTests = []struct {
	shell []Coord
	holes [][]Coord
	err   bool
	empty bool
}{
	{nil, nil, false, true},
	{[]Coord{NewCoord(0, 0), NewCoord(10, 10), NewCoord(10, 0), NewCoord(0, 0)}, nil, false, false},
	{
		[]Coord{NewCoord(0, 0), NewCoord(10, 10), NewCoord(10, 0), NewCoord(0, 0)},
		[][]Coord{[]Coord{NewCoord(2, 1), NewCoord(2, 2), NewCoord(3, 1), NewCoord(2, 1)}},
		false,
		false,
	},
	{[]Coord{NewCoord(0, 0), NewCoord(10, 10), NewCoord(10, 0)}, nil, true, false},
}

func TestPolygonConstructor(t *testing.T) {
	for i, test := range polygonConstructorTests {
		geom, err := NewPolygon(test.shell, test.holes...)
		if err != nil {
			if !test.err {
				t.Errorf("#%d: ctor: want no error, got: %v", i, err)
			}
			continue
		}
		empty, err := geom.IsEmpty()
		if err != nil {
			t.Errorf("#%d: empty error: %v", i, err)
			continue
		}
		if empty != test.empty {
			t.Errorf("#%d: empty: want %v, got %v", i, test.empty, empty)
		}
	}
}

func TestLineStringLinearRingEqual(t *testing.T) {
	line := Must(FromWKT("LINESTRING (0 0, 10 10, 10 0, 0 0)"))
	ring := Must(FromWKT("LINEARRING (0 0, 10 10, 10 0, 0 0)"))
	if !mustEqual(line.Equals(ring)) {
		t.Errorf("expected equal!")
	}
}

var relateTests = []struct {
	g1, g2 string
	pat    string
}{
	{
		"POLYGON ((60 160, 220 160, 220 20, 60 20, 60 160))",
		"POLYGON ((60 160, 20 200, 260 200, 140 80, 60 160))",
		"212101212",
	},
}

func TestRelate(t *testing.T) {
	for i, test := range relateTests {
		g1 := Must(FromWKT(test.g1))
		g2 := Must(FromWKT(test.g2))
		pat, err := g1.Relate(g2)
		if err != nil {
			t.Fatalf("#%d %v", i, err)
		}
		if pat != test.pat {
			t.Errorf("#%d want %v got %v", i, test.pat, pat)
		}
	}
}

var relatePatTests = []struct {
	g1, g2 string
	pat    string
	relate bool
}{
	{
		"POLYGON ((60 160, 220 160, 220 20, 60 20, 60 160))",
		"POLYGON ((60 160, 20 200, 260 200, 140 80, 60 160))",
		"212101212",
		true,
	},
}

func TestRelatePat(t *testing.T) {
	for i, test := range relatePatTests {
		g1 := Must(FromWKT(test.g1))
		g2 := Must(FromWKT(test.g2))
		ok, err := g1.RelatePat(g2, test.pat)
		if err != nil {
			t.Fatalf("#%d %v", i, err)
		}
		if ok != test.relate {
			t.Errorf("#%d want %v got %v", i, test.relate, ok)
		}
	}
}

func TestFromWKB(t *testing.T) {
	for i, test := range wkbDecoderTests {
		g1 := Must(FromWKB(test.wkb))
		g2 := Must(FromWKT(test.wkt))
		if !mustEqual(g1.Equals(g2)) {
			t.Errorf("#%d want %v got %v", i, test.wkt, g1.String())
		}
	}
}

func TestFromHex(t *testing.T) {
	for i, test := range wkbDecoderHexTests {
		g1 := Must(FromHex(test.hex))
		g2 := Must(FromWKT(test.wkt))
		if !mustEqual(g1.Equals(g2)) {
			t.Errorf("#%d want %v got %v", i, test.wkt, g1.String())
		}
	}
}

func TestWKB(t *testing.T) {
	for i, test := range wkbEncoderTests {
		g := Must(FromWKT(test.wkt))
		wkb, err := g.WKB()
		if err != nil {
			t.Fatalf("#%d %v", i, err)
		}
		if !bytes.Equal(wkb, test.wkb) {
			t.Errorf("#%d want %v got %v", test.wkb, wkb)
		}
	}
}

func TestHex(t *testing.T) {
	for i, test := range wkbEncoderHexTests {
		g := Must(FromWKT(test.wkt))
		hex, err := g.Hex()
		if err != nil {
			t.Fatalf("#%d %v", i, err)
		}
		if !bytes.Equal(hex, test.wkb) {
			t.Errorf("#%d want %v got %v", string(test.wkb), string(hex))
		}
	}
}

func TestLineInterpolatePointDistError(t *testing.T) {
	line := Must(FromWKT("LINESTRING(0 0, 10 10)"))
	_, err := line.LineInterpolatePoint(-0.1)
	if err != ErrLineInterpolatePointDist {
		t.Errorf("must not allow negative distance")
	}
	_, err = line.LineInterpolatePoint(1.1)
	if err != ErrLineInterpolatePointDist {
		t.Errorf("must not allow distance greater than 1.0")
	}
}

func TestLineInterpolatePointTypeError(t *testing.T) {
	pt := Must(FromWKT("POINT(0 0)"))
	_, err := pt.LineInterpolatePoint(0.0)
	if err != ErrLineInterpolatePointType {
		t.Errorf("only permitted on linestrings")
	}
}

func TestLineInterpolatePoint(t *testing.T) {
	var tests = []struct {
		line string
		dist float64
		pt   string
	}{
		{"LINESTRING(25 50, 75 75, 100 35)", 0.0, "POINT(25 50)"},
		{"LINESTRING(25 50, 75 75, 100 35)", 1.0, "POINT(100 35)"},
		{"LINESTRING(25 50, 100 125, 150 190)", 0.2, "POINT (51.5974135047432014 76.5974135047432014)"},
	}
	for i, test := range tests {
		line := Must(FromWKT(test.line))
		actual, err := line.LineInterpolatePoint(test.dist)
		if err != nil {
			t.Fatalf("#%d %s", i, err)
		}
		expected := Must(FromWKT(test.pt))
		if !mustEqual(actual.Equals(expected)) {
			t.Errorf("#%d want %v got %v", i, test.pt, actual.String())
		}
	}
}
