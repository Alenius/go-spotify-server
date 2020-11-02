package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
