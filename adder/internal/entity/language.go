package entity

import (
	"github.com/google/uuid"
	"github.com/leabago/share-radio/adder/docs/gen"
)

type Language struct {
	Id   uuid.UUID
	Name string
}

func (g Language) ConvertToHttpLanguage() gen.Language {
	return gen.Language{
		Id:   new(g.Id),
		Name: new(g.Name),
	}
}
