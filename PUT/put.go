package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	user := User{
		Id:   1,
		Name: "Mubina",
	}
	url := "https://httpbin.org/put"
	jsonUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonUser))
	if err != nil {
		log.Panic(err)
	}

	request.Header.Set("Content-type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	pp.Println("Result: ")

	fmt.Println(string(data))

}
