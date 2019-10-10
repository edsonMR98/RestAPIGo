package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func splitDescription(d string) string {
	sliceOfD := strings.Split(d, " ")
	return strings.Join(sliceOfD[:25], " ")
}

func main() {
	response, err := http.Get("https://newsapi.org/v2/top-headlines?country=us&category=sports&apiKey=565555cabbae4989934af419afac7973")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body) // return data in []byte

		// Unmarshal JSON to News struct.
		var news News
		json.Unmarshal(data, &news) // convert data ([]byte) to News struct

		fmt.Printf("Status = %v", news.Status)
		fmt.Println()
		fmt.Printf("Total Results = %v", news.TotalResults)
		fmt.Println()

		for x := 0; x < len(news.Articles); x++ { // information of all articles provided by API is displayed
			fmt.Println("Article", x+1)
			fmt.Println("Source{ Id:", news.Articles[x].Source.Id, "	Name:", news.Articles[x].Source.Name, "}")
			fmt.Println("Author:", news.Articles[x].Author)
			fmt.Println("Title:", news.Articles[x].Title)
			
			description := news.Articles[x].Description
			if len(strings.Split(news.Articles[x].Description, " ")) > 25 { // If the Description has more than 25 words...
				description = splitDescription(news.Articles[x].Description)
				fmt.Println("Description:", description, "...")
			} else {
				fmt.Println("Description:", description)
			}
			
			fmt.Println("")
		}
	}
}