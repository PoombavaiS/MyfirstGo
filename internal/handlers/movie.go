package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PoombavaiS/MyfirstGo/internal/db"
	"github.com/PoombavaiS/MyfirstGo/internal/moviebuff"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/movies/")
	fmt.Println(id)

	movie := db.GetMovie(id)
	if len(movie.Name) > 0 {
		fmt.Println(movie.ID, movie.Name)
	} else {
		movie, err := moviebuff.GetMovie(id)
		fmt.Println(movie)
		if err != nil {
			fmt.Println(err.Error())
		}
		var m db.Movie
		m.ID = movie.UUID
		m.Name = movie.Name
		err = db.UpsertMovie(m)
		if err != nil {
			fmt.Println("Coudn't upsert data")
		}
		movie_json := json.Marshal(movie)
		w.Header().Set("Content-Type", "application/json")
		w.Write(movie_json)
	}
}
