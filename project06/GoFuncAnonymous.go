package main

import "fmt"

//全局匿名函数
var (
	getResult = func(a, b int, meth_type string) int {
		if meth_type == "+" {
			return a + b
		} else if meth_type == "-" {

			return a - b
		} else if meth_type == "*" {
			return a * b
		} else {
			return 1
		}
	}
)

func funcMethod(a, b int) (int, int, int) {
	//加法
	add := getResult(a, b, "+")
	//减法
	sub := getResult(a, b, "-")
	//乘法
	mul := getResult(a, b, "*")

	return add, sub, mul

}

func main() {

	fmt.Println(funcMethod(1, 23))

}
