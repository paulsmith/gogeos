package geos

import (
	"testing"
)

var wktWriterTests = []struct{ in, out string }{
	{"POINT(-117 33)", "POINT (-117.0000000000000000 33.0000000000000000)"},
}

func TestWKTWriter(t *testing.T) {
	writer := NewWKTWriter()
	decoder := NewWKTDecoder()
	var geom *Geometry
	var err error
	for _, test := range wktWriterTests {
		geom, err = decoder.Decode(test.in)
		if err != nil {
			t.Errorf("WKTDecoder.Decode(): %v", err)
		}
		actual, err := writer.Encode(geom)
		if err != nil {
			t.Errorf("WKTWriter.Encode(): %v", err)
		}
		if actual != test.out {
			t.Errorf("WKTWriter.Encode(): want %v, got %v", test.out, actual)
		}
	}
}
