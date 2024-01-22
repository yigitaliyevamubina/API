package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

type Output struct {
	Headers struct {
		Encoding  string `json:"Accept-Encoding"`
		Host      string `json:"Host"`
		UserAgent string `json:"User-Agent"`
		ID        string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
}

func main() {
	URL := "https://httpbin.org/headers"
	response, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var result Output
	if err := json.Unmarshal(content, &result); err != nil {
		log.Println(err)
		return
	}

	pp.Println(result)
}
