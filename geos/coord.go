package geos

import (
	"fmt"
)

type Coord struct {
	X, Y, Z float64
}

func NewCoord(x, y float64) Coord {
	return Coord{x, y, 0}
}

func (c Coord) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}

func coordSlice(cs *CoordSeq) ([]Coord, error) {
	size, err := cs.size()
	if err != nil {
		return nil, err
	}
	coords := make([]Coord, 0, size)
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
