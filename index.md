---
layout: default
title: gogeos - Go library for spatial data operations and geometric algorithms
---

# gogeos - Go library for spatial data operations and geometric algorithms

gogeos is a library for Go that provides operations on spatial data and
geometric algorithms. It has bindings to the
[GEOS](http://trac.osgeo.org/geos/) C library.

## Quick usage

```go
package main

import (
	"fmt"
	"log"
	"github.com/paulsmith/gogeos/geos"
)

func main() {
	g, err := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	if err != nil {
		log.Fatal(err)
	}
	b, err := g.Buffer(2.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b) // outputs WKT
}
```

## Getting started / Installation

```bash
$ go get github.com/paulsmith/gogeos
```

## Examples

![](img/example2-a-b.png)

Intersection

![](img/example3-intersection.png)

Union

![](img/example4-union.png)

Difference

![](img/example5-difference.png)

Difference

![](img/example6-difference.png)

Symmetric difference

![](img/example7-symdifference.png)

Buffer

![](img/example8-buffer.png)

Merging linestrings

![](img/example9-unmerged-linestrings.png)
![](img/example9-merged-linestrings.png)
