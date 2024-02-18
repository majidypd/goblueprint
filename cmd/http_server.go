package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/majidypd/goblueprint/config"
	"github.com/majidypd/goblueprint/internal/router"
)

var httpServeCmd = &cobra.Command{
	Use:     "httpserver",
	Aliases: []string{"serv"},
	Short:   "Run http server",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(httpServeCmd)
}

func runServer() {
	container := NewDig()
	if err := container.Invoke(func(cfg *config.Config, router  *routes.ApiRouteHandler) {
		fmt.Printf("Server is listening on port %s...\n", cfg.App.Port)
		if err := http.ListenAndServe(":"+cfg.App.Port, router.Router); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}); err != nil {
		panic(err)
	}
}

