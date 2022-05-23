package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"https://www.google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://golang.org",
		"https://stackoverflow.com"}

	c := make(chan string)
	for _, url := range websites {
		go checkUrl(url, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkUrl(link, c)
		}(l)
	}
}

func checkUrl(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("%v is up\n", link)
	}
	c <- link
}
