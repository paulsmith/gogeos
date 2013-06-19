.PHONY: fmt check

SRCS = coord.go coordseq.go cwrappers.go geom.go geos.c geos.go geos.h helper.go prepared.go types.go wkb.go wkt.go

all: build

cwrappers.go: geoscapi.py
	python3.3 $< /usr/local/include/geos_c.h > $@
	gofmt -w $@

build: $(SRCS)
	go build

check:
	go test -v

fmt:
	go fmt
