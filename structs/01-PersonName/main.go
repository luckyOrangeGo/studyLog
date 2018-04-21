package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// var alex person
	alex := person{
		firstName: "埃里克森",
		lastName:  "安德森",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 92919,
		}, //此处必须有逗号
	}
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	alex.print()
	// (&alex).updateName("艾利克斯")
	alex.updateName("艾利克斯")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
} //如果直接传递值而不是指针，新命名不管用

func (p person) print() {
	fmt.Printf("%+v", p)
}
