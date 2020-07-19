SHELL=/bin/bash

# builds frontend to ui/public (will build lib first)
ui:
	cd ui && npm run build

# builds valist npm package
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
