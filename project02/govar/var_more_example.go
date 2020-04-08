package main

import "fmt"

//定义全局变量,凡是在函数以外的变量就叫全局变量
var m1 = 100


//定义多个全局变量
var (
	m3 = 300
	m4 = 400
	name2= "Monty"
)



func main() {
	// 声明多个变量后再赋值

	var u1, u2, u3 int
	u1 = 2
	u2 = 3
	u3 = 4

	fmt.Println("value:", u1, u2, u3)

	// 声明和赋值一起,并且不用声明其类型
	var l1, l2, l3 = 1, "roddy", "golang"
	fmt.Println(l1, l2, l3)

	// 使用:=声明多个变量

	a, b, c := 1, 2, 3

	fmt.Println(a, b, c,m1,m3,m4,name2)

	//在一个代码块中,不能定义一个重复的变量名   name redeclared
	//var name= "a"

	var name= "c"
	//可以修改name的值，但是不能在代码块中设置同一个变量名
	name="v"

	fmt.Println(name)

	//如果变量都是数值的话，使用+号就是做数值运算
	var ag1,ag2  = 1,2
	var sg1,sg2 = "string","string2"

	fmt.Println(ag1+ag2)
	fmt.Println(sg1+sg2)

}
