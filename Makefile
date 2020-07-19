SHELL=/bin/bash

# builds frontend to ui/public
ui:
	cd ui && npm run build

# runs local gatsby server
dev:
	cd ui && npm run develop

lib:
	cd lib && npm run build

# migrates/deploys Solidity contracts via Truffle
migrate:
	truffle migrate

deploy: migrate

.PHONY: ui lib
