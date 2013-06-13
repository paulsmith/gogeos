check:
	$(MAKE) -C geos check

fmt:
	$(MAKE) -C geos fmt

install:
	go install ./...
