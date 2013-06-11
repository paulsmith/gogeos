package geos

func Must(g *Geometry, err error) *Geometry {
	if err != nil {
		panic(err)
	}
	return g
}

func MustCoordSeq(c *CoordSeq, err error) *CoordSeq {
	if err != nil {
		panic(err)
	}
	return c
}

func MustCoords(c []Coord, err error) []Coord {
	if err != nil {
		panic(err)
	}
	return c
}
