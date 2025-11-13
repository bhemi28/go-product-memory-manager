build:
	@go build -o ./bin/go-product-memory-manager cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-product-memory-manager

cmd: build
