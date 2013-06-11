---
layout: default
title: gogeos - Go library for spatial data operations and geometric algorithms
---

# gogeos - Go library for spatial data operations and geometric algorithms

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
 * [Mailing list](https://groups.google.com/forum/?fromgroups#!forum/gogeos): [gogeos@googlegroups.com](mailto:gogeos@googlegroups.com)
 * IRC: [#gogeos](irc://irc.freenode.net/gogeos) on freenode

<h2 id="installation">Installation</h2>

### Requirements

 * GEOS 3.3.x or greater

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

<h2 id="examples">Examples</h2>

### Binary topology operations

gogeos provides binary topology operations, which produce a new geometry from
spatial analysis performed on two input geometries. For example, say we have
these two overlapping polygons, A in blue and B in orange:

![](img/example2-a-b.png)

*(NB: these graphics weren't produced by gogeos directly - I used the
excellent [draw2d](http://code.google.com/p/draw2d/draw2d) package to render
the output of gogeos functions.)*

Then the following operations will produce new geometries in magenta:

<div class="pure-g-r">
	<div class="pure-u-1-2">
		<h4>Intersection</h4>

		<p><img src="img/example3-intersection.png" alt=""></p>
	</div>

	<div class="pure-u-1-2">
		<h4>Union</h4>

		<p><img src="img/example4-union.png" alt=""></p>
	</div>

	<div class="pure-u-1-2">
		<h4>Difference (A of B)</h4>

		<p><img src="img/example5-difference.png" alt=""></p>
	</div>

	<div class="pure-u-1-2">
		<h4>Difference (B of A)</h4>

		<p><img src="img/example6-difference.png" alt=""></p>
	</div>

	<div class="pure-u-1-2">
		<h4>Symmetric difference</h4>

		<p><img src="img/example7-symdifference.png" alt=""></p>
	</div>
</div>

### Unary topology operations

gogeos can produce new geometries based on a operation performed on a single
geometry, perhaps with some input. For example, given a linestring, the
`Buffer()` method produces a new polygon:

#### Buffer

![](img/example8-buffer.png)

### Merging linestrings

For a collection of fully noded linestrings, a new collection can be produced
that merges together the linestrings that touch only at their start and end
points. This is provided by calling the `LineMerge()` method on a
MultiLineString collection:

```go
var linestrings = []*geos.Geometry{
	// ...
}
coll := geos.Must(geos.NewCollection(geos.MULTILINESTRING, linestrings...))
coll.LineMerge()
```

<div class="pure-g-r">
	<div class="pure-u-1-2">
	    <p>
	        <i>Before</i>
	    </p>
		<p>
	        <img src="img/example9-unmerged-linestrings.png">
		</p>
	</div>
	<div class="pure-u-1-2">
	    <p>
	        <i>After</i>
	    </p>
		<p>
	        <img src="img/example9-merged-linestrings.png">
		</p>
	</div>
</div>

*These examples were inspired by the developer’s guide to the
[JTS](http://www.vividsolutions.com/jts/JTSHome.htm)*
