# 有关Go语言的学习

## 20180330Learn 更新

## 学习计划

---
今天主要学习WebSorcket 和 Web加密

## WebSorcket

### 引入

---
Web用户代理（如浏览器）和Web服务器（如Apache）之间的交互标准模型是用户代理发出HTTP请求，服务器对每个请求进行单一回复。在浏览器的情况下，通过单击链接，在地址栏中输入URL，点击向前或向后按钮等来进行请求。响应被视为新页面并被加载到浏览器窗口中。

这种传统模式有许多缺点。首先是每个请求打开并关闭一个新的TCP连接。 HTTP 1.1通过允许持久连接解决了这个问题，因此连接可以在短时间内保持打开状态，以允许在同一台服务器上创建多个请求（例如图像）。尽管HTTP 1.1持久连接缓解了加载包含许多图形的页面的问题，但它并未改进交互模型。即使使用表单，该模型仍然是提交表单并将响应显示为新页面。 JavaScript有助于在提交之前对表单数据执行错误检查，但不会更改模型。

AJAX（异步JavaScript和XML）使用户交互模型取得了重大进展。这允许浏览器发出请求，并使用响应来使用HTML文档对象模型（DOM）更新显示。但是交互模型也是一样的。 AJAX只影响浏览器如何管理返回的页面。在Go for AJAX中没有明确的额外支持，因为不需要任何支持：HTTP服务器只能看到一个普通的HTTP POST请求，可能有一些XML或JSON数据，这可以使用已经讨论过的技术来处理。

所有这些仍然是浏览器（或用户代理）到服务器通信。缺少的是浏览器与服务器建立TCP连接并从服务器读取消息的服务器到浏览器通信。这可以由WebSockets填充：浏览器（或任何用户代理）保持打开与WebSocket服务器的长期TCP连接。 TCP连接允许任何一方发送任意数据包，所以任何应用协议都可以在WebSocket上使用。

如何启动WebSocket是由用户代理发送一个特殊的HTTP请求，指出“切换到WebSocket”。基于HTTP请求的TCP连接保持打开状态，但用户代理和服务器都切换到使用WebSockets协议，而不是获取HTTP响应并关闭套接字。

请注意，它仍然是启动WebSockets连接的浏览器或用户代理。浏览器不运行自己的TCP服务器。虽然作为IETF RFC6455的规范很复杂[请参阅](https://tools.ietf.org/html/rfc6455)，但该协议的设计非常易于使用。客户端打开HTTP连接，然后用自己的WS协议替换HTTP协议，重新使用相同的TCP或新的连接。
Go对子存储库中的WebSocket有一些支持，但实际上建议使用第三方软件包。本章考虑两个软件包。

### 网络协议模型回顾

---
Open Systems Interconnect(OSI)从第一层到第七层分别是：物理层、数据链路层、网络层、传输层、会话层、表示层、应用层。

实际中使用的TCP/IP四层协议，分别是：网络接口层、网间层、传输层、应用层。

socket编程的主要是面向三层和四层的协议，是主要是针对IP协议、TCP、UDP的，像HTTP这种基于TCP的七层协议不在讨论范围内。

### WebSockets服务器

---
一个WebSocket服务器从一个HTTP服务器开始，接受TCP连接并处理TCP连接上的HTTP请求。当一个请求将切换到WebSocket连接的连接时，协议处理程序从HTTP处理程序更改为WebSocket处理程序。所以只有TCP连接的角色发生了变化，服务器继续作为其他请求的HTTP服务器，而基于该连接的TCP套接字被用作WebSocket。

在第8章讨论的简单服务器之一，HTTP注册了各种处理程序，如文件处理程序或函数处理程序。为了处理WebSockets请求，我们只需注册一个不同类型的处理程序 - 一个WebSockets处理程序。服务器使用的处理程序基于URL模式。例如，一个文件处理程序可能被注册为/，/cgi-bin/...的函数处理程序，以及/ws的WebSockets处理程序。仅希望用于WebSockets的HTTP服务器可能运行如下：

``` go
func main() {
    http.Handle("/", websocket.Handler(WSHandler))
    err := http.ListenAndServe(":12345", nil)
    checkError(err)
    }
```

一个更复杂的服务器可能只需添加更多的处理程序就可以处理HTTP和WebSockets请求。

#### Go子存储库包

---
Go有一个名为golang.org/x/net/websocket的子资源库包。 要使用这个，你必须先下载它：
`go get golang.org/x/net/websocket`

软件包文档声明如下：

>这个软件包目前缺少一些替代和更主动维护的WebSockets软件包中的一些功能：[软件包](https://godoc.org/github.com/gorilla/websocket)

它建议你应该会更好地使用替代软件包。尽管如此，我们认为这个软件包与使用Go团队的软件包其余部分一致。后面的部分将介绍另一种软件包。

消息对象HTTP是一个流协议。 WebSockets是基于框架的。 您准备一个数据块（任意大小）并将其作为一组帧发送。帧可以包含UTF-8编码的字符串或字节序列。使用WebSockets最简单的方法就是准备一个数据块，并要求Go WebSockets库将它打包为一组帧数据，通过网络发送，并作为同一块接收。websocket包中包含一个名为Message的便捷对象来完成这个任务。Message对象有两个方法 - 发送和接收 - 它将WebSocket作为第一个参数。第二个参数是要存储数据的变量的地址或要发送的数据。

发送字符串数据的代码如下所示：

```go
msgToSend := "Hello"
err := websocket.Message.Send(ws, msgToSend)

var msgToReceive string
err := websocket.Message.Receive(conn, &msgToReceive)
```

发送字节数据的代码如下所示:

```go
 dataToSend := []byte{0, 1, 2}
 err := websocket.Message.Send(ws, dataToSend)

 var dataToReceive []byte
 err := websocket.Message.Receive(conn, &dataToReceive)

```

接下来给出一个发送和接收字符串数据的回显服务器。 请注意，在WebSockets中，任何一方都可以发起消息发送，并且在此服务器中，当它连接（发送/接收）而不是更常规的接收/发送服务器时，我们会将消息从服务器发送到客户端。服务器是[EchoServer.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/EchoServer.go)

与此服务器通话的客户端是[EchoClient.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/EchoClient.go)

运行方式如下：

`go run EchoClient.go ws://localhost:12345`

客户端的输出是服务器发送的内容：

`Received from server: Hello  0`
`Received from server: Hello  1`
`Received from server: Hello  2`
`Received from server: Hello  3`
`Received from server: Hello  4`
`Received from server: Hello  5`
`Received from server: Hello  6`
`Received from server: Hello  7`
`Received from server: Hello  8`
`Received from server: Hello  9`

#### JSON对象

---
预计许多WebSockets客户端和服务器将以JSON格式交换数据。 对于Go程序，这意味着Go对象将被编组为JSON格式，然后以UTF-8字符串形式发送，而接收器将读取该字符串并将其解组为Go对象。 称为JSON的websocket便利对象将为您执行此操作。 它具有用于发送和接收数据的Send和Receive方法，就像Message对象一样。

考虑某种情况，即客户端使用WebSockets（可以双向发送消息）将Person对象发送给服务器。 从客户端读取消息并将其打印到服务器的标准输出的服务器是[PersonServerJSON.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/PersonServerJSON.go)

以JSON格式发送Person对象的客户端是[PersonClientJSON.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/PersonClientJSON.go)

#### 编解码器类型

---
Message和JSON对象都是Codec类型的实例。 这种类型定义如下：

```go
type Codec struct {
    Marshal   func(v interface{}) (data []byte, payloadType byte, err error)
    Unmarshal func(data []byte, payloadType byte, v interface{}) (err error)
    }
```

编解码器类型实现了之前使用的发送和接收方法。 WebSockets很可能也将用于交换XML数据。 我们可以构建一个XML Codec对象，方法是将XML编组和解组方法包装为一个合适的Codec对象。 我们可以用这种方式创建一个名为XMLCodec.go的XMLCodec包：

```go
package xmlcodec
import (
    "encoding/xml"        "golang.org/x/net/websocket"
)

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
    msg, err = xml.Marshal(v)
    return msg, websocket.TextFrame, nil
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
    err = xml.Unmarshal(msg, v)
    return err
}

var XMLCodec = websocket.Codec{xmlMarshal, xmlUnmarshal}
```

该文件应该安装在GOPATH的src子目录中：
`$GOPATH/src/xmlcodec/XMLCodec.go`

然后，我们可以将诸如Person之类的Go对象序列化为XML文档，并将它们从客户端发送到服务器。 接收文档并将其打印到标准输出的服务器如 [PersonServerXML.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/PersonServerXML.go)

以XML格式发送Person对象的客户端是[PersonClientXML.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/PersonClientXML.go)

#### 基于TLS的WebSockets

---
WebSocket可以构建在安全的TLS套接字之上。 我们讨论了如何使用证书来使用TLS套接字。这在WebSockets中没有改变。也就是说，我们使用http.ListenAndServeTLS而不是http.ListenAndServe。

以下是使用TLS的回显服务器：[EchoServerTLS.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/EchoServerTLS.go)

客户端和以前一样是echo客户端。 所有这些变化都是使用wss模式而不是ws模式的URL：

`EchoClient wss://localhost:12345/`

如果服务器提供的TLS证书有效，那么这将工作正常。 我使用的证书不是：它是自签名的，这通常表示您正在进入危险区域。 如果你仍然想继续下去，你需要打开TLS InsecureSkipVerify标志来使用我们的“去除安全检查”。 这是通过EchoClientTLS.go程序完成的，该程序使用此标志设置配置，然后调用DialConfig代替拨号：
[EchoClientTLS.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/EchoClientTLS.go)

#### HTML页面中的WebSockets

---
WebSockets的原始驱动程序允许HTTP用户代理（如浏览器和服务器）之间进行全双工交互。 预期典型的用例涉及浏览器中的JavaScript程序与服务器交互。 在本节中，我们构建一个Web / WebSockets服务器，该服务器提供一个HTML页面，该页面设置WebSocket并使用WebSockets显示来自该服务器的信息。 我们正在研究图![全双工交互情况](https://github.com/luckyOrangeGo/studyLog/raw/master/WebSockets/Full%20duplex%20interaction%20situation.png)所示的情况。

物联网（IoT）的时代即将到来。因此，我们可以预期来自传感器和传感器网络的数据将用于驱动执行器并在浏览器中显示有关物联网网络的信息。关于使用Raspberry Pi和Arduinos构建传感器网络有很多书，但是我们将通过每隔几秒在网页上更新“传感器”来显示CPU温度，从而大幅度简化情况。来自Debian软件包lm-sensors的Linux传感器命令会向标准输出写入它所知道的传感器的值。我的台式机上的命令传感器产生如下输出：

```log
acpitz-virtual-0 Adapter: Virtual device
temp1:        +27.8°C  (crit = +105.0°C)
temp2:        +29.8°C  (crit = +105.0°C)
coretemp-isa-0000 Adapter: ISA adapter
Physical id 0:  +58.0°C  (high = +105.0°C,
crit = +105.0°C)
Core 0:         +57.0°C  (high = +105.0°C, crit = +105.0°C)
Core 1:         +58.0°C  (high = +105.0°C, crit = +105.0°C)
```

刷新时，通常Core 0和Core 1上的温度可能会发生变化。在Windows上，执行相同操作的命令是这样的：
`wmic /namespace:\\root\wmi PATH MSAcpi_ThermalZoneTemperature get CurrentTemperature`

当它运行时，它的输出如:

`42.4° C`

在Mac上，使用命令[osx-cpu-temp](https://github.com/lavoiesl/osx-cpu-temp)。
如果你不想通过这些步骤，只需要替换一个更普通的程序，如日期。
我们提供了一个Go程序来从ROOT_DIR目录传递HTML文档，然后从URL GetTemp中建立一个WebSocket。 服务器端的WebSocket每两秒从传感器获取输出并将其发送到套接字的客户端。 Web / WebSockets服务器在端口12345上运行，没有特殊原因。 该程序将在安装lm-sensors软件包后在Linux下运行。 对于其他系统，请将任何其他有趣的系统调用替换为exec.Command调用。

Go服务器是[TemperatureServer.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/TemperatureServer.go)

启动此功能的顶级HTML文件是websocket.html

该程序使用JavaScript打开WebSockets连接并处理onopen，onmessage和onclose事件。 它使用evt.data和send函数进行读写。 它以预先格式化的元素显示数据，与上面的数据完全一样。 它每两秒刷新一次。 HTML文档的结构基于HTML5 - [来自TutorialsPoint的WebSockets](https://www.tutorialspoint.com/html5/html5_websocket.htm).

#### The Gorilla Package

---
WebSockets的替代包是github.com/gorilla/websocket包。 要使用它，您需要运行以下命令：

`go get github.com/gorilla/websocket`

#### Echo Server

---
使用这个包的echo服务器是[EchoServerGorilla.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/EchoServerGorilla.go)。 通过引入对websocket.Upgrader对象的调用，它使HTTP到WebSockets的转换更加明确。 它还更清楚地区分发送文本和二进制消息。

使用这个包的echo客户端是[EchoClientGorilla.go](https://github.com/luckyOrangeGo/studyLog/blob/master/WebSockets/EchoClientGorilla.go)

### 结论

---
WebSockets标准是IETF RFC，因此不会有重大变化。 这将允许HTTP用户代理和服务器设置双向套接字连接，并且应该使某些交互样式更容易。Go支持WebSockets的两个包。

### 通用socket编程API

---

``` go
type IP []byte

type IPMask []byte

type IPAddr {IP IP Zone string }

type Conn interface {
    //从链接读数据
    Read(b []byte) (n int, err error)

    //写数据到链接
    Write(b []byte) (n int, err error)

    //关闭链接
    Close() error

    //定位本地网络网址
    LocalAddr() Addr

    //定位远程网络地址
    RemoteAddr() Addr

    // SetDeadline设置与连接关联的读取和写入最后期限。 这等同于调用SetReadDeadline和SetWriteDeadline。
    // 截止日期是I / O操作失败后的绝对时间（参见Error类型），而不是阻塞。 截止日期适用于所有将来和待处理的I / O，而不仅仅是紧接着的读取或写入调用。 超过截止日期后，可以通过设定未来的最后期限来刷新连接。
    // 在成功读取或写入调用后，可以通过重复延长最后期限来实现空闲超时。
    // t的零值意味着I /O操作不会超时。
    SetDeadline(t time.Time) error

    SetReadDeadline(t time.Time) error

    //即使写入超时，也可能返回n> 0，表示某些数据已成功写入。
    SetWriteDeadline(t time.Time) error
}

//保持连接到服务器，即使它没有任何要发送的东西。
func (c *TCPConn) SetKeepAlive(keepalive bool) error



```

`Dial`函数

```go
func Dial(network, address string) (Conn, error)
```

`network` 用于指定网络类型，目前支持的值有：tcp、tcp4(IPv4-only)、tcp6(IPv6-only)、udp、udp4(IPv4-only)、udp6(IPv6-only)、ip、ip4(IPv4-only)、ip6(IPv6-only)、unix、unixgram、unixpacket。

`address` 指定要连接的地址，对于TCP和UDP来说，地址格式为host:port。对于IPv6因为地址中已经有冒号了，所以需要用中括号将IP地址括起来，比如[::1]:80。如果省略掉host的话，比如:80，就认为是本地系统。

例如：

```go
Dial("tcp", "192.168.1.1:80")
Dial("tcp", "golang.org:http")
Dial("tcp", "[2001:db8::1]:http")
Dial("tcp", "[fe80::1%lo0]:80")
Dial("tcp", ":80")
```

如果是IP网络的话，network必须是ip、ip4或ip6，且必须在后面**加上冒号说明协议号或者名字**，比如：

```go
Dial("ip4:1", "192.168.2.1")
Dial("ip6:ipv6-icmp", "2001:db8::1")
```

如果address对应多个地址（比如可能是一个域名对用多个IP），那么Dial会挨个尝试直到成功连上某一个。

除了Dial，还有一个DialTimeout多提供了一个超时的功能：

```go
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
```

第一点需要注意的是正在进行的错误检查的数量过大。这对于联网程序来说是正常的：失败的机会远远大于独立程序。在客户端，服务器或中间的任何路由器和交换机上，硬件可能会失败;通信可能被防火墙阻止;由于网络负载可能会发生超时;服务器可能会在客户端与之通话时崩溃。执行以下检查：

1. 指定的地址中可能存在语法错误。
2. 尝试连接到远程服务可能会失败。例如，所请求的服务可能没有运行，或者可能没有这样的主机连接到网络。
3. 虽然连接已建立，但如果连接突然中断或网络超时，则对服务的写入可能会失败。
4. 同样，读取可能会失败。从服务器读取需要评论。在这种情况下，我们基本上从服务器读取单个响应。这将通过连接上的文件结束来终止。但是，它可能由几个TCP数据包组成，所以我们需要一直读取直到文件结束。

 io / ioutil函数ReadAll将会处理这些问题并返回完整的响应。

## 安全与加密

分布式系统的ISO OSI（开放系统互连）七层模型，ISO在这个架构上构建了一系列文档。对于我们来说，最重要的是ISO安全架构模型ISO 7498-2。这需要购买，但国际电联已制作出一份与之相符的文件X.800，该文件可从国际电联获取，网址为[X.800_zh](https://www.itu.int/ITU-T/recommendations/rec.aspx?rec=3794&lang=zh)

---
功能和级别安全系统所需的主要功能如下：

1. 身份验证/身份证明
2. 数据完整性/数据未被篡改
3. 机密性/数据不公开
4. 公证/签名
5. 访问控制
6. 保证/可用性

### 机制

---
达到这种安全级别的机制如下：

* 对等实体认证
  * 加密
  * 电子签名
  * 认证交换
* 数据来源认证
  * 加密
  * 电子签名
* 访问控制服务
  * 访问控制列表
  * 密码
  * 能力列表
  * 标签
* 连接机密性
  * 加密
  * 路由控制
* 无连接的机密性
  * 加密
  * 路由控制
* 选择性领域的机密性
  * 加密
* 流量机密性
  * 加密
  * 流量填充
  * 路由控制
* 连接完整性与恢复
  * 加密
  * 数据完整性
* 连接完整性无需恢复
  * 加密
  * 数据完整性
* 连接完整性选择字段
  * 加密
  * 数据完整性
* 无连接完整性
  * 加密
  * 数字签名
  * 数据完整性
* 无连接完整性选择字段
  * 加密
  * 数字签名
  * 数据完整性
* 原始不可否认
  * 数字签名
  * 数据完整性
  * 公证
* 收据的不可否认性
  * 数字签名
  * 数据完整性
  * 公证

#### 数据的完整性

---
确保数据完整性意味着提供一种数据未被篡改的测试手段。 通常这是通过在数据中的字节中形成一段简单的数字来完成的。 这个过程被称为散列，结果编号被称为散列值或散列值。

Go支持多种散列算法，包括MD4，MD5，RIPEMD-160，SHA1，SHA224，SHA256，SHA384和SHA512。

一个散列有一个io.Writer，并且你将要散列的数据写入这个作者。 您可以按大小查询哈希值中的字节数，并按和查询哈希值。

一个典型的例子是MD5哈希。这使用了md5软件包。 哈希值是一个16字节的数组。 这通常以ASCII格式打印成四个十六进制数字，每个数字由四个字节组成。 一个简单的程序是[MD5Hash.go](https://github.com/luckyOrangeGo/studyLog/blob/master/Security/MD5Hash.go)。

该程序打印“b1946ac9 2492d234 7c6235b4 d2611184”。 HMAC（Keyed-Hash Message Authentication Code）就是其中的一个变种，它为散列算法增加了一个密钥。使用这个没有什么变化。要与密钥一起使用MD5散列，请将呼叫`hash := md5.New()`替换为:
`hash := hmac.New(md5.New, []byte("secret"))`

#### 对称密钥加密

---
有两种用于加密数据的主要机制。对称密钥加密使用一个加密和解密相同的密钥。加密和解密代理都需要知道这个密钥。没有讨论如何在代理之间传输该密钥。与散列一样，有许多加密算法。现在很多人都知道存在弱点，并且随着计算机变得越来越快，一般来说算法越来越弱。 Go支持多种对称密钥算法，如AES和DES。

算法是块算法。也就是说，他们处理数据块。如果你的数据没有与块大小对齐，你将不得不在末尾添加额外的空白。每个算法都由一个Cipher对象表示。这由NewCipher在适当的包中创建，并将对称密钥作为参数。一旦你有密码，你可以使用它来加密和解密数据块。我们使用AES-128，其密钥大小为128位（16字节），块大小为128位。密钥的大小决定了使用哪个版本的AES。

#### 公钥加密

---
另一种主要的加密类型是公钥加密。 公钥加密和解密需要两个密钥：一个用于加密，另一个用于解密。 加密密钥通常以某种方式公开，以便任何人都可以加密给你的消息。 解密密钥必须保密; 否则，每个人都可以解密这些消息！ 公钥系统是不对称的，对于不同的用途使用不同的密钥。 Go支持很多公钥加密系统。 典型的是RSA方案。 从随机数生成RSA私钥和公钥的程序是[GenRSAKeys.go](https://github.com/luckyOrangeGo/studyLog/blob/master/Security/GenRSAKeys.go)。

#### X.509证书

---
公钥基站（Public Key Infrastructure，PKI）是公钥集合的框架，以及诸如所有者名称和位置等附加信息，以及它们之间的链接，提供某种审批机制。 今天使用的主要PKI基于X.509证书。 例如，Web浏览器使用它们来验证网站的身份。 为我的网站生成自签名X.509证书并将其存储在.cer文件中的示例程序是[GenX509Cert.go](https://github.com/luckyOrangeGo/studyLog/blob/master/Security/GenX509Cert.go)

#### 传输层安全性(TLS)

---
如果你必须自己完成所有繁重的工作，加密/解密方案的用处有限。 目前Internet上支持加密消息传递的最流行机制是TLS（传输层安全性）(Transport Layer Security)，以前称为SSL（安全套接字层）(Secure Sockets Layer)。 在TLS中，客户端和服务器使用X.509证书协商身份。 一旦完成，就在它们之间发明一个秘密密钥，所有的加密/解密都是使用这个密钥完成的。 谈判相对较慢，但一旦完成，将使用更快的秘密密钥机制。 服务器需要有证书; 如果需要，客户可能有一个。

##### 基本客户 A Basic Client

---
我们首先说明如何连接到具有由“知名”证书颁发机构（CA）(Certificate Authority)（如RSA）签名的证书的服务器。 从Web服务器获取头部信息的程序可以调整为从TLS Web服务器获取头部信息。该计划是[TLSGetHead.go](https://github.com/luckyOrangeGo/studyLog/blob/master/Security/TLSGetHead.go)。

在针对诸如 www.google.com:443 的适当网站运行时：

```go
go run TLSGetHead.go www.google.com:443
```

它产生这样的输出：

``` sh
HTTP/1.0 302 Found
Cache-Control: private
Content-Type: text/html; charset=UTF-8
Location: https://www.google.com.au/?gfe_rd=cr&ei=L3lvWKSXMdPr8AfvhqKIBg
Content-Length: 263
Date: Fri, 06 Jan 2017 11:02:07 GMT
Alt-Svc: quic=":443"; ma=2592000; v="35,34"
```

##### 服务器使用自签名证书 Server Using a Self-Signed Certificate

---
如果服务器使用自签名证书（可能在组织内部使用或在试验时使用），那么Go软件包何时会生成错误：“x509：由未知权威机构签名的证书”。 证书必须安装到客户端的操作系统中（这将与O / S相关），或者客户端必须将证书安装为根CA. 我们将展示第二种方式。 使用带有任何证书的TLS的回显服务器是[TLSEchoServer.go](https://github.com/luckyOrangeGo/studyLog/blob/master/Security/TLSEchoServer.go)

如果证书是自签名的，那么简单的TLS客户端将不会与此服务器一起工作，它位于此处。 我们需要将配置设置为TLS.Dial的第三个参数，它将我们的证书安装为根证书。 乔希Bleecher斯奈德在: “[入门X509证书由未知机构签发](https://groups.google.com/forum/#!topic/golang-nuts/v5ShM8R7Tdc)”，显示如何做到这一点。然后，服务器与客户端[TLSEchoClient.go](https://github.com/luckyOrangeGo/studyLog/blob/master/Security/TLSEchoClient.go)工作。

### Security结论

---
安全本身就是一个巨大的领域，本章几乎没有涉及到它。但是，主要概念已被涵盖。没有强调的是设计阶段需要构建多少安全性措施：作为事后考虑的安全性几乎总是失败的。