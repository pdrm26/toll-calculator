build:
	@go build -o bin/obu obu/main.go

run: build
	@./bin/obu

test:
	@go test -v ./...
