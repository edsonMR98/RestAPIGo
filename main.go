package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sources struct {
	Id   string
	Name string
}

type Articles struct {
	Source      Sources
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

type News struct {
	Status       string
	TotalResults int
	Articles     []Articles
}

func main() {
	response, err := http.Get("https://newsapi.org/v2/top-headlines?country=us&category=sports&apiKey=565555cabbae4989934af419afac7973")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body) // return data in []byte
		fmt.Println(string(data))

		// Unmarshal JSON to News struct.
		var news News
		json.Unmarshal(data, &news) // convert data ([]byte) to News struct

		fmt.Printf("Status = %v", news.Status)
		fmt.Println()
		fmt.Printf("Total Results = %v", news.TotalResults)
		fmt.Println()
		fmt.Printf("Total Results = %v", news.Articles[3].Author)
		fmt.Println()
	}
}
