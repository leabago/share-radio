package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/leabago/share-radio/adder/docs/gen"
)

type Station struct {
	Id        string
	Name      string
	Url       string
	Genre     string
	Language  *string
	Icon      *string
	IsActive  *bool
	IsNew     *bool
	AddedAt   *time.Time
	UpdatedAt *time.Time
}

func (s Station) ConvertToHttpStation() (gen.Station, error) {

	id, err := uuid.Parse(s.Id)
	if err != nil {
		return gen.Station{}, err
	}

	return gen.Station{
		Id:        new(id),
		Name:      new(s.Name),
		Url:       new(s.Url),
		Genre:     new(s.Genre),
		Language:  s.Language,
		Icon:      s.Icon,
		IsActive:  s.IsActive,
		IsNew:     s.IsNew,
		AddedAt:   s.AddedAt,
		UpdatedAt: s.UpdatedAt,
	}, nil
}

func (s Station) ConvertToHttpStationDetail() (gen.StationDetail, error) {
	id, err := uuid.Parse(s.Id)
	if err != nil {
		return gen.StationDetail{}, err
	}

	return gen.StationDetail{
		Id:        new(id),
		Name:      new(s.Name),
		Url:       new(s.Url),
		Genre:     new(s.Genre),
		Language:  s.Language,
		Icon:      s.Icon,
		IsActive:  s.IsActive,
		IsNew:     s.IsNew,
		AddedAt:   s.AddedAt,
		UpdatedAt: s.UpdatedAt,
	}, nil
}

func ConvertCreateStationRequestObject(request gen.CreateStationRequestObject) (Station, error) {

	return Station{
		Name:     request.Body.Name,
		Url:      request.Body.Url,
		Genre:    request.Body.Genre,
		Language: request.Body.Language,
	}, nil
}

type StationId struct {
	Id string `db:"id"`
}
