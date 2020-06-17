package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	// var results map[string]string // v초기화 되지않은 거에 값을 넣으려고 하면 패닉
	//var results = map[string]string{} // v이런식으로 해줘야 넣을 수있음
	var results = make(map[string]string)
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	// fmt.Println(results)
	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
