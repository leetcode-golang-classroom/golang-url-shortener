.PHONY=build

build:
	@go build -o bin/main cmd/main.go

run: build
	@./bin/main

test:
	@go test -v -cover ./test/...

build-tool:
	@go build -o bin/tool tool/msgpack.go
run-tool:
	@./bin/tool