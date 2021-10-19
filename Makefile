SHELL=/bin/bash

all: install valist

bin:
	go build -ldflags "-s -w" ./cmd/valist

bin-linux-amd64:
	GOOS=linux   GOARCH=amd64 go build -ldflags "-s -w" -o dist/linux-amd64/valist       ./cmd/valist

bin-linux-arm64:
	GOOS=linux   GOARCH=arm64 go build -ldflags "-s -w" -o dist/linux-arm64/valist       ./cmd/valist

bin-darwin-amd64:
	GOOS=darwin  GOARCH=amd64 go build -ldflags "-s -w" -o dist/darwin-amd64/valist      ./cmd/valist

bin-darwin-arm64:
	GOOS=darwin  GOARCH=arm64 go build -ldflags "-s -w" -o dist/darwin-arm64/valist      ./cmd/valist

bin-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/windows-amd64/valist.exe ./cmd/valist

bin-multi: bin-linux-amd64 bin-linux-arm64 bin-darwin-amd64 bin-darwin-arm64 bin-windows-amd64

valist: web bin

install: install-lib install-relay

install-lib:
	npm install --prefix ./web/lib

install-relay:
	npm install --prefix ./web/relay

install-docs:
	pip install mkdocs mkdocs-material

dev-lib:
	npm run dev --prefix ./web/lib

dev-relay:
	npm run dev --prefix ./web/relay

dev-docs:
	mkdocs serve

dev:
	@make -j 2 dev-lib dev-relay

web-lib:
	npm run build --prefix ./web/lib

web-relay:
	rm -rf ./web/relay/out
	npm run build --prefix ./web/relay
	npm run export --prefix ./web/relay

web: web-lib web-relay

lint-valist:
	golangci-lint run

lint-web-lib:
	npm run lint --prefix ./web/lib

lint-web-relay:
	npm run lint --prefix ./web/relay

lint: lint-valist lint-web-lib lint-web-relay

test-valist:
	go test ./...

test-web-lib:
	npm run test --prefix ./web/lib

test: test-valist test-web-lib

docs:
	mkdocs build

clean:
	rm -rf ./web/relay/.next
	rm -rf ./web/relay/out
	rm -rf ./web/relay/node_modules
	rm -rf ./web/lib/node_modules
	rm -rf ./web/lib/dist
	rm -rf dist site

publish:
	clean install web bin-multi

.PHONY: web docs
