# studyLog

## 20180329Learn

### 学习计划

复习structs, map 的更多用法，复习GO Web ，掌握使用http包，了解Socket编程。

### 主要知识点

1. http包建立Web服务器
[Web.go](https://github.com/luckyOrangeGo/studyLog/tree/master/HttpPg)
2. Socket编程

``` go
type IP []byte       //定义IP

ParseIP(s string) IP //IP转化

func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error) //客户端和服务器交互通道
```

>`DialTCP` 建立TCP连接

>`ListenTCP` 接收来自客户端请求

>`SetTimeout` 设置连接的超时时间

>`SetKeepAlive` 保持持续连接

UDP与TCP的区别在于服务端处理多个客户端请求数据包的方式不同，UDP缺少了客户端连接请求的Accept函数，其他几乎一样，将TCP换成UDP即可。

### 指令用法积累

panic

* 停止当前函数
* 一直向上返回，执行每一层的defer
* 如果没有遇见recover, 程序退出

recover

* 仅在defer 调用中使用
* 获取panic 的值
* 如果无法处理，可重新panic

``` go
defer func() {
        r := recover()
        if r == nil {
        fmt.Println("Nothing to recover. " +
            "Please try uncomment errors " +
                "below.")
            return
        }
        if err, ok := r.(error); ok {
            fmt.Println("Error occurred:", err)
        } else {
            panic(fmt.Sprintf(
                "I don't know what to do: %v", r))
        }
    }()
```
