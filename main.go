package main

import (
	"log"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rfauzi44/vehicle-rental/router"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {

	r, err := router.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("APP_PORT")
	log.Printf("server running on port %s", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}

}
