package nodes

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"poc/app/app_mock"
	"poc/bus/bus_mock"
	"poc/config"
	"poc/model"
	"poc/protos/cloud"
	"testing"
)

func Test_ShouldTransferObjectFromNodeCLientToNodeServer(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	errChan := make(chan error)

	id := "some-mock-time-uuid"
	val := &cloud.TestEntity{Name: "Test Entity"}
	serialized, err := proto.Marshal(val)
	if err != nil {
		log.Fatal("could not serialize", err)
	}
	entity := &anypb.Any{TypeUrl: string(val.ProtoReflect().Descriptor().FullName()), Value: serialized}
	iso := model.NewInternalServerObject(&cloud.CloudObject{
		Id:     id,
		Entity: entity,
	})

	transferChannelName := "transfer.test"

	var actualIso *model.InternalServerObject
	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().Publish(transferChannelName, gomock.Any()).Do(func(channnamName string, arg *model.InternalServerObject) {
		actualIso = arg
	})

	cfg := &config.CloudConfig{
		NodeServer: config.NodeServerConfig{
			Port: ":60001",
		},
	}
	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("errChan").Return(errChan)
	mockAppContext.EXPECT().Get("config").Return(cfg)
	mockAppContext.EXPECT().Get(model.TRANSFER_CHANNEL_NAME).Return(transferChannelName)

	nodeServer := NewNodeServer(mockAppContext)
	nodeServer.Start()

	serverAddress := "localhost:60001"
	nodeClient := NewNodeClient(serverAddress)
	nodeClient.Start()

	// When
	err = nodeClient.Transfer(iso)

	// Then - all expected calls are done
	assert.Equal(t, nil, err, "client should transfer obj without error")
	assert.Equal(t, iso.Object.Id, actualIso.Object.Id, "Object with another ID was published")
	assert.Equal(t, iso.Object.Entity.Value, actualIso.Object.Entity.Value, "Object with another entity value was published")
}
