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
