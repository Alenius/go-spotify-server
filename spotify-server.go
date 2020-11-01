package main

import (
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const SPOTIFY_BASE_URL = "https://api.spotify.com"
const SPOTIFY_AUTH_URL = "https://accounts.spotify.com/api/token"
const SPOTIFY_CLIENT_ID = "867357abf03643fab0ee2cad0b8903f9"

type Client struct {
	username string
	password string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		req, err := http.NewRequest("POST", SPOTIFY_AUTH_URL, strings.NewReader("grant_type=client_credentials"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		SPOTIFY_API_KEY := os.Getenv("SPOTIFY_AUTH_KEY")
		req.SetBasicAuth(SPOTIFY_CLIENT_ID, SPOTIFY_API_KEY)
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
		req, err = http.NewRequest("GET", SPOTIFY_BASE_URL+"/v1/tracks/11dFghVXANMlKmJXsNCbNl", nil)
		req.Header.Add("Authorization", "Bearer "+access_token)
		res, err = http.DefaultClient.Do(req)
		checkErr(err, w)

		parsedBody, err = parseHttpBody(res.Body)
		log.Println(parsedBody)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func checkErr(e error, w http.ResponseWriter) {
	if e != nil {
		fmt.Fprintf(w, e.Error())
	}
}

func parseHttpBody(body io.ReadCloser) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.NewDecoder(body).Decode(&result)

	return result, err
}
