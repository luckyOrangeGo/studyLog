package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const FILE = "./test.txt"

func main() {
	dstr := []byte("this is a test")
	err := ioutil.WriteFile(FILE, dstr, 0777)
	if err != nil {
		fmt.Println("write file err")
	}
	dstr = []byte("this is a test")
	fd, err := os.OpenFile(FILE, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open file error")
	}
	defer fd.Close()
	num, err := fd.Write(dstr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
