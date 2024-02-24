package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	if err := container.Invoke(func(cfg *config.Config, router *routes.ApiRouteHandler) {
		fmt.Printf("Server is listening on port %s...\n", cfg.App.Port)

		// Create a new HTTP server
		srv := &http.Server{
			Addr:    ":" + cfg.App.Port,
			Handler: router.Router,
		}

		// Start the server in a goroutine
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Error starting server: %s", err)
			}
		}()

		// Wait for an interrupt signal
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		// Initiate a graceful shutdown
		log.Println("Shutting down gracefully, press Ctrl+C again to force")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")
	}); err != nil {
		panic(err)
	}
}
