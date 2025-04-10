package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/users/"

	jsonBody := []byte(`{"name":"sendrequest","password":"123"}`)
	bodyReader := bytes.NewReader(jsonBody)

	res, err := http.Post(url, "application/json", bodyReader)
	// defer res.Body.Close()
	fmt.Println(err)
	json_response, _ := io.ReadAll(res.Body)
	fmt.Println(string(json_response))

}
