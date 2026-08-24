package stations

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const createStation = `
INSERT INTO public.radio_stations
(id, "name", url, genre, country, "language", is_active, popular_rating, total_plays, added_by, added_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
returning id;
`

type Repo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) Repo {
	return Repo{
		Postgres: pg,
	}
}

func (r *Repo) CreateStation(ctx context.Context, task *entity.Station) (string, error) {

	row, err := r.Postgres.Pool.Query(
		ctx,
		createStation,
		task.Id,
		task.Name,
		task.Url,
		task.Genre,
		task.Country,
		task.Language,
		task.IsActive,
		task.PopularRating,
		task.TotalPlays,
		task.AddedBy,
		task.AddedAt,
		task.UpdatedAt,
	)
	if err != nil {
		return "", err
	}

	res, err := pgx.CollectOneRow(row, pgx.RowToStructByName[entity.StationId])
	if err != nil {
		return "", err
	}

	return res.Id, nil

}
