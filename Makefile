SHELL=/bin/bash -o pipefail

# Prechecks of needed packages, some can be autoinstalled later on
ARTIFACT_NAME := theenglishcutd

GOPROXY ?= ""

ARTIFACTS_DIR ?= _build

all: build

_artifacts:
	mkdir -p ${ARTIFACTS_DIR}

_bin:
	mkdir -p bin


define compile
	$(eval os = $1)
	$(eval extension = $2)

	$(eval branch = $(shell git rev-parse --abbrev-ref HEAD))
	$(eval commit = $(shell git rev-parse --short HEAD))
	$(eval build_time = $(shell date -u +%s))
	$(eval tag = $(shell git describe --tags))
	$(eval ldflags = "-X main.commitHash=$(commit) -X main.buildTime=$(build_time) -X main.branch=$(branch) -X cmd/version.version=$(tag)")

	echo "building $(os) binary"
	go build -ldflags=$(ldflags) -o bin/$(ARTIFACT_NAME) ./cmd/$(ARTIFACT_NAME).go
endef

.PHONY: build
build: _bin lint test
	$(call compile)

.PHONY: test
test: _artifacts
	@echo "Executing tests"
	gotestsum --format short-verbose --junitfile ${ARTIFACTS_DIR}/junit.xml -- -coverprofile=${ARTIFACTS_DIR}/coverage_ut.out ./...
	@echo "Generating coverage report"
	go tool cover -html=${ARTIFACTS_DIR}/coverage_ut.out -o ${ARTIFACTS_DIR}/coverage_ut.html

.PHONY: lint
lint:
	@echo "Executing linters"
	golangci-lint run

dev:
	gin -a 9876 -d cmd
