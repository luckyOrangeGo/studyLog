package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://baidu.com",
		"http://qq.com",
		"http://golang.org",
	}
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be donw!")
		return
	}

	fmt.Println(link, "is up!")
}
