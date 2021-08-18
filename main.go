package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func main() {
  log.WithFields(log.Fields{
    	"Is server started": "True",
  }).Info("Go app server started")
  http.HandleFunc("/ping", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
    	"Request Received": "Yes",
  	}).Info(r)
	fmt.Fprintf(w, "200")
}