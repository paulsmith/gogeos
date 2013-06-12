package geos

// Must is a helper that wraps a call to a function returning (*Geometry, error)
// and panics if the error is non-nil.
func Must(g *Geometry, err error) *Geometry {
	if err != nil {
		panic(err)
	}
	return g
}

// MustCoords is a helper that wraps a call to a function returning ([]Coord, error)
// and panics if the error is non-nil.
func MustCoords(c []Coord, err error) []Coord {
	if err != nil {
		panic(err)
	}
	return c
}
