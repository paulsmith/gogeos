package geos

import (
	"testing"
)

func TestCoordSeq(t *testing.T) {
	var test = struct {
		size, dims, idx int
		x, y, z         float64
	}{
		3, 2, 1, 3.14, 0.0, -1000.1,
	}
	cs := newCoordSeq(test.size, test.dims)
	if cs == nil {
		t.Errorf("newCoordSeq(): got nil from C API")
	}
	var err error
	if err = cs.setX(test.idx, test.x); err != nil {
		t.Errorf("CoordSeq.setX(): %v", err)
	}
	if err = cs.setY(test.idx, test.y); err != nil {
		t.Errorf("CoordSeq.setY(): %v", err)
	}
	if err = cs.setZ(test.idx, test.z); err != nil {
		t.Errorf("CoordSeq.setZ(): %v", err)
	}
	var val float64
	if val, err = cs.x(test.idx); err != nil {
		t.Errorf("CoordSeq.x(%v): %v", test.idx, err)
	}
	if val != test.x {
		t.Errorf("CoordSeq.x(%v): want %v, got %v", test.idx, test.x, val)
	}
	if val, err = cs.y(test.idx); err != nil {
		t.Errorf("CoordSeq.y(%v): %v", test.idx, err)
	}
	if val != test.y {
		t.Errorf("CoordSeq.y(%v): want %v, got %v", test.idx, test.y, val)
	}
	if val, err = cs.z(test.idx); err != nil {
		t.Errorf("CoordSeq.z(%v): %v", test.idx, err)
	}
	if val != test.z {
		t.Errorf("CoordSeq.z(%v): want %v, got %v", test.idx, test.z, val)
	}
	size, err := cs.size()
	if err != nil {
		t.Fatalf("size(): error: %v", err)
	}
	if size != test.size {
		t.Errorf("CoordSeq.size(): want %v, got %v", test.size, size)
	}
	dims, err := cs.dims()
	if err != nil {
		t.Fatalf("dims(): error: %v", err)
	}
	if dims != test.dims {
		t.Errorf("CoordSeq.dims(): want %v, got %v", test.dims, dims)
	}
}
