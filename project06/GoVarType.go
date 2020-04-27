package main

import "fmt"

// 外部变量
var (
	Name string = "default"
	Age  int64
)

func change() {
	//修改的就是全局变量
	Name = "1234"
	a := 124
	fmt.Println(a)
}

func main() {

	Age = 100
	fmt.Println(Name, Age)

	change()
	Age = 100
	fmt.Println(Name, Age)

}
