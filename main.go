package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    response, err := http.Get("https://newsapi.org/v2/top-headlines?country=us&category=sports&apiKey=565555cabbae4989934af419afac7973")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}