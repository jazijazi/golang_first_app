package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func createUserRequest() {
	url := "http://localhost:8080/users/"

	jsonBody := []byte(`{"name":"jazi22","password":"ab12345"}`)
	bodyReader := bytes.NewReader(jsonBody)

	res, err := http.Post(url, "application/json", bodyReader)
	// defer res.Body.Close()
	fmt.Println(err)
	json_response, _ := io.ReadAll(res.Body)
	fmt.Println(string(json_response))

}

func login() {
	url := "http://localhost:8080/users/login/"

	jsonBody := []byte(`{"name":"jazi22","password":"ab12345"}`)
	bodyReader := bytes.NewReader(jsonBody)

	res, err := http.Post(url, "application/json", bodyReader)
	// defer res.Body.Close()
	fmt.Println(err)
	json_response, _ := io.ReadAll(res.Body)
	fmt.Println(string(json_response))
}

func verify() {
	url := "http://localhost:8080/users/verify/"

	jsonBody := []byte(`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQyOTExNDMsIm5hbWUiOiJqYXppMjIiLCJyb2xlIjoiIn0.xvkQOL7K1gOt_UeOOWuvIzhB_HMMcFEf4Y3VnZqPsZU"}`)
	bodyReader := bytes.NewReader(jsonBody)

	res, err := http.Post(url, "application/json", bodyReader)
	// defer res.Body.Close()
	fmt.Println(err)
	json_response, _ := io.ReadAll(res.Body)
	fmt.Println(string(json_response))
}

func main() {
	// login()
	verify()
}

// TODO
// make air port
// what is clime
// clean with chatgpt
