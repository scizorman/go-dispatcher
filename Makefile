SHELL := /bin/bash -e

all:
	@more $(MAKEFILE_LIST)

lint:
	go vet ./...

fmt:
	go fmt ./...

test:
	go test -v ./...

.PHONY: all lint fmt test
