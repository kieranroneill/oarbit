SHELL := /bin/bash

.PHONY: dev install install_go_deps install_go_dev_deps install_js_deps format format_go format_js migrations_down_dev migrations_up_dev test test_unit test_go_unit

all: install

###
# installation
###

install:
	${MAKE} install_go_deps
	${MAKE} install_go_dev_deps
	${MAKE} install_js_deps

install_go_deps:
	@echo ">>> installing go dependencies"
	go mod download
	go mod verify

install_go_dev_deps:
	@echo ">>> installing go dev dependencies"
	go install github.com/pressly/goose/v3/cmd/goose@latest

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
# database
###

migrations_down_dev:
	docker compose \
		-f ./deployments/compose.development.yml \
		-p oarbit-dev \
		 run \
		--rm migrations_down

migrations_up_dev:
	docker compose \
		-f ./deployments/compose.development.yml \
	 	-p oarbit-dev \
		run \
		--rm migrations_up

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
