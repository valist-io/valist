SHELL=/bin/bash

build-subgraph:
	cd valist-subgraph && npm install && npm run codegen && npm run build

deploy-subgraph:
	cd valist-subgraph && npm run deploy:ganache

build-contracts:
	cd valist-contracts && npm install && npm run build

deploy-contracts:
	cd valist-contracts && npm run deploy:ganache

migrate-contracts:
	cd valist-contracts && npm run migrate:ganache

build-js:
	cd valist-js && lerna bootstrap && lerna run build

dev-js:
	cd valist-js && lerna bootstrap && lerna run dev --parallel

docker-compose:
	bash ./docker-compose.sh

dev:
	$(MAKE) -j 2 docker-compose dev-js

build: build-subgraph build-contracts build-js

deploy: deploy-contracts deploy-subgraph

migrate: migrate-contracts