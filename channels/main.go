// 访问域名并打印可访问性，和访问次数

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
		"http://baidu.com",
		"http://qq.com",
		"http://golang.org",
	}

	// channel
	c := make(chan string)

	// 计数器
	m := make(map[string]int)

	// 独立进程分配
	for _, link := range links {
		m[link] = 0
		go checkLink(link, c, m) // 在函数前面加go 使其变成独立线程

	}

	//延迟1秒后重复访问
	for l := range c {
		go func(link string) { //function literal 需要传入的参数
			time.Sleep(1 * time.Second)
			checkLink(link, c, m)
		}(l) //实际传入的参数

	}

}

// 检查链接可访问性，并记录访问次数
func checkLink(link string, c chan string, m map[string]int) {
	_, err := http.Get(link) //单线程在这里卡住了
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	//计数器++
	m[link]++

	fmt.Println(link, "is up!", m[link])
	c <- link
}
