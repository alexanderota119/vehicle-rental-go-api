package server

import (
	"log"
	"net/http"
	"os"

	"github.com/rfauzi44/vehicle-rental/router"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	RunE:  serve,
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

	log.Println("App running on port", address)

	return http.ListenAndServe(address, mainRoute)

}
