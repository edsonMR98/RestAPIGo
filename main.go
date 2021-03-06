package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Sources define the struct of the source of each article
type Sources struct {
	Id   string
	Name string
}

// Articles define the struct of each article
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

// News define the API´s news struct
type News struct {
	Status       string
	TotalResults int
	Articles     []Articles
}

// splitDescription, take a description with +25words and split it showing the first 25 words
// d: string description of each article
// retun string with 25 words
func splitDescription(d string) string {
	sliceOfD := strings.Split(d, " ")
	return strings.Join(sliceOfD[:25], " ")
}

func main() {
	response, err := http.Get(URL_API_KEY)
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
