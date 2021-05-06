SHELL=/bin/bash

# builds static frontend
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

start: lib
	cd relay && npm run start

# compile contracts
contracts:
	cd eth && npm run compile

# migrates/deploys Solidity contracts via Truffle
migrate:
	cd eth && npm run migrate

# runs local ganache cli
blockchain:
	cd eth && npm run develop

# launches truffle console
console:
	cd eth && npm run console

deploy: migrate

# build frontend
frontend: lib relay

# build all artifacts
all: contracts lib relay

compile: all

build: all

install-lib:
	cd lib && npm i

install-relay:
	cd relay && npm i

install-eth:
	cd eth && npm i
	pip3 install slither-analyzer

install-all: install-lib install-relay install-eth

install: install-all

update-all:
	cd eth && npm update
	cd lib && npm update
	cd relay && npm update
	make audit-fix

update: update-all

audit-fix:
	cd eth && npm audit fix
	cd lib && npm audit fix
	cd relay && npm audit fix

audit-contracts:
	slither eth --filter-paths "@openzeppelin" --truffle-build-directory "../lib/src/abis" --truffle-ignore-compile

docs:
	cd docs && mkdocs build
	cd lib && npm run docs

.PHONY: relay lib contracts docs
