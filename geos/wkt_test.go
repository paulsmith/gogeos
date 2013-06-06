package geos

import (
	"testing"
)

var wktWriterTests = []struct{ in, out string }{
	{"POINT(-117 33)", "POINT (-117.0000000000000000 33.0000000000000000)"},
}

func TestWKTWriter(t *testing.T) {
	writer := NewWKTWriter()
	reader := NewWKTReader()
	var geom *Geometry
	var err error
	for _, test := range wktWriterTests {
		geom, err = reader.Read(test.in)
		if err != nil {
			t.Errorf("WKTReader.Read(): %v", err)
		}
		actual, err := writer.Write(geom)
		if err != nil {
			t.Errorf("WKTWriter.Write(): %v", err)
		}
		if actual != test.out {
			t.Errorf("WKTWriter.Write(): want %v, got %v", test.out, actual)
		}
	}
}
