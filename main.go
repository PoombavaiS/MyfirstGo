package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/PoombavaiS/go_training_exercises/internal/db"
	moviebuff "github.com/RealImage/moviebuff-sdk-go"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"Is server started": "True",
	}).Info("Go app server started")
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/movies", get_movies)
	http.ListenAndServe(":8080", nil)
	err := godotenv.Load(".env")
	cfg := moviebuff.Config{
		HostURL:     os.Getenv("MB_URL"), //https://moviepass-v2.herokuapp.com/
		StaticToken: os.Getenv("MBAPI_TOKEN"),
	}
	movieData, err := moviebuff.New(cfg).GetMovie("ccbacb63-ce02-4117-a9d1-87869f50369d")
	if err != nil {
		logrus.WithError(err).Errorln("Failed to get Movie Information from Moviebuff for given MovieID")
	}
	fmt.Println(movieData)

	dbConn := db.NewConnection()

	var name string
	var user_id string
	err = dbConn.QueryRow(context.Background(), "select name, user_id from users").Scan(&name, &user_id)
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

func get_movies(w http.ResponseWriter, r *http.Request) {
	dbConn := db.NewConnection()
	fmt.Println("Movies Request")
	var name string
	var movie_id string
	dbConn.QueryRow(context.Background(), "select name, movie_id from movies").Scan(&name, &movie_id)
	fmt.Println(name, movie_id)
}
