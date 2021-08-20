SHELL=/bin/bash

all: valist

valist: web
	go build ./cmd/valist

install: install-lib install-relay valist
	go install ./cmd/valist

install-lib:
	npm install --prefix ./web/lib

install-relay:
	npm install --prefix ./web/relay

web-lib:
	npm run build --prefix ./web/lib

web-relay:
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
