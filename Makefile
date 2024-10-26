.PHONY: all docs sdk test

all: docs sdk

docs:
	go run cmd/docgen/main.go

sdk:
	go run cmd/sdkgen/main.go

test:
	go test ./...

generate:
	go generate ./...
