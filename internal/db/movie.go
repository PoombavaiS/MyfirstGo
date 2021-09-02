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
	err := dbConn.ConnPool.QueryRow("select * from movies where id = $1", id).Scan(&movie_id, &name)
	conninfo := dbConn.ConnPool.QueryRow("SELECT CURRENT_USER, SESSION_USER")
	fmt.Println(conninfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	movie := Movie{}
	if len(name) > 0 {
		movie.ID = movie_id
		movie.Name = name
	}
	return movie
}

func UpsertMovie(m Movie) error {
	dbConn := NewConnection()
	_, err := dbConn.ConnPool.Exec("Insert into movies(id, name) values($1,$2) ON CONFLICT (id) DO UPDATE SET name = $2", m.ID, m.Name)
	return err
}
