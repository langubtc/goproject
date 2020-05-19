package main

import (
	"fmt"
	"reflect"
)

func main() {

	type RoddyInt int //自定义类型  就是RoddyInt

	var a RoddyInt

	type CCInt = int //类型别名  还是int

	var b CCInt

	a = 2

	b = 3

	fmt.Println(a, b, reflect.TypeOf(a), reflect.TypeOf(b))
	//输出结果  2 3 main.RoddyInt int

}
