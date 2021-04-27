package model

import (
	"poc/protos/cloud"
	"poc/protos/nodes"
)

func NewIsoFromCloudObject(cloudObj *cloud.CloudObject) *nodes.ISO {
	return &nodes.ISO{
		CloudObj: cloudObj,
		Metadata: &nodes.IsoMeta{},
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
)
