package base

import "github.com/google/uuid"

type Object struct {
	Id uuid.UUID `json:"id"`
}

func CreateObject(id uuid.UUID) *Object {
	return &Object{Id:id}
}