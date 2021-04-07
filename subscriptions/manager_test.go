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

func Test_ShouldRegisterSubscriptionSuccessfullyAndSendIncomingObjectViaStream(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	incomingTopic := "incoming"
	incomingChan := make(bus.DataChannel)

	entity := &cloud.TestEntity{Name: "My Test Entity"}
	serialized, err := proto.Marshal(entity)
	if err != nil {
		t.Fatal("could not serialize:", err)
	}
	dataEvent := bus.DataEvent{
		Data:  &anypb.Any{TypeUrl: string(entity.ProtoReflect().Descriptor().FullName()), Value: serialized},
		Topic: incomingTopic,
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().Subscribe(incomingTopic, incomingChan)

	mockUtils := utils_mock.NewMockIUtils(mockCtrl)
	mockSubId := "mock-sub-id"
	mockUtils.EXPECT().GenerateUuid().Return(mockSubId)

	mockStream := cloud_mock.NewMockCloud_SubscribeServer(mockCtrl)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("utils").Return(mockUtils)
	mockAppContext.EXPECT().Get("incomingChan").Return(incomingChan)
	mockAppContext.EXPECT().Get("incomingTopic").Return(incomingTopic)

	subscriptionManager := NewSubscriptionManager(mockAppContext)

	// When
	err = subscriptionManager.RegisterSubscription("mockObjType", mockStream)
	incomingChan <- dataEvent
	time.Sleep(time.Second) //todo

	// Then
	assert.Equal(t, nil, err, "RegisterSubscription should not return an error")
}
