package language

import (
	"context"

	"github.com/leabago/share-radio/adder/docs/gen"
	"github.com/leabago/share-radio/adder/internal/repo"
)

type LanguageCase struct {
	repo repo.LanguageRepo
}

func New(r repo.LanguageRepo) *LanguageCase {
	return &LanguageCase{
		repo: r,
	}
}

func (g *LanguageCase) ListLanguages(ctx context.Context, request gen.ListLanguagesRequestObject) (gen.LanguageList, error) {

	allLanguages, err := g.repo.ListLanguages(ctx)
	if err != nil {
		return gen.LanguageList{}, err
	}

	languages := make([]gen.Language, len(allLanguages))

	for i, language := range allLanguages {
		languages[i] = language.ConvertToHttpLanguage()
	}

	resp := gen.LanguageList{
		Languages: &languages,
	}

	return resp, nil
}
