SHELL=/bin/bash

# builds frontend to ui/public
relay:
	cd relay && npm run build

# builds valist npm package
lib:
	cd lib && npm run build

# runs local next server
dev-relay:
	cd relay && npm run dev

# runs local next server
dev-lib:
	cd lib && npm run dev

# runs both dev servers in parallel, piping output to same shell
dev:
	@make -j 2 dev-lib dev-relay

# compile contracts
contracts:
	cd eth && npm run compile

# migrates/deploys Solidity contracts via Truffle
migrate:
	cd eth && npm run migrate

deploy: migrate

# build frontend
frontend: lib relay

# build all artifacts
all: contracts lib relay

install:
	cd lib && npm i
	cd relay && npm i

install-all:
	cd lib && npm i
	cd relay && npm i
	cd eth && npm i

.PHONY: relay lib contracts
