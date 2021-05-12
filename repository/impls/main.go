package impls

import (
	"poc/protos/nodes"
)

type (
	IRepositoryImpl interface {
		Start()
		Stop()
		SaveIso(obj *nodes.ISO) error
	}
)
