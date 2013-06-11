package geos

import (
	"testing"
)

var wktEncoderTests = []struct{ in, out string }{
	{"POINT(-117 33)", "POINT (-117.0000000000000000 33.0000000000000000)"},
}

func TestWktEncoder(t *testing.T) {
	encoder := newWktEncoder()
	decoder := newWktDecoder()
	var geom *Geometry
	var err error
	for _, test := range wktEncoderTests {
		geom, err = decoder.decode(test.in)
		if err != nil {
			t.Errorf("wktDecoder.decode(): %v", err)
		}
		actual, err := encoder.encode(geom)
		if err != nil {
			t.Errorf("wktEncoder.encode(): %v", err)
		}
		if actual != test.out {
			t.Errorf("wktEncoder.encode(): want %v, got %v", test.out, actual)
		}
	}
}
