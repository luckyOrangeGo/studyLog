# Go learn

```
%d	decimal integer
%x, %o, %b	integer in hexadecimal, octal, binary
%f, %g, %e	floating-point number: 3.141593 3.141592653589793 3.141593e+00
%t	boolean: true or false
%c	rune (Unicode code point)
%s	string
%q	quoted string "abc" or rune 'c'
%v	any value in a natural format
%T	type of any value
%%	literal percent sign (no operand)
```

```go

package main

import "fmt"

//const 未显示设定的第一常量默认设为0，后续同上。
//iota 在 const 里 自动初始为0，后续++
//make new ???

//函数变参   ...int

func biancan(arg ...int) {
}

//延迟语句 defer 在这之后的指定的函数会在函数退出前调用。

//struct中，自定义类型和内置类型可以作为匿名字段，字段重复时，最外层优先访问。
//重写method
type Human struct {
  name  string
  age   int
  phone string
}
type Std struct {
  Human
  school string
}
type Emp struct {
  Human
  company string
}

func (h *Human) SayHi() {
  fmt.Printf("%s, %s/n", h.name, h.phone)
}
func (e *Human) SayHi() {
  fmt.Printf("%s, %s, %s/n", e.name, e.company, e.phone)
}

//mark := Std{Human{.,.,.}, "MIT"}
//sam := Emp{Human{.,.,.}, "PeerSafe"}
//mark.SayHi()
//sam.SayHi()
//----------------interface+++++++++++++++++
type Child interface {
  SayHi()
  Sing(song string)
  BorrowMoney(amount float32)
}
type Elder interface {
  SayHi()
  Sing(song string)
  SpendMoney(amount float32)
}
type Men interface {
  SayHi()
  Sing(song string)
}

//interface 只能通过其他非interface 实现，不能自我实现。
//duck-typing

//传指针比较轻量级（8bytes），
func add1(a *int) int {
  *a = *a + 1
  return *a
}
func add() {
  x := 3
  fmt.Println("x=", x)
  fmt.Println("x+1=", add1(&x))
}
```