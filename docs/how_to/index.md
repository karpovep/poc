## How to tips

* To compile protoc
```$bash
./compile_protos
```

* To run cloud server
```$bash
go run cloud_server/main.go
```

* To run cloud client
```$bash
go run cloud_client/main.go
```

* To install linter
```$bash
GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.26.0
```

* To run linter
```$bash
golangci-lint run
```
