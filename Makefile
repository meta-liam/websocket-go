SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

utest:
	go mod tidy; \
	export PROJECT_ENV="unit" && go test -coverpkg=./... -coverprofile=coverage.data ./... -gcflags=-l ;

build-mac:
	go mod tidy; \
	rm -rf build/mac | echo "no build dir"; \
	mkdir build/mac; \
	CGO_ENABLED=0 go build -ldflags '-w -s' -o build/mac/main ./v1/main.go;