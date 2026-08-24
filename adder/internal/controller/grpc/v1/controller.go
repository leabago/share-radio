package v1

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/leabago/share-radio/adder/docs/proto/v1"
	"github.com/leabago/share-radio/adder/internal/usecase"
	"github.com/leabago/share-radio/adder/pkg/logger"
)

// TranslationController -.
type TranslationController struct {
	v1.UnimplementedTranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
