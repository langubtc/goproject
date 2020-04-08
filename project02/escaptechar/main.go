package main

import "fmt"


func main(){

	//转义字符的使用 \t 制表符号
	fmt.Println("tom\t jack")

	// \n 换行符号
	fmt.Println("mysql\ndba")

	// \转义字符  c:\path\c
	fmt.Println("mysql c:\\path\\c")

	fmt.Println("I say:\"jack  hello\"")

	/** 输出
	姓名	年龄	籍贯	住址
	join	12		河北	北京 **/

	fmt.Println("姓名\t年龄\t籍贯\t住址\njoin\t12\t\t河北\t北京")

	// \r 后面内容替换以前的内容
	fmt.Println("姓名\rtest")
}
