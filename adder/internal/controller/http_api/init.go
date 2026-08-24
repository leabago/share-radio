package http_api

import (
	"net/http"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/docs"
	"github.com/evrone/go-clean-template/docs/gen"
	"github.com/evrone/go-clean-template/internal/controller/http-api/api"
	"github.com/evrone/go-clean-template/internal/controller/http-api/middleware"

	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/jwt"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, cfg *config.Config, t usecase.Translation, u usecase.User, tk usecase.Task, jwtManager *jwt.Manager, l logger.Interface) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Prometheus metrics
	if cfg.Metrics.Enabled {
		prometheus := fiberprometheus.New("my-service-name")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Swagger
	if cfg.Swagger.Enabled {

		app.Use(swagger.New(swagger.Config{
			BasePath:    "/",
			FileContent: docs.ApiSchema,
			Path:        "swagger",

			//CacheAge: 1,
		}))

	}

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	bffController := api.NewController(u)
	wrappedHandler := gen.NewStrictHandler(bffController, nil)
	gen.RegisterHandlers(app, wrappedHandler)

	tasks := app.Group("tasks")
	tasks.Use(middleware.Tasks)

}
