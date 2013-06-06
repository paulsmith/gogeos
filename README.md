gogeos - Go library for spatial data operations and geometric algorithms
========================================================================

[![Build Status](https://travis-ci.org/paulsmith/gogeos.png?branch=master)](https://travis-ci.org/paulsmith/gogeos)

gogeos is a library for Go that provides operations on spatial data and
geometric algorithms.

It provides bindings to the [GEOS](http://trac.osgeo.org/geos/) C library.

<h2 id="quickstart">Quick start</h2>

```go
package main

import (
	"fmt"
	"log"

	"github.com/paulsmith/gogeos/geos"
)

func main() {
	line, err := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := line.Buffer(2.5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf)
	// POLYGON ((18.2322330470336311 21.7677669529663689, 18.61…
}
```

<h2 id="overview">Overview</h2>

### Functionality

 * Binary predicates - intersects, disjoint, etc.
 * Topology operations - difference, union, etc.
 * Polygonization, line merging, and simplification
 * Prepared geometries (for better performance for common binary predicates)
 * Validity checking
 * DE-9IM
 * Geometry info - area, length, distance, etc.
 * IO - WKT & WKB read/write

gogeos is an open source project (MIT license).

### Community

 * [Source code: GitHub project](https://github.com/paulsmith/gogeos)
 * [Issues tracker](https://github.com/paulsmith/gogeos/issues)
 * [Mailing list: gogeos@googlegroups.com](https://groups.google.com/forum/?fromgroups#!forum/gogeos)
 * [IRC: #gogeos on freenode](irc://irc.freenode.net/gogeos)

<h2 id="installation">Installation</h2>

### Requirements

 * GEOS 3.2.x or greater

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
$ go get github.com/paulsmith/gogeos
```

License
-------

MIT. See `COPYING`.

Copyright (c) 2013 Paul Smith
