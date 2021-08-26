package config

import "os"

var (
	MOVIEBUFF_HOST_URL     string
	MOVIEBUFF_ACCESS_TOKEN string
	DatabaseURL            string
	MigrationFilePath      string
)

func init() {
	MovieBuffHostURL = os.Getenv("MB_URL")
	MovieBuffAccessToken = os.Getenv("MBAPI_TOKEN")
	DatabaseURL = os.Getenv("DATABASE_URL")
	MigrationFilePath = os.Getenv("MIG_FILES_PATH")
}
