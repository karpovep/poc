BINARY_NAME=dpc

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./...

lint:
	golangci-lint run

mock-gen:
	# go get github.com/golang/mock/gomock
	# go get github.com/golang/mock/mockgen
	mockgen -destination=bus/bus_mock/main.go -package=bus_mock poc/bus IEventBus
	mockgen -destination=protos/protos_mock/cloud_grpc.pb.go -package=cloud_mock poc/protos Cloud_SubscribeServer
	mockgen -destination=utils/utils_mock/main.go -package=utils_mock poc/utils IUtils
	mockgen -destination=utils/utils_mock/cancellable_timer.go -package=utils_mock poc/utils ICancellableTimer
	mockgen -destination=app/app_mock/context.go -package=app_mock poc/app IAppContext

compile_protos:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        protos/cloud.proto

docker-start:
	./docker/local/scripts/docker-start

docker-stop:
	./docker/local/scripts/docker-start
