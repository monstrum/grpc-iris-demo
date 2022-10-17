# Start server

```shell
go run *.go
```

# Start client api

```shell
cd example/client
go run *.go

# Access over web browser

# Access grpc over generated stub
# http://localhost:18080

# Access grpc service over https
# http://localhost:18080/test
```

Tried using IP instead
```go
const (
	address = "127.0.0.1:3443" // Not working
)

// Response
// rpc error: code = Unimplemented desc = unexpected HTTP status code received from server: 404 (Not Found); transport: received unexpected content-type "text/plain; charset=utf-8"
// but works under http://localhost:18080/test
```
Tried using Localhost instead
```go
const (
	address = "localhost:3443" // Not working
)

// Same response
// rpc error: code = Unimplemented desc = unexpected HTTP status code received from server: 404 (Not Found); transport: received unexpected content-type "text/plain; charset=utf-8"
// but works under http://localhost:18080/test
```
Both cases still able be reproduced, same case with other credential method.
```go
cred, _ := credentials.NewClientTLSFromFile("./server.crt", "localhost")

// Expected same response.
// rpc error: code = Unimplemented desc = unexpected HTTP status code received from server: 404 (Not Found); transport: received unexpected content-type "text/plain; charset=utf-8"
// still works under http://localhost:18080/test
```

Now test with standard grpc implementation with different port:
```shell
# Uncomment this line on `example/client/main.go`
go s.serveGrpc(serv, handler)
```

# Then update the address in client example
```go
const (
	address = "localhost:50051"
)

// then using insecure credential
cred := insecure.NewCredentials()
interceptors := []grpc.DialOption{
    grpc.WithTransportCredentials(cred),
    grpc.WithBlock(),
}
```
It works.
