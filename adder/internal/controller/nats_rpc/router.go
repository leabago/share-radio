package v1

import (
	v1 "github.com/leabago/share-radio/adder/internal/controller/nats_rpc/v1"
	"github.com/leabago/share-radio/adder/internal/usecase"
	"github.com/leabago/share-radio/adder/pkg/jwt"
	"github.com/leabago/share-radio/adder/pkg/logger"
	"github.com/leabago/share-radio/adder/pkg/nats/nats_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation, u usecase.User, tk usecase.Task, j *jwt.Manager, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)

	{
		v1.NewRoutes(routes, t, u, tk, j, l)
	}

	return routes
}
