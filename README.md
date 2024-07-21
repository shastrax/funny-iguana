# funny-iguana
Project Funny gRPC

## buf
[buf](https://buf.build/docs/tutorials/getting-started-with-buf-cli)
[protobuf-go](https://github.com/protocolbuffers/protobuf-go)
[connectrpc](https://github.com/connectrpc/connect-go)
[h2c](https://pkg.go.dev/golang.org/x/net/http2/h2c)
[python grpc](https://github.com/grpc/grpc)

## generate go and python
```buf generate proto```
```buf lint proto```

## go application
```go get github.com/shastrax/funny-iguana/proto/iguana/v1```
```go get github.com/shastrax/funny-iguana/gen/iguana/v1/iguanav1connect```

## github action 
Runs linter and generates fresh proto stubs