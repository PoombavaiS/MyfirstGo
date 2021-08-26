package db

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v4"
)

var cPool *pgx.ConnPool

type DBConnection struct {
	ConnPool *pgx.ConnPool
}

func init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

func NewConnection() *DBConnection {
	return (&DBConnection{ConnPool: cPool})
}

//var name string
//var user_id string
//err = conn.QueryRow(context.Background(), "select name, user_id from users").Scan(&name, &user_id)
//if err != nil {
//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//	os.Exit(1)
//}

//fmt.Println(name, user_id)
