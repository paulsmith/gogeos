package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

import (
	"fmt"
)

// Required for the thread-safe GEOS C API (the "*_r" functions)
var handle = C.gogeos_initGEOS()

// XXX: store last error message from handler in a global var (chan?)

func Version() string {
	return C.GoString(C.GEOSversion())
}

func Error() error {
	return fmt.Errorf("geos: %s", C.GoString(C.gogeos_get_last_error()))
}
