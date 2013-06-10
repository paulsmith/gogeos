package geos

import (
	"bytes"
	"testing"
)

var wkbReaderTests = []struct {
	wkb []byte
	wkt string
}{
	{[]byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}, "POINT(-117 35)"},
}

func TestWKBReaderRead(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	for i, test := range wkbReaderTests {
		g1 := Must(DefaultWKBReader.Read(test.wkb))
		g2 := Must(wktDecoder.Decode(test.wkt))
		if !mustEqual(g1.Equals(g2)) {
			t.Errorf("#%d: should equal! got %v want %v", i, g1, g2)
		}
	}
}

var wkbReaderHexTests = []struct {
	hex string
	wkt string
}{
	{"01010000000000000000405DC00000000000804140", "POINT(-117 35)"},
}

func TestWKBReaderHexRead(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	for i, test := range wkbReaderHexTests {
		g1 := Must(DefaultWKBReader.ReadHex(test.hex))
		g2 := Must(wktDecoder.Decode(test.wkt))
		if !mustEqual(g1.Equals(g2)) {
			t.Errorf("#%d: should equal! got %v want %v", i, g1, g2)
		}
	}
}

var wkbWriterTests = []struct {
	wkt string
	wkb []byte
}{
	{"POINT(-117 35)", []byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}},
}

func TestWKBWriterWrite(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	for i, test := range wkbWriterTests {
		g1 := Must(wktDecoder.Decode(test.wkt))
		actual, err := DefaultWKBWriter.Write(g1)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(actual, test.wkb) {
			t.Errorf("#%d: want %v got %v", i, test.wkb, actual)
		}
	}
}

var wkbWriterHexTests = []struct {
	wkt string
	wkb []byte
}{
	{"POINT(-117 35)", []byte("01010000000000000000405DC00000000000804140")},
}

func TestWKBWriterWriteHex(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	for i, test := range wkbWriterHexTests {
		g1 := Must(wktDecoder.Decode(test.wkt))
		actual, err := DefaultWKBWriter.WriteHex(g1)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(actual, test.wkb) {
			t.Errorf("#%d: want %v got %v", i, string(test.wkb), string(actual))
		}
	}
}
