package utils

import (
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type (
	IUtils interface {
		GenerateUuid() string
		GenerateTimeUuid() string
	}

	Utils struct{}
)

func NewUtils() IUtils {
	return &Utils{}
}

func (u *Utils) GenerateUuid() string {
	return uuid.New().String()
}

func (u *Utils) GenerateTimeUuid() string {
	return gocql.TimeUUID().String()
}
