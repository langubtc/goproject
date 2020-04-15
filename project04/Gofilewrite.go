package main

import (
	"log"
	"os"
)

func main() {
	//创建一个文件
	f, err := os.Create("B.txt")
	if err != nil {
		log.Fatal(err)
	}
	//写入内容到文件
	f.WriteString("abdcewsdsd\n")
	f.Seek(2, os.SEEK_SET)
	f.WriteString("oK")
	f.Close()

	//第二种方法使用os.OpenFile方法
	//O_
	v, err := os.OpenFile("c.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	v.WriteString("Vvvv\n")
	v.Close()

}
