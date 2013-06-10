package geos

import (
	"testing"
)

var wktEncoderTests = []struct{ in, out string }{
	{"POINT(-117 33)", "POINT (-117.0000000000000000 33.0000000000000000)"},
}

func TestWKTEncoder(t *testing.T) {
	encoder := NewWKTEncoder()
	decoder := NewWKTDecoder()
	var geom *Geometry
	var err error
	for _, test := range wktEncoderTests {
		geom, err = decoder.Decode(test.in)
		if err != nil {
			t.Errorf("WKTDecoder.Decode(): %v", err)
		}
		actual, err := encoder.Encode(geom)
		if err != nil {
			t.Errorf("WKTEncoder.Encode(): %v", err)
		}
		if actual != test.out {
			t.Errorf("WKTEncoder.Encode(): want %v, got %v", test.out, actual)
		}
	}
}
