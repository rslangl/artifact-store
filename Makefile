# --------------------
# Project variables
# --------------------

MAIN_PATH := ./cmd/artifact_store
BUILD_PATH := ./build/package
OUTPUT_PATH := ./output
CONFIG_PATH := ./config
TMP_PATH := ${OUTPUT_PATH}/tmp
BIN_NAME := artifactstore
DOCKER_FILE_PATH := ./build/package/Dockerfile
DOCKER_TAG := local

# --------------------
# Helpers
# --------------------

.PHONY: help
help:
	@printf "%s\n" "Usage:" \
		"" "Helpers:" \
		"   make help 			This helptext" \
		"   make tools 			Install tools" \
		"" "Quality control:" \
		"   make audit 			Run quality control checks" \
		"   make test 			Run all tests" \
		"   make test/cover 		Run all tests and display coverage" \
		"   make upgradeable 		List direct dependencies that have upgrades available" \
		"" "Build/development:" \
		"   make codegen 		Generate API code from specification" \
		"   make build 			Build binary" \
		"   make docker-build 		Build docker image" \
		"" "Run:" \
		"   make run 			Run binary" \
		"   make docker-run 		Run docker container"

.PHONY: tools
tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# --------------------
# Quality control
# --------------------

.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=${TMP_PATH}/coverage.out ./...
	go tool cover -html=${TMP_PATH}/coverage.out

.PHONY: upgradeable
upgradeable:
	go run github.com/oligot/go-mod-upgrade@latest

# --------------------
# Build/development
# --------------------

.PHONY: tidy
tidy:
	go mod tidy -v ./...
	go fix ./...
	go fmt ./...

.PHONY: codegen
codegen:
	oapi-codegen --config=${CONFIG_PATH}/codegen.yaml api/openapi.yaml

.PHONY: build
build: codegen
	mkdir -p ${OUTPUT_PATH}
	go build -o ${OUTPUT_PATH}/${BIN_NAME} ${MAIN_PATH}

.PHONY: docker-build
docker-build: codegen
	docker build -t ${BIN_NAME}:${DOCKER_TAG} -f ${DOCKER_FILE_PATH} .

# --------------------
# Build/development
# --------------------

.PHONY: run
run: build
	go run ${MAIN_PATH}

.PHONY: docker-run
docker-run:
	docker run --rm -it ${BIN_NAME}:${DOCKER_TAG}
