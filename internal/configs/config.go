package config

import "os"

var (
	MovieBuffHostURL     string
	MovieBuffAccessToken string
	DatabaseURL          string
	MigrationFilePath    string
	Server               string
	Host                 string
	Port                 string
	Database             string
	User                 string
	Password             string
)

func init() {
	MovieBuffHostURL = os.Getenv("MB_URL")
	MovieBuffAccessToken = os.Getenv("MBAPI_TOKEN")
	DatabaseURL = os.Getenv("DATABASE_URL")
	MigrationFilePath = os.Getenv("MIG_FILES_PATH")
	Server = os.Getenv("DB_SERVER")
	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	Database = os.Getenv("DB_NAME")
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
}
