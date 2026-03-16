MAIN_PATH=./cmd/artifact_store
BUILD_PATH=./build/package
OUTPUT_PATH=./output
BIN_NAME=as

.PHONY: help clean build

clean:
	go clean
	rm ${OUTPUT_PATH}/*

test:
	go test -v ./...

codegen:
	oapi-codegen --config=codegen.yaml api/openapi.yaml
	#go tool oapi-codegen --config=codegen.yaml api/openapi.yaml

build: codegen
	mkdir -p ${OUTPUT_PATH}
	go build -o ${OUTPUT_PATH}/${BIN_NAME} ${MAIN_PATH}

docker-build:
	docker build -t ${BIN_NAME}:local ${BUILD_PATH}

run: build
	go run ${MAIN_PATH}
