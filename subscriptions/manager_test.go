package subscriptions

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"poc/app/app_mock"
	"poc/bus"
	"poc/bus/bus_mock"
	"poc/model"
	"poc/protos/cloud"
	"poc/protos/cloud/cloud_mock"
	"poc/utils/utils_mock"
	"testing"
	"time"
)

func Test_ShouldHandleObjectFromInboundChannelByPublishingItToOutboundChannelIfThereIsNoAnySubscription(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inboundChannelName := "inbound.test"
	transferChannelName := "transfer.test"
	cachedChannelName := "cached.test"
	outboundChannelName := "outbound.test"
	processedChannelName := "processed.test"
	managerChan := make(bus.DataChannel)

	entity := &cloud.TestEntity{Name: "My Test Entity"}
	entityType := string(entity.ProtoReflect().Descriptor().FullName())
	serialized, err := proto.Marshal(entity)
	if err != nil {
		t.Fatal("could not serialize:", err)
	}
	cloudObj := &cloud.CloudObject{Entity: &anypb.Any{TypeUrl: entityType, Value: serialized}}
	internalServerObject := model.NewIsoFromCloudObject(cloudObj)
	dataEvent := bus.DataEvent{
		Data:  internalServerObject,
		Topic: inboundChannelName,
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().CreateDataChannel().Return(managerChan)
	mockEventBus.EXPECT().Subscribe(inboundChannelName, managerChan)
	mockEventBus.EXPECT().Subscribe(transferChannelName, managerChan)
	mockEventBus.EXPECT().Subscribe(cachedChannelName, managerChan)
	mockEventBus.EXPECT().Publish(outboundChannelName, internalServerObject)

	mockUtils := utils_mock.NewMockIUtils(mockCtrl)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("utils").Return(mockUtils)
	mockAppContext.EXPECT().Get(model.INBOUND_CHANNEL_NAME).Return(inboundChannelName)
	mockAppContext.EXPECT().Get(model.TRANSFER_CHANNEL_NAME).Return(transferChannelName)
	mockAppContext.EXPECT().Get(model.CACHED_CHANNEL_NAME).Return(cachedChannelName)
	mockAppContext.EXPECT().Get(model.OUTBOUND_CHANNEL_NAME).Return(outboundChannelName)
	mockAppContext.EXPECT().Get(model.PROCESSED_CHANNEL_NAME).Return(processedChannelName)

	NewSubscriptionManager(mockAppContext)

	// When
	managerChan <- dataEvent
	time.Sleep(time.Millisecond * 100) // wait for the event to be processed

	// Then - all expected calls are done
}

func Test_ShouldRegisterSubscriptionSuccessfullyAndSendObjectViaStreamAndReceiveAcknowledgementAndUnregisterSubscription(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inboundChannelName := "inbound.test"
	transferChannelName := "transfer.test"
	cachedChannelName := "cached.test"
	outboundChannelName := "outbound.test"
	processedChannelName := "processed.test"
	managerChan := make(bus.DataChannel)

	entity := &cloud.TestEntity{Name: "My Test Entity"}
	entityType := string(entity.ProtoReflect().Descriptor().FullName())
	serialized, err := proto.Marshal(entity)
	if err != nil {
		t.Fatal("could not serialize:", err)
	}
	cloudObj := &cloud.CloudObject{Entity: &anypb.Any{TypeUrl: entityType, Value: serialized}}
	internalServerObject := model.NewIsoFromCloudObject(cloudObj)
	dataEvent := bus.DataEvent{
		Data:  internalServerObject,
		Topic: inboundChannelName,
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().CreateDataChannel().Return(managerChan)
	mockEventBus.EXPECT().Subscribe(inboundChannelName, managerChan)
	mockEventBus.EXPECT().Subscribe(transferChannelName, managerChan)
	mockEventBus.EXPECT().Subscribe(cachedChannelName, managerChan)
	mockEventBus.EXPECT().Publish(processedChannelName, internalServerObject)

	mockUtils := utils_mock.NewMockIUtils(mockCtrl)
	mockSubId := "mock-sub-id"
	mockUtils.EXPECT().GenerateUuid().Return(mockSubId)

	mockStream := cloud_mock.NewMockCloud_SubscribeServer(mockCtrl)
	mockStream.EXPECT().Send(cloudObj)
	mockAck := &cloud.Acknowledge{}
	ackType := string(mockAck.ProtoReflect().Descriptor().FullName())
	serializedAck, err := proto.Marshal(entity)
	if err != nil {
		t.Fatal("could not serialize:", err)
	}
	mockStream.EXPECT().Recv().Return(&cloud.CloudObject{Entity: &anypb.Any{TypeUrl: ackType, Value: serializedAck}}, nil)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("utils").Return(mockUtils)
	mockAppContext.EXPECT().Get(model.INBOUND_CHANNEL_NAME).Return(inboundChannelName)
	mockAppContext.EXPECT().Get(model.TRANSFER_CHANNEL_NAME).Return(transferChannelName)
	mockAppContext.EXPECT().Get(model.CACHED_CHANNEL_NAME).Return(cachedChannelName)
	mockAppContext.EXPECT().Get(model.OUTBOUND_CHANNEL_NAME).Return(outboundChannelName)
	mockAppContext.EXPECT().Get(model.PROCESSED_CHANNEL_NAME).Return(processedChannelName)

	subscriptionManager := NewSubscriptionManager(mockAppContext)

	clientCloseChan := make(chan bool)

	// When
	subId, err := subscriptionManager.RegisterSubscription(entityType, mockStream, clientCloseChan)
	managerChan <- dataEvent
	time.Sleep(time.Millisecond * 100) // wait for the event to be processed
	subscriptionManager.UnregisterSubscription(entityType, subId)

	// Then
	assert.Equal(t, nil, err, "RegisterSubscription should not return an error")
}
