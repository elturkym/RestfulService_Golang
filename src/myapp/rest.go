package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "./domains"
	"../daos"
)



func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func getJsonResponse() ([]byte, error) {

	return json.MarshalIndent(daos.GetPostById(1), "", "  ")
}
