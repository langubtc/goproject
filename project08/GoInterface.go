package main

import "fmt"

func nullInterface(a interface{}) {
	fmt.Println(a)

}

func main() {
	//空接口可以作为函数的参数

	nullInterface("2323323232332")

	//空接口可以赋予任何类型
	mapList := make(map[string]interface{}, 16)
	mapList["name"] = "roddy"
	mapList["age"] = 123
	mapList["list"] = []string{"1", "2", "3"}

	fmt.Println(mapList)

	//接口可以作类型断言
	var a interface{}
	a = "2323"
	ret, ok := a.(string)
	fmt.Println(ret, ok, a)

}
