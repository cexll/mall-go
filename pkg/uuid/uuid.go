package uuid

import "github.com/google/uuid"

func NewUuid() string {
	return uuid.New().String()
}
