package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//获取文件名
	filename := os.Args[1]
	//fmt.Println(filename)
	//打开文件
	f, _ := os.Open(filename)

	//将文件名交给带有缓存的IO Reader
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n') //指定分割符
		//判断是否是文件结尾
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	_ := f.Close()
}
