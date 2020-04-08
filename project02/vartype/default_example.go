package main

import "fmt"

//存放单个字母就用byte
//字符串是一串固定长度的字符连接起来，golang的字符中是由字节组成

func main() {
	//所有数据类型都有一个默认类型，如果没有赋值，就会使用这个默认值。默认值又叫零值。
	// int  0
	// float 0
	// bool false
	// string ""

	var v bool
	fmt.Println("v default:",v)

	var aInt int
	fmt.Println("int default:",aInt)
}

