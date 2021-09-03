package moviebuff

import (
	"fmt"
	"os"

	"github.com/RealImage/moviebuff-sdk-go"
	"github.com/sirupsen/logrus"
)

func GetMovie(movie_id string) (*moviebuff.Movie, error) {

	fmt.Println("Fetch from MB")
	cfg := moviebuff.Config{
		HostURL:     os.Getenv("MB_URL"),
		StaticToken: os.Getenv("MBAPI_TOKEN"),
	}
	movieData, err := moviebuff.New(cfg).GetMovie(movie_id)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to get Movie Information from Moviebuff for given MovieID")
	}
	return movieData, err
}
