BINARY_NAME=dpc

build:
	go build -o ${BINARY_NAME} main.go

run-dev:
	go run main.go

run: build
	./dpc

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./...

lint:
	golangci-lint run

mock_gen:
	# go get github.com/golang/mock/gomock
	# go get github.com/golang/mock/mockgen
	mockgen --build_flags=--mod=mod -destination=bus/bus_mock/main.go -package=bus_mock poc/bus IEventBus
	mockgen --build_flags=--mod=mod -destination=protos/cloud/cloud_mock/cloud_grpc.pb.go -package=cloud_mock poc/protos/cloud Cloud_SubscribeServer
	mockgen --build_flags=--mod=mod -destination=utils/utils_mock/main.go -package=utils_mock poc/utils IUtils
	mockgen --build_flags=--mod=mod -destination=utils/utils_mock/cancellable_timer.go -package=utils_mock poc/utils ICancellableTimer
	mockgen --build_flags=--mod=mod -destination=app/app_mock/context.go -package=app_mock poc/app IAppContext
	mockgen --build_flags=--mod=mod -destination=repository/repository_mock/main.go -package=repository_mock poc/repository IRepository

compile_protos:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        protos/cloud/cloud.proto protos/nodes/nodes.proto

docker-start:
	./docker/scripts/docker-start

docker-stop:
	./docker/scripts/docker-stop

install:
	brew install protobuf
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go_vendor:
	go mod vendor
