// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/leabago/share-radio/adder/docs/gen"
	"github.com/leabago/share-radio/adder/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=./mocks_usecase_test.go -package=usecase_test

type (

	// Translation -.
	Translation interface {
		Translate(ctx context.Context, userID string, t entity.Translation) (entity.Translation, error)
		History(ctx context.Context, userID string) (entity.TranslationHistory, error)
	}

	// User -.
	User interface {
		Register(ctx context.Context, username, email, password string) (entity.User, error)
		Login(ctx context.Context, email, password string) (string, error)
		GetUser(ctx context.Context, userID string) (entity.User, error)
	}

	// Task -.
	Task interface {
		Create(ctx context.Context, userID, title, description string) (entity.Task, error)
		Get(ctx context.Context, userID, taskID string) (entity.Task, error)
		List(ctx context.Context, userID string, status *entity.TaskStatus, limit, offset int) ([]entity.Task, int, error)
		Update(ctx context.Context, userID, taskID, title, description string) (entity.Task, error)
		Transition(ctx context.Context, userID, taskID string, newStatus entity.TaskStatus) (entity.Task, error)
		Delete(ctx context.Context, userID, taskID string) error
	}

	Station interface {
		CreateStation(ctx context.Context, request gen.CreateStationRequestObject) (gen.StationId, error)
		ListStations(ctx context.Context, request gen.ListStationsRequestObject) (gen.StationListResponse, error)
		GetStation(ctx context.Context, request gen.GetStationRequestObject) (gen.Station, error)
	}

	Genre interface {
		ListGenres(ctx context.Context, request gen.ListGenresRequestObject) (gen.GenreList, error)
	}

	Language interface {
		ListLanguages(ctx context.Context, request gen.ListLanguagesRequestObject) (gen.LanguageList, error)
	}
)
