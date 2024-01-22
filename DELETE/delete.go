package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

func main() {
	url := "https://httpbin.org/delete"
	request, err := http.NewRequest(http.MethodDelete, url, nil)
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

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	pp.Println("Successfully deleted!")
	fmt.Println(string(data))
}
