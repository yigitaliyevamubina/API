package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/k0kubun/pp"
)

func main() {
	formData := url.Values{}
	formData.Add("id", "1")
	formData.Add("name", "Mubina")
	formData.Add("surname", "Yigitaliyeva")
	formData.Add("group", "Golang Backend fn8")

	url := "https://httpbin.org/post"
	response, err := http.PostForm(url, formData)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	var result map[string]interface{}
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
