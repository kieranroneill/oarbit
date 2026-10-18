SHELL := /bin/bash

.PHONY: dev test

all: install

###
# installation
###

install:
	${MAKE} install_js_deps

install_js_deps:
	@echo ">>> installing javascript dependencies"
	pnpm install

###
# docker
###

dev:
	docker compose \
	 	-f ./deployments/compose.development.yml \
	 	-p oarbit-dev \
	 	--env-file ./configs/api/.env.dev \
	 	--env-file ./configs/postgres/.env.dev \
	 	--env-file ./configs/shared/.env.dev \
	 	--env-file ./configs/web/.env.dev \
		up \
		--build

###
# formatting
###

format:
	${MAKE} format_go
	${MAKE} format_js

format_go:
	@echo ">>> formatting go files"
	gofmt -w .

format_js:
	@echo ">>> formatting javascript files"
	pnpm format

###
# testing
###

test:
	${MAKE} test_unit

test_unit:
	${MAKE} test_go_unit

test_go_unit:
	@echo ">>> running go unit tests"
	go test -v ./...
