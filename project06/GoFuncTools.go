package main

import "fmt"

//定义函数使用func 关键字
//func  函数名  参数 [参数类型]  返回值类型

func add(x int, y int) int {
	return x + y
}

//定义返回值r(裸返回) syntactic sugar 语法糖
func subtraction(x int, y int) (r int) {

	r = x - y
	return //直接返回会自动返回所定义r变量
}

//省略相同类型的变量类型声明
func getUser(username, password string, age int) string {
	fmt.Println(username, password, age)
	return username

}

//多返回值
func getUserDetail(username, password string, age int) (string, string) {
	fmt.Println(username, password, age)
	return username, password

}

//函数可变参数，args定义参数名  ...int 多个未知参数，返回值为int

func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}

}

//匿名函数func关键字后面没有函数名称
//maper := func(r, rune) rune{
//	return r
//}
//fmt.Println(maper,2)

//return n

//闭包是由函数和与其相关的环境组合而成的实体

func main() {

	result := add(1, 5)
	fmt.Println(result)
	fmt.Println(add(5, 6))

	fmt.Println(subtraction(7, 6))

	fmt.Println(getUser("roddy", "passwrod1", 23))
	fmt.Println(getUserDetail("roddy", "passwrod1", 23))

	//多返回值接收
	name, password := getUserDetail("roddy", "pwd1", 28)

	//使用_放到垃圾桶
	name1, _ := getUserDetail("roddy", "pwd1", 28)
	fmt.Println(name, password)
	fmt.Println(name1)

	fmt.Println(sum(1, 2, 3, 4, 5))

}
