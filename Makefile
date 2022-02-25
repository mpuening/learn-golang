-include .env

#VERSION := $(shell git describe --tags)
#BUILD := $(shell git rev-parse --short HEAD)
#PROJECTNAME := $(shell basename "$(PWD)")
#ADDR := http://localhost:8080

VERSION := 0.0.0
BUILD := 0.0.0
PROJECTNAME := server
ADDR := http://localhost:8080

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard cmd/*.go)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.version=$(VERSION) -X=main.build=$(BUILD) -X=main.buildtime='`date +%F-%T`"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## build: Build app.
build: compile test

## start: Start in development mode.
start:
	@bash -c "trap 'make stop' EXIT;"
	@$(MAKE) compile start-server

## stop: Stop development mode.
stop: stop-server

start-server: stop-server
	@echo "  >  $(PROJECTNAME) is available at $(ADDR)"
	@$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@touch $(PID)
	@kill `cat $(PID)` 2> /dev/null || true
	@rm $(PID)

restart-server: stop-server start-server

## compile: Compile the binary.
compile:
	@touch $(STDERR)
	@rm $(STDERR)
	@$(MAKE) -s go-compile 2>&1 > $(STDERR)
	@cat $(STDERR)

## test: Test the binaries.
test: go-test

## tidy: It is good to stay tidy.
tidy: go-tidy

## clean: Clean build files. Runs `go clean` internally.
clean:
	@rm -rf $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@$(MAKE) go-clean

go-compile: go-get go-fmt go-build

go-tidy:
	@echo "  >  Tidy up modules..."
	@go mod tidy

go-vendor:
	@echo "  >  Checking vendor issues..."
	@go mod vendor

go-get:
	@echo "  >  Checking for missing dependencies..."
	@go get ./...

go-fmt:
	@echo "  >  Checking for format issues..."
	#@go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	@test -z $(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

go-build:
	@echo "  >  Building binary..."
	@go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-test:
	@echo "  >  Running test cases..."
	@go test ./...

go-clean:
	@echo "  >  Cleaning build and test cache"
	@go clean -testcache
	@/bin/rm -rf $(GOBIN)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

