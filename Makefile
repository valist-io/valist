SHELL=/bin/bash

all: install valist

bin:
	go build ./cmd/valist

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
	npm run docs --prefix ./web/lib

.PHONY: web docs