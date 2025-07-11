.PHONY: obu receiver


obu:
	@go build -o bin/obu ./obu
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

test:
	@go test -v ./...
