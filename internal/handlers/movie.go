package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PoombavaiS/MyfirstGo/internal/db"
	"github.com/PoombavaiS/MyfirstGo/internal/moviebuff"
	"github.com/sirupsen/logrus"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/movies/")
	fmt.Println(id)

	movie := db.GetMovie(id)
	fmt.Println(movie)
	if len(movie.Name) > 0 {
		fmt.Println("Movie present in DB")
		fmt.Println(movie.ID, movie.Name)
		movie_json, err := json.Marshal(movie)
		if err != nil {
			logrus.Infoln(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(movie_json)
	} else {
		fmt.Println("Movie NOT present in DB")
		movie, err := moviebuff.GetMovie(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		if movie != nil {
			var m db.Movie
			fmt.Println(movie)
			m.ID = movie.UUID
			m.Name = movie.Name
			err = db.UpsertMovie(m)
			if err != nil {
				fmt.Println("Coudn't upsert data")
			}
			movie_json, err := json.Marshal(movie)
			if err != nil {
				logrus.Infoln(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(movie_json)
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404! Not found")
		}

	}
}
