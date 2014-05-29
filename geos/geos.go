// Package geos provides support for creating and manipulating spatial data.
// At its core, it relies on the GEOS C library for the implementation of
// spatial operations and geometric algorithms.
package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

import (
	"fmt"
	"sync"
)

var (
	// Required for the thread-safe GEOS C API (the "*_r" functions).
	handle = C.gogeos_initGEOS()
	// Protects the handle from being used concurrently in multiple C threads.
	handlemu sync.Mutex
)

// XXX: store last error message from handler in a global var (chan?)

// Version returns the version of the GEOS C API in use.
func Version() string {
	return C.GoString(cGEOSversion())
}

// Error gets the last error that occured in the GEOS C API as a Go error type.
func Error() error {
	return fmt.Errorf("geos: %s", C.GoString(C.gogeos_get_last_error()))
}
