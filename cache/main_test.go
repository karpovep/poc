package cache

import (
	"github.com/golang/mock/gomock"
	"poc/app/app_mock"
	"poc/bus"
	"poc/bus/bus_mock"
	"poc/model"
	"poc/protos/cloud"
	"poc/utils/utils_mock"
	"testing"
	"time"
)

func Test_ShouldProcessObjectBySchedulingReprocessing(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	retryChannelName := "retry.test"
	cachedChannelName := "cached.test"
	retryChan := make(bus.DataChannel)

	cloudObj := &cloud.CloudObject{}
	internalServerObject := model.NewIsoFromCloudObject(cloudObj)
	retryIn := int32(60)
	internalServerObject.Metadata.RetryIn = retryIn
	dataEvent := bus.DataEvent{
		Data:  internalServerObject,
		Topic: retryChannelName,
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().Subscribe(retryChannelName, retryChan)
	mockEventBus.EXPECT().Publish(cachedChannelName, internalServerObject)

	mockTimerCh := make(chan bool)
	mockCacheTimer := utils_mock.NewMockICancellableTimer(mockCtrl)
	mockCacheTimer.EXPECT().After(time.Second * time.Duration(retryIn)).Return(mockTimerCh)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("cacheTimer").Return(mockCacheTimer)
	mockAppContext.EXPECT().Get("retryChan").Return(retryChan)
	mockAppContext.EXPECT().Get(model.RETRY_CHANNEL_NAME).Return(retryChannelName)
	mockAppContext.EXPECT().Get(model.CACHED_CHANNEL_NAME).Return(cachedChannelName)

	NewCache(mockAppContext)

	// When
	retryChan <- dataEvent
	time.Sleep(time.Millisecond * 50) // wait for the event to be processed
	mockTimerCh <- true               // mimics that time passed
	time.Sleep(time.Millisecond * 50) // wait for the event to be processed

	// Then - all expected calls are done
}
