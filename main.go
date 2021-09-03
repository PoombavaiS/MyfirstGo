package main

import (
	"fmt"
	"net/http"

	db "github.com/PoombavaiS/MyfirstGo/internal/db"
	handlers "github.com/PoombavaiS/MyfirstGo/internal/handlers"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"Is server started": "True",
	}).Info("Go app server started")
	db.Migrations()

	http.HandleFunc("/ping", handler)
	http.HandleFunc("/movies/", handlers.GetMovies)
	http.ListenAndServe(":8080", nil)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("could not load env")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"Request Received": "Yes",
	}).Info(r)
	fmt.Fprintf(w, "200")
}
