package model

import (
	"poc/protos/cloud"
	"poc/protos/nodes"
)

type RepositoryType string

func NewIsoFromCloudObject(cloudObj *cloud.CloudObject) *nodes.ISO {
	return &nodes.ISO{
		CloudObj:           cloudObj,
		Metadata:           &nodes.IsoMeta{},
		TransferredByNodes: map[string]int64{},
	}
}

func NewIsoFromCloudObjectAndMeta(cloudObj *cloud.CloudObject, meta *nodes.IsoMeta) *nodes.ISO {
	return &nodes.ISO{
		CloudObj:           cloudObj,
		Metadata:           meta,
		TransferredByNodes: map[string]int64{},
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

	CASSANDRA_REPOSITORY_TYPE RepositoryType = "cassandra"
)
