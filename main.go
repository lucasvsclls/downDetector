package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.golang.org",
		"http://www.amazon.com",
		"http://www.microsoft.com",
		"http://www.facebook.com",
		"http://www.claude.com",
		"http://www.chatgpt.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(4 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	message := link

	if err != nil {
		message += " might be down"
		fmt.Println(message)

		c <- link
		return
	}

	message += " is available"
	fmt.Println(message)

	c <- link
	return
}
