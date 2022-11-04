all:
	@more $(MAKEFILE_LIST)
.PHONY: all

fmt:
	go fmt ./...
test:
	go test -v ./...
.PHONY: fmt test
