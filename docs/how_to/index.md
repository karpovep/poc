## How to tips

* To compile protoc
```$bash
make compile_protos
```

* To run cloud server
```$bash
make run
```

* To build
```$bash
make build
```

* To clean-up
```$bash
make clean
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
make lint
```

* To generate mocks
```$bash
make mock_gen
```

* To run tests
```$bash
make test
```

* To start dockers
```$bash
make docker-start
```

* To stop dockers
```$bash
make docker-stop
```
