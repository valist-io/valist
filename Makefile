SHELL=/bin/bash

# builds valist npm package
lib:
	cd lib && npm run build

# builds cli
cli:
	cd cli && npm run build

# builds static frontend
relay:
	cd relay && npm run build

# runs local typescript compiler in watch mode
dev-lib:
	cd lib && npm run dev

# runs local next server
dev-relay:
	cd relay && npm run dev

# hot reload docs
dev-docs:
	mkdocs serve

# runs both dev servers in parallel, piping output to same shell
dev:
	@make -j 2 dev-lib dev-relay

# builds and runs relay in production mode
start: lib
	cd relay && npm run start

# build frontend
frontend: lib relay

# build static site (no APIs) into /relay/out
static: frontend
	cd relay && npm run export

# compile contracts
contracts:
	cd hardhat && npm run compile

# migrates/deploys Solidity contracts via Truffle
migrate:
	cd hardhat && npm run migrate

deploy: migrate

# launches truffle console
console:
	cd hardhat && npm run console

# runs local ganache cli
blockchain:
	cd hardhat && npm run develop

# build all artifacts
all: contracts lib relay

compile: all

build: all

install-hardhat:
	cd hardhat && npm i

install-lib:
	cd lib && npm i

install-cli:
	cd cli && npm i

install-relay:
	cd relay && npm i

install-frontend: install-lib install-relay

install-docs:
	pip install mkdocs mkdocs-material

install-all: install-hardhat install-lib install-cli install-relay

install: install-all

update-all:
	cd hardhat && npm update
	cd lib && npm update
	cd cli && npm update
	cd relay && npm update
	make audit-fix

update: update-all

audit-fix:
	cd hardhat && npm audit fix
	cd lib && npm audit fix
	cd cli && npm audit fix
	cd relay && npm audit fix

audit-contracts:
	slither hardhat --filter-paths "@openzeppelin" --truffle-build-directory "../lib/src/abis" --truffle-ignore-compile

docs:
	mkdocs build
	cd lib && npm run docs

.PHONY: relay lib contracts docs cli
