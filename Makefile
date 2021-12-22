SHELL=/bin/bash

LDFLAGS=-ldflags "-s -w -X 'main.Version=v0.6.3'"

all: valist

valist: bin

bin:
	go build $(LDFLAGS) ./cmd/valist

bin-linux-amd64:
	GOOS=linux   GOARCH=amd64 go build $(LDFLAGS) -o dist/linux-amd64/valist       ./cmd/valist

bin-linux-arm64:
	GOOS=linux   GOARCH=arm64 go build $(LDFLAGS) -o dist/linux-arm64/valist       ./cmd/valist

bin-darwin-amd64:
	GOOS=darwin  GOARCH=amd64 go build $(LDFLAGS) -o dist/darwin-amd64/valist      ./cmd/valist

bin-darwin-arm64:
	GOOS=darwin  GOARCH=arm64 go build $(LDFLAGS) -o dist/darwin-arm64/valist      ./cmd/valist

bin-windows-amd64:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/windows-amd64/valist.exe ./cmd/valist

bin-multi: bin-linux-amd64 bin-linux-arm64 bin-darwin-amd64 bin-darwin-arm64 bin-windows-amd64

install-docs:
	pip install mkdocs mkdocs-material

dev-docs:
	cd docs && mkdocs serve

start:
	go run ./cmd/valist daemon

lint:
	golangci-lint run

test:
	go test ./...

docs:
	cd docs && mkdocs build

clean:
	rm -Rf dist site

publish: clean
	goreleaser build
	go run ./cmd/valist publish

publish-dryrun: clean bin-multi
	go run ./cmd/valist publish --dryrun

.PHONY: docs
