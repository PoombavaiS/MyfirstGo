package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PoombavaiS/MyfirstGo/internal/db"
	migrate "github.com/golang-migrate/migrate/v4"
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
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/movies", services.get_movies)
	http.ListenAndServe(":8080", nil)
	err := godotenv.Load(".env")

	dbConn := db.NewConnection()

	var name string
	var user_id string
	err = dbConn.ConnPool.QueryRow("select name, user_id from users").Scan(&name, &user_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, user_id)

	m, err := migrate.New(os.Getenv("MIGRATION_FILES"), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error 1")
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrLocked {
		fmt.Println("Error 2")
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"Request Received": "Yes",
	}).Info(r)
	fmt.Fprintf(w, "200")
}
