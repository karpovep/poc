package model

import (
	"github.com/gocql/gocql"
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
	if cloudObj.Id == "" {
		//todo check cloudObj.Id for the TimeUUID validity
		cloudObj.Id = gocql.TimeUUID().String()
	}
	return &InternalServerObject{
		Object: cloudObj,
		Metadata: &ObjectMeta{
			Status: NEW,
		},
	}
}

const (
	INBOUND_CHANNEL_NAME     string = "inbound"
	TRANSFER_CHANNEL_NAME    string = "transfer"
	OUTBOUND_CHANNEL_NAME    string = "outbound"
	UNPROCESSED_CHANNEL_NAME string = "unprocessed"
	RETRY_CHANNEL_NAME       string = "retry"
	CACHED_CHANNEL_NAME      string = "cached"
	PROCESSED_CHANNEL_NAME   string = "processed"

	NEW       ObjectStatus = "NEW"
	PROCESSED ObjectStatus = "PROCESSED"
)
