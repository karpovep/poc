package impls

import (
	"poc/protos/nodes"
)

type (
	IRepositoryImpl interface {
		Start()
		Stop()
		SaveInternalServerObject(obj *nodes.ISO) error
	}
)
