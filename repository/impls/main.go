package impls

import (
	"poc/protos/nodes"
)

type (
	IRepositoryImpl interface {
		Start()
		Stop()
		SaveIso(obj *nodes.ISO) error
		ListActiveIso(nodeId string, limit int, page []byte) ([]*nodes.ISO, []byte, error)
		FindIsoByTypeAndId(objType string, id string) (*nodes.ISO, error)
		DeleteActiveIso(iso *nodes.ISO) error
	}
)
