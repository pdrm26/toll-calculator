.PHONY: obu receiver invoicer


obu:
	@go build -o bin/obu ./obu
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calc:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

invoicer:
	@go build -o bin/invoicer ./invoicer
	@./bin/invoicer

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto	

gate:
	@go build -o bin/gate ./gateway
	@./bin/gate

test:
	@go test -v ./...
