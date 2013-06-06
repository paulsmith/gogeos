package geos

import (
	"testing"
)

func TestVersion(t *testing.T) {
	expected := "3.3.8-CAPI-1.7.8"
	if actual := Version(); actual != expected {
		t.Errorf("Version(): want %v, got %v", expected, actual)
	}
}
