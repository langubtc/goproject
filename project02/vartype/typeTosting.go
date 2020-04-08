package main

import (
	"fmt"
	"strconv"
)

func main()  {

	// 在开发中，经常会将基本数据类型转成string 或者将string转成基本数据类型

	var i  = 10
	var c   = true
	var s string = "23"
	var d  = 0.2345

	//将s转换成10进制 类型
	var n3 int64 = 11
	n3,_ = strconv.ParseInt(s,10,64)
	fmt.Println(n3)



	// 使用方式
	var str string
	// %d  整数
	// %f  float
	// %t  bool
	// %c  string


	str=fmt.Sprintf("%d",i)

	//f=fmt.Sprintf("%d",i)
	fmt.Printf("int to string:%q\n",str)

	str=fmt.Sprintf("%f",d)
	fmt.Printf("float to string:%q\n",str)

	str=fmt.Sprintf("%t",c)
	fmt.Printf("bool to string:%s\n",str)

}