SHELL=/bin/bash

# builds frontend to ui/public
ui:
	cd ui && npm run build

# builds valist npm package
lib:
	cd lib && npm run build

# runs local gatsby server
dev: lib
	cd ui && npm run develop

# compile contracts
contracts:
	truffle compile

# migrates/deploys Solidity contracts via Truffle
migrate:
	truffle migrate

deploy: migrate

# build frontend
frontend: lib ui

# build all artifacts
all: contracts lib ui

install:
	cd lib && npm i
	cd ui && npm i

.PHONY: ui lib contracts
