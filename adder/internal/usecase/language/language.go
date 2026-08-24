package genre

import (
	"context"

	"github.com/leabago/share-radio/adder/docs/gen"
	"github.com/leabago/share-radio/adder/internal/repo"
)

// UseCase -.
type GenreCase struct {
	repo repo.GenreRepo
}

// New returns a Task usecase instrumented with OpenTelemetry tracing spans.
func New(r repo.GenreRepo) *GenreCase {
	return &GenreCase{
		repo: r,
	}
}

func (g *GenreCase) ListGenres(ctx context.Context, request gen.ListGenresRequestObject) (gen.GenreList, error) {

	allGenres, err := g.repo.ListGenres(ctx)
	if err != nil {
		return gen.GenreList{}, err
	}

	genres := make([]gen.Genre, len(allGenres))

	for i, genre := range allGenres {
		genres[i] = genre.ConvertToHttpGenre()
	}

	resp := gen.GenreList{
		Genres: &genres,
	}

	return resp, nil
}
