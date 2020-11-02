package main

import (
	"go-server/src/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", getSpotifyAuth)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
