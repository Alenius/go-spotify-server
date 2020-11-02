package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const spotifyBaseUrl = "https://api.spotify.com"
const spotifyAuthUrl = "https://accounts.spotify.com/api/token"
const spotifyClientId = "867357abf03643fab0ee2cad0b8903f9"

func getSpotifyAuth(w http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest("POST", spotifyAuthUrl, strings.NewReader("grant_type=client_credentials"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	spotifyApiKey := os.Getenv("SPOTIFY_AUTH_KEY")
	req.SetBasicAuth(spotifyClientId, spotifyApiKey)
	checkErr(err, w)

	res, err := http.DefaultClient.Do(req)
	checkErr(err, w)

	parsedBody, err := parseHttpBody(res.Body)
	checkErr(err, w)

	if res.StatusCode != 200 {
		log.Fatalf("No auth")
	}

	fmt.Fprintf(w, "Login success")

	var access_token string = parsedBody["access_token"].(string)
	req, err = http.NewRequest("GET", spotifyBaseUrl+"/v1/tracks/11dFghVXANMlKmJXsNCbNl", nil)
	req.Header.Add("Authorization", "Bearer "+access_token)
	res, err = http.DefaultClient.Do(req)
	checkErr(err, w)

	parsedBody, err = parseHttpBody(res.Body)
	log.Println(parsedBody)
}
