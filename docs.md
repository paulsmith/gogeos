---
layout: default
title: gogeos Documentation
---

gogeos Documentation
====================

gogeos is a spatial data library for Go, providing bindings to the
production-quality and time tested GEOS C library.

Creating geometries
-------------------

Geometries in gogeos can be created by either from a Well-Know Text (WKT)
string, or by building up from coordinates and sequences of coordinates.
Working with WKT is generally easier and is one of the common formats for
transferring geometries between systems:

```go
poly, err := geos.FromWKT(`POLYGON ((10 10, 20 10, 20 20, 10 20, 10 10))`)
if err != nil {
	log.Fatalf("error decoding WKT: %v", err)
}
```

Building up geometries from coordinates TK ...

More documentation TK!
----------------------

In the meantime, see the [godoc API docs](http://godoc.org/github.com/paulsmith/gogeos/geos)

...

Miscellaneous
-------------

gogeos, and the underlying GEOS C API, use a floating precision model -- that
is, coordinates are represented by floating-point numbers. Therefore,
coordinates generated in the course of computation may lose precision, as they
have more digits than are representable by the floating-point data type (a
`float64` in Go and a `double` in C).
