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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(statusCode)
	gw := gzip.NewWriter(w)
	defer gw.Close()
	json.NewEncoder(gw).Encode(body)
}
