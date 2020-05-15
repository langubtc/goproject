package main

import "fmt"

func getname() string {

	return "name"
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("触发宕机捕获异常:", err)
		}
	}()

	fmt.Println("starting......!")

	a := 23
	if a == 23 {
		panic("crash") //手动触发宕机,后面的代码不会执行
	}
	fmt.Println(getname())

}
