package subscriptions

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"poc/app/app_mock"
	"poc/bus"
	"poc/bus/bus_mock"
	cloud "poc/protos"
	cloud_mock "poc/protos/protos_mock"
	"poc/utils/utils_mock"
	"testing"
	"time"
)

func Test_ShouldRegisterSubscriptionSuccessfullyAndSendIncomingObjectViaStreamAndUnregisterSubscription(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	incomingTopic := "incoming"
	incomingChan := make(bus.DataChannel)

	entity := &cloud.TestEntity{Name: "My Test Entity"}
	entityType := string(entity.ProtoReflect().Descriptor().FullName())
	serialized, err := proto.Marshal(entity)
	if err != nil {
		t.Fatal("could not serialize:", err)
	}
	cloudObj := &cloud.CloudObject{Entity: &anypb.Any{TypeUrl: entityType, Value: serialized}}
	dataEvent := bus.DataEvent{
		Data:  cloudObj,
		Topic: incomingTopic,
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().Subscribe(incomingTopic, incomingChan)

	mockUtils := utils_mock.NewMockIUtils(mockCtrl)
	mockSubId := "mock-sub-id"
	mockUtils.EXPECT().GenerateUuid().Return(mockSubId)

	mockStream := cloud_mock.NewMockCloud_SubscribeServer(mockCtrl)
	mockStream.EXPECT().Send(cloudObj)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("utils").Return(mockUtils)
	mockAppContext.EXPECT().Get("incomingChan").Return(incomingChan)
	mockAppContext.EXPECT().Get("incomingTopic").Return(incomingTopic)

	subscriptionManager := NewSubscriptionManager(mockAppContext)

	clientCloseChan := make(chan bool)

	// When
	subId, err := subscriptionManager.RegisterSubscription(entityType, mockStream, clientCloseChan)
	incomingChan <- dataEvent
	time.Sleep(time.Millisecond * 100) // wait for the event to be processed
	subscriptionManager.UnregisterSubscription(entityType, subId)

	// Then
	assert.Equal(t, nil, err, "RegisterSubscription should not return an error")
}
