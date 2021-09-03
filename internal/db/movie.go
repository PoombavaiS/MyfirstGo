package db

import (
	"fmt"
	"os"
)

type Movie struct {
	ID   string
	Name string
}

func GetMovie(id string) Movie {
	dbConn := NewConnection()

	var movie_id, name string
	err := dbConn.ConnPool.QueryRow("select id, name from movies where id = $1", id).Scan(&movie_id, &name)
	if err != nil {
		fmt.Println("Movie NOT found in DB")
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	movie := Movie{}
	fmt.Println(movie)
	if len(name) > 0 {
		movie.ID = movie_id
		movie.Name = name
	}
	return movie
}

func UpsertMovie(m Movie) error {
	dbConn := NewConnection()
	fmt.Println("Insert Movie")
	_, err := dbConn.ConnPool.Exec("Insert into movies(id, name) values($1,$2) ON CONFLICT (id) DO UPDATE SET name = $2", m.ID, m.Name)
	return err
}
