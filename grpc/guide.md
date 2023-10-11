### Prepare Protocol Buffers Environment

1. download Protocol Buffers (protobuf) from
   [here](https://github.com/protocolbuffers/protobuf/releases/latest)

2. install the Go protocol buffers plugin (usage: compile *.proto to *.go)

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Run the Compiler

```
example: execute at root directory
protoc --go_out=. --go-grpc_out=. --go_opt=module="grpc" --go-grpc_opt=module="grpc" pb/user.proto
```

---

> - [Protocol Buffer Basic: Go](https://pkg.go.dev/google.golang.org/protobuf/proto)
> - [proto3 Language Guide](https://protobuf.dev/programming-guides/proto3)
> - [Go api](https://pkg.go.dev/google.golang.org/protobuf/proto)
> - [Go grpc](https://grpc.io/docs/languages/go)
