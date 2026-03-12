MAIN_PATH=./cmd/artifact_store
OUTPUT_PATH=./output
BIN_NAME=as

.PHONY: help clean

clean:
	go clean
	rm ${OUTPUT_PATH}/*

test:
	go test -v ./...

build:
	mkdir -p ${OUTPUT_PATH}
	go build -o ${OUTPUT_PATH}/${BIN_NAME} ${MAIN_PATH}

run: build
	go run ${MAIN_PATH}
