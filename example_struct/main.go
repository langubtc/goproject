package main

import "fmt"

type Books struct {
	SN     int64
	Name   string
	Title  string
	Author string
}

func main() {
	newBook := Books{Author: "roddy", SN: 2323232323, Title: "Python", Name: "Python 自动化"}
	fmt.Println(newBook)
}
