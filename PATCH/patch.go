package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://httpbin.org/patch"
	params := map[string]string{
		"name": "updated name",
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonParams))
	if err != nil {
		log.Fatal(err)
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(content))
}
