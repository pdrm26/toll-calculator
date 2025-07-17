# Toll Calculator

```
docker run -d -p 9092:9092 --name broker apache/kafka:latest
```

## Installing protobuf compiler

```
sudo apt install -y protobuf-compiler
```

## Installing GRPC and Protobuffer plugins for Golang.

1. Protobuffers

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

2. GRPC

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. ️Make sure ~/go/bin is in your PATH:
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
