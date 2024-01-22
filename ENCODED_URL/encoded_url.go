package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	URL, err := url.Parse("https://www.golang.com")
	if err != nil {
		log.Println(err)
		return
	}

	params := url.Values{}
	params.Add("id", "1")
	params.Add("name", "Mubina")

	URL.RawQuery = params.Encode()

	fmt.Println("Encoded URL:", URL)
}
