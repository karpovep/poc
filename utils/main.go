package utils

import "github.com/google/uuid"

type (
	IUtils interface {
		GenerateUuid() string
	}

	Utils struct{}
)

func NewUtils() IUtils {
	return &Utils{}
}

func (u *Utils) GenerateUuid() string {
	return uuid.New().String()
}
