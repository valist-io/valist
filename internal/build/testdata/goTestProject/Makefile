SHELL=/bin/bash

windows: 
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows64 src/main.go

linux: 
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux64 src/main.go

darwin: 
	GOOS=darwin GOARCH=amd64 go build -o ./dist/darwin64 src/main.go

all: windows linux darwin