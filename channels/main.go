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

	c := make(chan string) //在channel 中

	for _, link := range links {
		go checkLink(link, c) // 在函数前面加go 使其变成独立线程
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //单线程在这里卡住了
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
