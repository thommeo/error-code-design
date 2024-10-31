.PHONY: all docs sdk test test-verbose

all: test docs sdk

docs:
	go run cmd/docgen/main.go

sdk:
	go run cmd/sdkgen/main.go

test:
	go test ./... -v

test-verbose:
	go test ./... -v -count=1

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

generate:
	go generate ./...
