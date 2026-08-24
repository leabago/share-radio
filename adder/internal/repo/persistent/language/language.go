package genre

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/leabago/share-radio/adder/internal/entity"
	"github.com/leabago/share-radio/adder/pkg/postgres"
	"github.com/leabago/share-radio/adder/pkg/postgres/sql_query"
)

const selectGenres = `
SELECT id, "name"
FROM public.genres
`

type Repo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{
		Postgres: pg,
	}
}

func (r *Repo) ListGenres(ctx context.Context) ([]entity.Genre, error) {

	order := sql_query.NewOrder().OrderBy("display_order", sql_query.Asc, false).SQL()

	selectGenresReq := selectGenres + order

	row, err := r.Postgres.Pool.Query(
		ctx,
		selectGenresReq,
	)
	if err != nil {
		return nil, err
	}

	genre, err := pgx.CollectRows(row, pgx.RowToStructByName[entity.Genre])
	if err != nil {
		return nil, err
	}

	return genre, nil
}
