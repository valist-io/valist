SHELL=/bin/bash

# builds frontend to ui/public
ui:
	cd ui && npm run build

# builds valist npm package
lib:
	cd lib && npm run build

# runs local gatsby server
dev-ui:
	cd ui && npm run develop

# runs local gatsby server
dev-lib:
	cd lib && npm run develop

# runs both dev servers in parallel, piping output to same shell
dev:
	@make -j 2 dev-lib dev-ui

# compile contracts
contracts:
	cd eth && npm run compile

# migrates/deploys Solidity contracts via Truffle
migrate:
	cd eth && npm run migrate

deploy: migrate

# build frontend
frontend: lib ui

# build all artifacts
all: contracts lib ui

install:
	cd lib && npm i
	cd ui && npm i

install-all:
	cd lib && npm i
	cd ui && npm i
	cd eth && npm i

.PHONY: ui lib contracts
