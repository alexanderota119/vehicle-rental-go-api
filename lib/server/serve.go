package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rfauzi44/vehicle-rental/router"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	config := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return config
}

func serve(cmd *cobra.Command, args []string) error {
	mainRoute, err := router.NewApp()
	if err != nil {
		return err
	}

	var address string = "0.0.0.0:8080"
	if appPort := os.Getenv("APP_PORT"); appPort != "" {
		address = ":" + appPort
	}

	corsConfig := corsHandler()

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
		Handler:      corsConfig.Handler(mainRoute),
	}

	log.Println("app run on port", address)

	return srv.ListenAndServe()

}
