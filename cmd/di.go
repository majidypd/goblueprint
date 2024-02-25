package cmd

import (
	"go.uber.org/dig"

	"github.com/majidypd/goblueprint/config"
	"github.com/majidypd/goblueprint/internal/app"
	"github.com/majidypd/goblueprint/internal/database"
	"github.com/majidypd/goblueprint/internal/router"
)

// NewDig service provider...
func NewDig() *dig.Container {
	container := dig.New()
	if err := container.Provide(routes.NewRouteHandler); err != nil {
		panic(err)
	}
	if err := container.Provide(config.NewConfig); err != nil {
		panic(err)
	}
	if err := container.Provide(app.NewHealthHandler); err != nil {
		panic(err)
	}
	if err := container.Provide(database.NewDB); err != nil {
		panic(err)
	}
	return container
}
