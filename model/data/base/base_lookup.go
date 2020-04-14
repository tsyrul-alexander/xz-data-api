package base

import "github.com/google/uuid"

type Lookup struct {
	Object
	Name string `json:"name"`
}

func CreateLookup(id uuid.UUID, name string) *Lookup {
	return &Lookup{
		Object: Object{
			Id: id,
		},
		Name: name,
	}
}
