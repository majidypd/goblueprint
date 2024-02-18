package routes

import (
	"github.com/go-chi/chi"
	"github.com/majidypd/goblueprint/internal/app"
)

type ApiRouteHandler struct {
	Router *chi.Mux
}

// NewRouteHandler constructor of RouteHandler
func NewRouteHandler(
	healthHandler *app.HealthHandler) *ApiRouteHandler {

	apiRouteHandler := &ApiRouteHandler{
		Router: chi.NewRouter(),
	}

	apiRouteHandler.Router.Get("/health", healthHandler.Health)
	return apiRouteHandler
}
