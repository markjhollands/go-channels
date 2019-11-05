package main

import (
	"log"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.co.uk",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// An alternative to the loop above we can simplify
	// to this code below:

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		log.Println(link, "might be down!")
		c <- link
		return
	}

	log.Println(link, "is up!")
	c <- link
}
