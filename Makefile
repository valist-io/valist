SHELL=/bin/bash

# builds frontend to ui/public
ui:
	cd ui && npm run build

lib:
	cd lib && npm run build

# runs local gatsby server
dev: lib
	cd ui && npm run develop

# migrates/deploys Solidity contracts via Truffle
migrate:
	truffle migrate

deploy: migrate

.PHONY: ui lib
