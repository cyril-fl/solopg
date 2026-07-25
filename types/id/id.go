package id

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

func New() ID {
	return ID(uuid.New())
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}