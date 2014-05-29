gogeos - Go library for spatial data operations and geometric algorithms
========================================================================

[![Build Status](https://travis-ci.org/paulsmith/gogeos.png?branch=master)](https://travis-ci.org/paulsmith/gogeos)

gogeos is a library for Go that provides operations on spatial data and
geometric algorithms.

It provides bindings to the [GEOS](http://trac.osgeo.org/geos/) C library.

Quick start
-----------

```go
package main

import (
	"fmt"

	"github.com/paulsmith/gogeos/geos"
)

func main() {
	line, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	buf, _ := line.Buffer(2.5)
	fmt.Println(buf)
	// Output: POLYGON ((18.2322330470336311 21.7677669529663689, 18.61…
}
```

Overview
--------

### Functionality

 * Binary predicates - intersects, disjoint, etc.
 * Topology operations - difference, union, etc.
 * Polygonization, line merging, and simplification
 * Prepared geometries (for better performance for common binary predicates)
 * Validity checking
 * DE-9IM
 * Geometry info - area, length, distance, etc.
 * IO - WKT & WKB read/write

gogeos is an open source project.

### Community

 * [Source code: GitHub project](https://github.com/paulsmith/gogeos)
 * [Issues tracker](https://github.com/paulsmith/gogeos/issues)
 * [Mailing list: gogeos@googlegroups.com](https://groups.google.com/forum/?fromgroups#!forum/gogeos)
 * [IRC: #gogeos on freenode](irc://irc.freenode.net/gogeos)

Installation
------------

### Requirements

 * GEOS 3.3.8 or 3.3.9

GEOS must be installed on your system to build gogeos.

#### Ubuntu

```bash
$ apt-get install libgeos-dev
```

#### OS X - homebrew

```bash
$ brew install geos
```

#### From source (all OSes)

```bash
$ wget http://download.osgeo.org/geos/geos-3.3.8.tar.bz2
$ tar xvfj geos-3.3.8.tar.bz2
$ cd geos-3.3.8
$ ./configure
$ make
$ sudo make install
```

### Installing gogeos

```bash
$ go get github.com/paulsmith/gogeos/geos
```

Documentation
-------------

 * [Main gogeos documentation](http://paulsmith.github.io/gogeos/)
 * [godoc](http://godoc.org/github.com/paulsmith/gogeos/geos)

Example
-------

Let’s say you have two polygons, A (blue) and B (orange).

![](http://paulsmith.github.io/gogeos/img/example2-a-b.png)

One of the most common things to do with a spatial data library like gogeos is
compute the intersection of two or more geometries. Intersection is just
a method on geometry objects in gogeos, which takes one argument, the other
geometry, and computes the intersection with the receiver. The result is a new
geometry, C (magenta):

```go
C := geos.Must(A.Intersection(B))
```

![](http://paulsmith.github.io/gogeos/img/example3-intersection.png)

`geos.Must` is just a convenience function that takes the output of any gogeos
function or method that returns a geometry and an error. It panics if the
error is non-null, otherwise returning the geometry, making it more convenient
to use in single-value contexts. In production code, though, you’ll want to
check the error value.

*(NB: these graphics weren't produced by gogeos directly - I used the
excellent [draw2d](http://code.google.com/p/draw2d/draw2d) package to render
the output of gogeos functions.)*

License
-------

MIT. See `COPYING`.

Copyright (c) 2013 Paul Smith
