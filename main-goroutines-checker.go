package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	// results["hello"] = "Hello"
	// fmt.Println("results : ", results)
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

	// fmt.Println(results)
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }
}

var errRequestFailed = errors.New("Request Failed")

//<- 를 넣으면 Send only
// func hitURL(url string, c chan<- result) {
// 	fmt.Println("Checking : ", url)
// 	// c <- result{}
// 	fmt.Println(<-c)
// }

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
