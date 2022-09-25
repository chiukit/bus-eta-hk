package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	// read response body
	body, ioErr := ioutil.ReadAll(res.Body)
	if ioErr != nil {
		fmt.Println(ioErr)
	}
	// close response body
	res.Body.Close()
	return body
}

func WriteJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}
