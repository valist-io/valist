SHELL=/bin/bash

all: install valist

bin:
	go build -ldflags "-s -w" ./cmd/valist

valist: web bin

install: install-lib install-relay

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

# runs local typescript compiler in watch mode
dev-lib:
	npm run dev --prefix ./web/lib

# runs local next server
dev-relay:
	npm run dev --prefix ./web/relay

# hot reload docs
dev-docs:
	mkdocs serve

# runs both dev servers in parallel, piping output to same shell
dev:
	@make -j 2 dev-lib dev-relay
