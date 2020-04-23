PROJECT_NAME := "cloudlint"
PKG := "github.com/pipetail/cloudlint"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint

all: build

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

functions := $(shell find internal/app -name \*main.go | awk -F'/' '{print $$3}')

build: ## Build golang binaries
	@for function in $(functions) ; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$function internal/app/$$function/*.go ; \
	done

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

# race: dep ## Run data race detector
# 	@go test -race -short ${PKG_LIST}

# msan: dep ## Run memory sanitizer
# 	@go test -msan -short ${PKG_LIST}

# coverage: ## Generate global code coverage report
# 	./tools/coverage.sh;

# coverhtml: ## Generate global code coverage report in HTML
# 	./tools/coverage.sh html;

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u golang.org/x/lint/golint
