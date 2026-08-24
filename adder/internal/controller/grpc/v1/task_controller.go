package v1

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/leabago/share-radio/adder/docs/proto/v1"
	"github.com/leabago/share-radio/adder/internal/usecase"
	"github.com/leabago/share-radio/adder/pkg/logger"
)

// TaskController -.
type TaskController struct {
	v1.UnimplementedTaskServiceServer

	tk usecase.Task
	l  logger.Interface
	v  *validator.Validate
}
