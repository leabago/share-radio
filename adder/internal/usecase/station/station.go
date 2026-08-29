package station

import (
	"context"

	"github.com/google/uuid"
	"github.com/leabago/share-radio/adder/docs/gen"
	"github.com/leabago/share-radio/adder/internal/entity"
	"github.com/leabago/share-radio/adder/internal/repo"
)

// UseCase -.
type StationCase struct {
	repo repo.StationRepo
}

// New returns a Task usecase instrumented with OpenTelemetry tracing spans.
func New(r repo.StationRepo) *StationCase {
	return &StationCase{
		repo: r,
	}
}

// Create -.
func (sc *StationCase) CreateStation(ctx context.Context, request gen.CreateStationRequestObject) (gen.StationId, error) {

	station, err := entity.ConvertCreateStationRequestObject(request)
	if err != nil {
		return gen.StationId{}, err
	}

	id, err := sc.repo.CreateStation(ctx, &station)
	if err != nil {
		return gen.StationId{}, err
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return gen.StationId{}, err
	}

	resp := gen.StationId{
		Id: new(idUuid),
	}
	return resp, nil
}

func (sc *StationCase) ListStations(ctx context.Context, request gen.ListStationsRequestObject) (gen.StationListResponse, error) {

	stations, err := sc.repo.ListStations(ctx, request)
	if err != nil {
		return gen.StationListResponse{}, err
	}

	getStations := entity.ConvertToHttpStationList(stations)

	resp := gen.StationListResponse{
		Stations: new(getStations),
		Pagination: &gen.Pagination{
			Limit:  request.Params.Limit,
			Offset: request.Params.Offset,
			Total:  new(len(getStations)),
		},
	}

	return resp, nil
}

func (sc *StationCase) GetStation(ctx context.Context, request gen.GetStationRequestObject) (gen.Station, error) {
	stationEntity, err := sc.repo.GetStation(ctx, request.Id.String())
	if err != nil {
		return gen.Station{}, err
	}

	station := stationEntity.ConvertToHttpStation()

	return *station, nil
}
