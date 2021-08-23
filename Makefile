.PHONY: all build clean run check cover test
BIN_FILE=stock

all: check build

build:
	@go build -o "${BIN_FILE}"

clean:
	@go clean
	rm -f cover.out cover.html "${BIN_FILE}"

test:
	@echo "=== go ut ==="
	@go test ./...

check:
	@echo "=== go lint ==="
	@go get -u golang.org/x/lint/golint
	@echo "run golint ..."
	@golint -set_exit_status ./...

cover:
	@go test -coverprofile cover.out ./...
	@go tool cover -html=cover.out -o cover.html

run:
	./"${BIN_FILE}"

