package geos

import (
	"fmt"
)

// Coord represents a coordinate in 3-dimensional space.
type Coord struct {
	X, Y, Z float64
}

// NewCoord is the constructor for a Coord object.
func NewCoord(x, y float64) Coord {
	return Coord{x, y, 0}
}

// String returns a (2d) string representation of a Coord.
func (c Coord) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}

// coordSlice constructs a slice of Coord objects from a coordinate sequence.
func coordSlice(cs *coordSeq) ([]Coord, error) {
	size, err := cs.size()
	if err != nil {
		return nil, err
	}
	coords := make([]Coord, size)
	for i := 0; i < size; i++ {
		x, err := cs.x(i)
		if err != nil {
			return nil, err
		}
		y, err := cs.y(i)
		if err != nil {
			return nil, err
		}
		coords[i] = Coord{X: x, Y: y}
	}
	return coords, nil
}
