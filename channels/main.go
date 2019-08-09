package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, Link := range links {
		go checkLink(Link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	time.Sleep(time.Second)
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link

}
