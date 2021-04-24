package model

import (
	"github.com/google/uuid"
	"poc/protos/cloud"
	"time"
)

type (
	ObjectStatus string

	ObjectMeta struct {
		Id      string
		Status  ObjectStatus
		RetryIn time.Duration
	}

	InternalServerObject struct {
		Object   *cloud.CloudObject
		Metadata *ObjectMeta
	}
)

func NewInternalServerObject(cloudObj *cloud.CloudObject) *InternalServerObject {
	return &InternalServerObject{
		Object: cloudObj,
		Metadata: &ObjectMeta{
			Id:     uuid.New().String(),
			Status: NEW,
		},
	}
}

const (
	INBOUND_CHANNEL_NAME     string = "inbound"
	OUTBOUND_CHANNEL_NAME    string = "outbound"
	UNPROCESSED_CHANNEL_NAME string = "unprocessed"
	RETRY_CHANNEL_NAME       string = "retry"
	CACHED_CHANNEL_NAME      string = "cached"
	PROCESSED_CHANNEL_NAME   string = "processed"

	NEW       ObjectStatus = "NEW"
	PROCESSED ObjectStatus = "PROCESSED"
)
