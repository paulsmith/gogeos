// +build ignore

package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"code.google.com/p/draw2d/draw2d"

	"github.com/paulsmith/gogeos/geos"
)

const (
	WIDTH  = 512
	HEIGHT = 512
	PAD    = 20
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
}

func ex1() {
	example("example1.png", geos.Must(geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")))
}

var (
	a = geos.Must(geos.FromWKT("POLYGON ((0 3, 2 3, 3 1, 1 0, 0 1.5, 0 3))"))
	b = geos.Must(geos.FromWKT("POLYGON ((1 2, 1.5 4, 3.5 4, 4.5 2.5, 1 2))"))
)

func ex2() {
	a := geos.Must(geos.FromWKT("LINEARRING (0 3, 2 3, 3 1, 1 0, 0 1.5, 0 3)"))
	b := geos.Must(geos.FromWKT("LINEARRING (1 2, 1.5 4, 3.5 4, 4.5 2.5, 1 2)"))
	example("example2-a-b.png", a, b)
}

func ex3() {
	example("example3-intersection.png", a, b, geos.Must(a.Intersection(b)))
}

func ex4() {
	example("example4-union.png", a, b, geos.Must(a.Union(b)))
}

func ex5() {
	example("example5-difference.png", a, b, geos.Must(a.Difference(b)))
}

func ex6() {
	example("example6-difference.png", a, b, geos.Must(b.Difference(a)))
}

func ex7() {
	example("example7-symdifference.png", geos.Must(a.SymDifference(b)))
}

func ex8() {
	l := geos.Must(geos.FromWKT("LINESTRING (0 2, 6 2, 3 4, 4.5 8, 1 8, 3 0)"))
	b := geos.Must(l.Buffer(0.5))
	example("example8-buffer.png", b, l)
}

func ex9() {
	wkt := []string{
		"LINESTRING (2 10, 3 10, 4 9, 5 8)",
		"LINESTRING (5 8, 6 6, 6 5)",
		"LINESTRING (6 5, 9 6, 10 7, 11 9)",
		"LINESTRING (11 9, 11 10, 10 11)",
		"LINESTRING (6 5, 3 4, 2 3)",
		"LINESTRING (2 3, 1.5 2.5, 1 1, 1 0)",
		"LINESTRING (1 0, 0 0, 0 -1, 1 -2)",
		"LINESTRING (1 -2, 2 -2, 3 -1, 3 0, 1 0)",
		"LINESTRING (6 5, 6 3, 6.5 2.5)",
		"LINESTRING (6.5 2.5, 7.5 2, 8.5 1.5)",
		"LINESTRING (8.5 1.5, 9 0.5, 10 0.5, 11 1.5)",
	}
	var linestrings []*geos.Geometry
	for i := range wkt {
		linestrings = append(linestrings, geos.Must(geos.FromWKT(wkt[i])))
	}
	example("example9-unmerged-linestrings.png", linestrings...)
	merged := geos.Must(geos.Must(geos.NewCollection(geos.MULTILINESTRING, linestrings...)).LineMerge())
	example("example9-merged-linestrings.png", merged)
}

func ex10() {
	l := geos.Must(geos.FromWKT("LINESTRING (0 2, 6 2, 3 4, 4.5 8, 1 8, 3 0)"))
	b := geos.Must(l.ConvexHull())
	example("example10-convex-hull.png", b, l)
}

func example(filename string, geoms ...*geos.Geometry) {
	m, ctxt := newContext(WIDTH, HEIGHT)
	ctxt.Clear()
	drawGeoms(m, ctxt, geoms...)
	saveImageToFile(m, filename)
}

func test() {
	g1 := geos.Must(geos.FromWKT("LINESTRING(0 0, 3 3, 3 2, 4 2, 7 5)"))
	g2 := geos.Must(geos.FromWKT("LINESTRING(0 10, 1 0, 0 10, 10 10)"))
	g3 := geos.Must(geos.FromWKT("POINT(5 5)"))
	g4 := geos.Must(g3.Buffer(1))
	g5 := geos.Must(geos.FromWKT("MULTIPOINT(5 6, 6 6, 3 8, 7 8)"))
	example("test.png", g1, g2, g4, g3, g5)
}

var (
	blue    = color.RGBA{0x4C, 0x94, 0xFF, 0xFF}
	orange  = color.RGBA{0xFF, 0xB8, 0x4C, 0xFF}
	magenta = color.RGBA{0xFF, 0x4C, 0xEE, 0xFF}
	green   = color.RGBA{0x4C, 0xFF, 0x5E, 0xFF}
	purple  = color.RGBA{0x8B, 0x4C, 0xFF, 0xFF}
	ygreen  = color.RGBA{0xC1, 0xFF, 0x4C, 0xFF}
	red     = color.RGBA{0xFF, 0x4C, 0x67, 0xFF}
	cyan    = color.RGBA{0x4C, 0xFF, 0xE5, 0xFF}
)

var colors = []color.Color{
	blue,
	orange,
	magenta,
	green,
	purple,
	ygreen,
	red,
	cyan,
}

func drawGeoms(m image.Image, ctxt draw2d.GraphicContext, geoms ...*geos.Geometry) {
	// get envelope and calculate scale fn
	var env Envelope
	if len(geoms) > 1 {
		coll, err := geos.NewCollection(geos.GEOMETRYCOLLECTION, geoms...)
		if err != nil {
			log.Fatal(err)
		}
		union, err := coll.UnaryUnion()
		if err != nil {
			log.Fatal(err)
		}
		env = envelope(union)
	} else {
		env = envelope(geoms[0])
	}
	scale := func(x, y float64) (float64, float64) {
		x = env.Px(x)*(WIDTH-2*PAD) + PAD
		y = HEIGHT - (env.Py(y)*(HEIGHT-2*PAD) + PAD)
		return x, y
	}

	var draw func(geoms ...*geos.Geometry)

	draw = func(geoms ...*geos.Geometry) {
		for i, g := range geoms {
			// pick color
			c := colors[i%len(colors)]
			// switch type
			_type, err := g.Type()
			if err != nil {
				log.Fatal(err)
			}
			switch _type {
			case geos.POINT:
				drawPoint(ctxt, g, c, 4.0, scale)
			case geos.LINESTRING, geos.LINEARRING:
				drawLine(ctxt, g, c, 4.0, scale)
			case geos.POLYGON:
				drawPolygon(ctxt, g, c, darken(c), 3.0, scale)
			case geos.MULTIPOINT, geos.MULTILINESTRING, geos.MULTIPOLYGON, geos.GEOMETRYCOLLECTION:
				n, err := g.NGeometry()
				if err != nil {
					log.Fatal(err)
				}
				var subgeoms []*geos.Geometry
				for i := 0; i < n; i++ {
					subgeoms = append(subgeoms, geos.Must(g.Geometry(i)))
				}
				draw(subgeoms...)
			default:
				log.Fatalf("unknown geometry type %v", _type)
			}
		}
	}

	draw(geoms...)
}

func darken(c color.Color) color.Color {
	return color.RGBA{
		R: (c.(color.RGBA).R & 0xfe) >> 1,
		G: (c.(color.RGBA).G & 0xfe) >> 1,
		B: (c.(color.RGBA).B & 0xfe) >> 1,
		A: c.(color.RGBA).A,
	}
}

type Envelope struct {
	Min, Max Point
}

type Point struct {
	X, Y float64
}

func Env(minx, miny, maxx, maxy float64) Envelope {
	return Envelope{Point{minx, miny}, Point{maxx, maxy}}
}

func (e Envelope) Dx() float64 {
	return e.Max.X - e.Min.X
}

func (e Envelope) Dy() float64 {
	return e.Max.Y - e.Min.Y
}

func (e Envelope) Px(x float64) float64 {
	return (x - e.Min.X) / e.Dx()
}

func (e Envelope) Py(y float64) float64 {
	return (y - e.Min.Y) / e.Dy()
}

func envelope(g *geos.Geometry) Envelope {
	env, err := g.Envelope()
	if err != nil {
		log.Fatal(err)
	}
	ring, err := env.ExteriorRing()
	if err != nil {
		log.Fatal(err)
	}
	cs, err := ring.coordSeq()
	if err != nil {
		log.Fatal(err)
	}
	getX := getOrd(cs, (*geos.CoordSeq).GetX)
	getY := getOrd(cs, (*geos.CoordSeq).GetY)
	return Env(getX(0), getY(0), getX(2), getY(2))
}

func getOrd(cs *geos.CoordSeq, fn func(*geos.CoordSeq, int) (float64, error)) func(int) float64 {
	return func(idx int) float64 {
		ord, err := fn(cs, idx)
		if err != nil {
			log.Fatal(err)
		}
		return ord
	}
}

func drawPoint(ctxt draw2d.GraphicContext, g *geos.Geometry, c color.Color, radius float64, scale func(x, y float64) (float64, float64)) {
	if c != nil {
		ctxt.SetFillColor(c)
	}
	x, err := g.X()
	if err != nil {
		log.Fatal(err)
	}
	y, err := g.Y()
	if err != nil {
		log.Fatal(err)
	}
	x, y = scale(x, y)
	ctxt.MoveTo(x, y)
	ctxt.ArcTo(x, y, radius, radius, 0, 2*math.Pi)
	ctxt.Fill()
}

func drawLine(ctxt draw2d.GraphicContext, g *geos.Geometry, c color.Color, width float64, scale func(x, y float64) (float64, float64)) {
	if c != nil {
		ctxt.SetStrokeColor(c)
	}
	if width != 0.0 {
		ctxt.SetLineWidth(width)
	}
	// XXX: should get a [] of points
	cs, err := g.coordSeq()
	if err != nil {
		log.Fatal(err)
	}
	lineCoordSeq(ctxt, cs, scale)
	ctxt.Stroke()
}

func lineCoordSeq(ctxt draw2d.GraphicContext, cs *geos.CoordSeq, scale func(x, y float64) (float64, float64)) {
	n := cs.Size()
	if n == 0 {
		return
	}
	// XXX: interface like sql.Scan() and .Error()
	getX := getOrd(cs, (*geos.CoordSeq).GetX)
	getY := getOrd(cs, (*geos.CoordSeq).GetY)
	x, y := getX(0), getY(0)
	ctxt.MoveTo(scale(x, y))
	for i := 1; i < n; i++ {
		x, y = getX(i), getY(i)
		x, y = scale(x, y)
		ctxt.LineTo(x, y)
	}
}

func drawPolygon(ctxt draw2d.GraphicContext, g *geos.Geometry, fillColor color.Color, strokeColor color.Color, width float64, scale func(x, y float64) (float64, float64)) {
	ctxt.SetFillColor(fillColor)
	ctxt.SetStrokeColor(strokeColor)
	ctxt.SetLineWidth(width)
	// exterior ring
	ring := geos.Must(g.ExteriorRing())
	cs, err := ring.coordSeq()
	if err != nil {
		log.Fatal(err)
	}
	lineCoordSeq(ctxt, cs, scale)
	ctxt.FillStroke()
	// interior rings...
}

func newContext(w, h int) (image.Image, draw2d.GraphicContext) {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	ctxt := draw2d.NewGraphicContext(m)
	ctxt.SetFillColor(image.White)
	ctxt.SetStrokeColor(image.Black)
	return m, ctxt
}

func saveImageToFile(m image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	err = png.Encode(w, m)
	if err != nil {
		log.Fatal(err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
