package main

import "fmt"

func main()  {
	// 变量, 相当于内存中一个数据存储空间，变量名可以访问到具体的变量值
	// 步骤 1. 声明  2.赋值  3.使用

	// golang使用var 来定义一个变量, 如果不赋值就是默认值,int的默认值为0
	var userAge int
	fmt.Println("value:",userAge)

	//<<<<<<<<<<<<<<<<<<<<
	// 声明变量使用var 来定义变量，然后是变量名，后面必须声明变量类型
	var age int

	age = 17

	fmt.Println("年龄:",age)

	//<<<<<<<<<<<<<<<<<<<<

	//类型推导，根据变量类型自动判断

	var School = "成都七中"
	var name = "roddy"
	var UserAge = 27
	fmt.Println(School,name,UserAge)

	//<<<<<<<<<<<<<<<<<<<<

	//省略var 定义方式  直接使用:= 来定义变量
	UserName := "Tom"
	fmt.Println(UserName)



}
