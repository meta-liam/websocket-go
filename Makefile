SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

utest:
	go mod tidy; \
	export PROJECT_ENV="unit" && go test -coverpkg=./... -coverprofile=coverage.data ./... -gcflags=-l ;
