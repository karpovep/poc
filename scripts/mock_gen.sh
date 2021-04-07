#!/usr/bin/env bash

# go get github.com/golang/mock/gomock
# go get github.com/golang/mock/mockgen

mockgen -destination=bus/bus_mock/main.go -package=bus_mock poc/bus IEventBus
mockgen -destination=protos/protos_mock/cloud_grpc.pb.go -package=cloud_mock poc/protos Cloud_SubscribeServer
mockgen -destination=utils/utils_mock/main.go -package=utils_mock poc/utils IUtils
mockgen -destination=app/app_mock/context.go -package=app_mock poc/app IAppContext
