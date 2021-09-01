package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PoombavaiS/MyfirstGo/internal/moviebuff"
)

func get_movies(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/movies/")
	fmt.Println(id)

	if len(id) != 0 {
		movie, err := moviebuff.GetMovie(id)
		fmt.Println(movie)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	// dbConn := db.NewConnection()
	// fmt.Println("Movies Request")
	// var name string
	// var movie_id string
	// err := dbConn.ConnPool.QueryRow("select name, movie_id from movies").Scan(&name, &movie_id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(name, movie_id)
}
