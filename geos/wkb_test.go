package geos

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

var wkbDecoderTests = []struct {
	wkb []byte
	wkt string
}{
	{[]byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}, "POINT(-117 35)"},
}

func TestWkbDecoderRead(t *testing.T) {
	wktDecoder := newWktDecoder()
	wkbDecoder := newWkbDecoder()
	for i, test := range wkbDecoderTests {
		g1 := Must(wkbDecoder.decode(test.wkb))
		g2 := Must(wktDecoder.decode(test.wkt))
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

func TestWkbDecoderHexRead(t *testing.T) {
	wktDecoder := newWktDecoder()
	wkbDecoder := newWkbDecoder()
	for i, test := range wkbDecoderHexTests {
		g1 := Must(wkbDecoder.decodeHex(test.hex))
		g2 := Must(wktDecoder.decode(test.wkt))
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

func TestWkbEncoderEncode(t *testing.T) {
	wktDecoder := newWktDecoder()
	wkbEncoder := newWkbEncoder()
	for i, test := range wkbEncoderTests {
		g1 := Must(wktDecoder.decode(test.wkt))
		actual, err := wkbEncoder.encode(g1)
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

var ewkbEncoderHexTests = []struct {
	wkt  string
	srid int
	wkb  []byte
}{
	{"POINT(-117 35)", 4326, []byte("0101000020E61000000000000000405DC00000000000804140")},
	{"POINT(-117 35)", 900913, []byte("010100002031BF0D000000000000405DC00000000000804140")},
	{"POINT(-117 35)", 0, []byte("01010000000000000000405DC00000000000804140")},
}

func TestWkbEncoderEncodeHex(t *testing.T) {
	wktDecoder := newWktDecoder()
	wkbEncoder := newWkbEncoder()
	for i, test := range wkbEncoderHexTests {
		g1 := Must(wktDecoder.decode(test.wkt))
		actual, err := wkbEncoder.encodeHex(g1)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(actual, test.wkb) {
			t.Errorf("#%d: want %v got %v", i, string(test.wkb), string(actual))
		}
	}
	for i, test := range ewkbEncoderHexTests {
		g1 := Must(wktDecoder.decode(test.wkt))
		g1.SetSRID(test.srid)
		eg1, err := wkbEncoder.encodeEWkb(g1)
		if err != nil {
			panic(err)
		}
		actual := []byte(strings.ToUpper(hex.EncodeToString(eg1)))
		if !bytes.Equal(actual, test.wkb) {
			t.Errorf("#%d: want %v got %v", i, string(test.wkb), string(actual))
		}
	}
}
