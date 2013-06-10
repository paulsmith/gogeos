package geos

import (
	"bytes"
	"testing"
)

var wkbDecoderTests = []struct {
	wkb []byte
	wkt string
}{
	{[]byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}, "POINT(-117 35)"},
}

func TestWKBDecoderRead(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	wkbDecoder := NewWKBDecoder()
	for i, test := range wkbDecoderTests {
		g1 := Must(wkbDecoder.Decode(test.wkb))
		g2 := Must(wktDecoder.Decode(test.wkt))
		if !mustEqual(g1.Equals(g2)) {
			t.Errorf("#%d: should equal! got %v want %v", i, g1, g2)
		}
	}
}

var wkbDecoderHexTests = []struct {
	hex string
	wkt string
}{
	{"01010000000000000000405DC00000000000804140", "POINT(-117 35)"},
}

func TestWKBDecoderHexRead(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	wkbDecoder := NewWKBDecoder()
	for i, test := range wkbDecoderHexTests {
		g1 := Must(wkbDecoder.DecodeHex(test.hex))
		g2 := Must(wktDecoder.Decode(test.wkt))
		if !mustEqual(g1.Equals(g2)) {
			t.Errorf("#%d: should equal! got %v want %v", i, g1, g2)
		}
	}
}

var wkbEncoderTests = []struct {
	wkt string
	wkb []byte
}{
	{"POINT(-117 35)", []byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}},
}

func TestWKBEncoderEncode(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	wkbEncoder := NewWKBEncoder()
	for i, test := range wkbEncoderTests {
		g1 := Must(wktDecoder.Decode(test.wkt))
		actual, err := wkbEncoder.Encode(g1)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(actual, test.wkb) {
			t.Errorf("#%d: want %v got %v", i, test.wkb, actual)
		}
	}
}

var wkbEncoderHexTests = []struct {
	wkt string
	wkb []byte
}{
	{"POINT(-117 35)", []byte("01010000000000000000405DC00000000000804140")},
}

func TestWKBEncoderEncodeHex(t *testing.T) {
	wktDecoder := NewWKTDecoder()
	wkbEncoder := NewWKBEncoder()
	for i, test := range wkbEncoderHexTests {
		g1 := Must(wktDecoder.Decode(test.wkt))
		actual, err := wkbEncoder.EncodeHex(g1)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(actual, test.wkb) {
			t.Errorf("#%d: want %v got %v", i, string(test.wkb), string(actual))
		}
	}
}
