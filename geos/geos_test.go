package geos

import (
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	const re = `3\.[34]?\.\d+-CAPI-1\.[78]?\.\d+( .+)?$`
	version := Version()
	matched, err := regexp.MatchString(re, version)
	if err != nil {
		t.Fatal("Version regex:", err)
	}
	if !matched {
		t.Errorf("Version(): %q didn't match regex \"%s\"", version, re)
	}
}
