package entity

import (
	"github.com/google/uuid"
	"github.com/leabago/share-radio/adder/docs/gen"
)

type Genre struct {
	Id   uuid.UUID
	Name string
}

func (g Genre) ConvertToHttpGenre() gen.Genre {
	return gen.Genre{
		Id:   new(g.Id),
		Name: new(g.Name),
	}
}
