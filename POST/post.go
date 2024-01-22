package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	JSON struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"json"`
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func main() {
	URL := "https://httpbin.org/post"

	user := User{
		ID:   10,
		Name: "Mubina",
	}

	jsonReq, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	response, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	var result Result
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatal(err)
		return
	}

	pp.Println(result)

}
