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

// 判断坐标数组是否是逆时针
// 逆时针(oriented counter-clockwise)返回 true
// 顺时针返回 false
func IsCCW(ring []Coord) bool {
	// 默认最后一个点和第一个点相同 ,去掉最后一个点
	_ring := ring[:len(ring)-1]

	inPts := len(_ring)
	if inPts < 3 {
		// 环的点数量少于3
		return false
	}
	// 找最右边点
	var indRightX, indPrev, indNext int = -1, -1, -1
	indRightX = maxRingX(_ring)

	if indRightX != 0 {
		//	不是第一个点
		indPrev = indRightX - 1
	} else {
		indPrev = inPts - 1
	}
	if indRightX != inPts-1 {
		//	不是最后一个
		indNext = indRightX + 1
	} else {
		indNext = 0
	}

	v1X := _ring[indRightX].X - _ring[indPrev].X
	v1Y := _ring[indRightX].Y - _ring[indPrev].Y
	v2X := _ring[indNext].X - _ring[indRightX].X
	v2Y := _ring[indNext].Y - _ring[indRightX].Y
	// cross product of v1 and v2
	zval := v1X*v2Y - v1Y*v2X
	if zval < 0 {
		// 顺时针
		return false
	} else if zval > 0 {
		// 逆时针
		return true
	} else if v1Y > 0 {
		// 顺时针
		return false
	} else {
		// 逆时针
		return true
	}
	return false
}

// 获取ring 最右边的X坐标索引
func maxRingX(ring []Coord) int {
	var maxX float64 = -99999999.0
	var maxXInd int = -1
	for ind, point := range ring {
		if point.X > maxX {
			maxX = point.X
			maxXInd = ind
		}
	}

	if maxXInd == -1 {
		fmt.Errorf("maxRingX run error")
	}
	return maxXInd
}
