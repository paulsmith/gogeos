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
	cs := NewCoordSeq(test.size, test.dims)
	if cs == nil {
		t.Errorf("NewCoordSeq(): got nil from C API")
	}
	var err error
	if err = cs.SetX(test.idx, test.x); err != nil {
		t.Errorf("CoordSeq.SetX(): %v", err)
	}
	if err = cs.SetY(test.idx, test.y); err != nil {
		t.Errorf("CoordSeq.SetY(): %v", err)
	}
	if err = cs.SetZ(test.idx, test.z); err != nil {
		t.Errorf("CoordSeq.SetZ(): %v", err)
	}
	var val float64
	if val, err = cs.GetX(test.idx); err != nil {
		t.Errorf("CoordSeq.GetX(%v): %v", test.idx, err)
	}
	if val != test.x {
		t.Errorf("CoordSeq.GetX(%v): want %v, got %v", test.idx, test.x, val)
	}
	if val, err = cs.GetY(test.idx); err != nil {
		t.Errorf("CoordSeq.GetY(%v): %v", test.idx, err)
	}
	if val != test.y {
		t.Errorf("CoordSeq.GetY(%v): want %v, got %v", test.idx, test.y, val)
	}
	if val, err = cs.GetZ(test.idx); err != nil {
		t.Errorf("CoordSeq.GetZ(%v): %v", test.idx, err)
	}
	if val != test.z {
		t.Errorf("CoordSeq.GetZ(%v): want %v, got %v", test.idx, test.z, val)
	}
	if size := cs.Size(); size != test.size {
		t.Errorf("CoordSeq.Size(): want %v, got %v", test.size, size)
	}
	if dims := cs.Dims(); dims != test.dims {
		t.Errorf("CoordSeq.Dims(): want %v, got %v", test.dims, dims)
	}
}
