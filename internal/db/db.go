package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	config "github.com/PoombavaiS/MyfirstGo/internal/configs"
	"github.com/golang-migrate/migrate/v4"
	pgx "gopkg.in/jackc/pgx.v2"
)

var cPool *pgx.ConnPool

type DBConnection struct {
	ConnPool *pgx.ConnPool
}

func init() {
	fmt.Println("Db Init called")

	var err error

	p, _ := strconv.ParseUint(config.Port, 0, 16)
	config := pgx.ConnConfig{
		User:     config.Server,
		Password: config.Password,
		Host:     config.Host,
		Port:     uint16(p),
		Database: config.Database,
	}
	cPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 20,
	})
	if err != nil {
		log.Fatalf("Connection error: %s", err)
	}

	// os.Setenv("DATABASE_URL", "postgres://poombavai:poombavai@localhost:5432/go_training_db?sslmode=disable")
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())
}

func NewConnection() *DBConnection {
	return (&DBConnection{ConnPool: cPool})
}

func Migrations() {
	m, err := migrate.New(os.Getenv("MIGRATION_FILES"), os.Getenv("DATABASE_URL"))
	fmt.Println(os.Getenv("MIGRATION_FILES"))
	fmt.Println(os.Getenv("DATABASE_URL"))
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
