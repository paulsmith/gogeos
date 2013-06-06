all: build

cwrappers.go: geoscapi.py
	python3.3 $< /usr/local/include/geos_c.h > $@
	gofmt -w $@

build: geom.go
	go build

check: geom.go
	go test -v
