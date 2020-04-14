package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	str1 := "this is roddy"

	str2 := str1 + "\nstr2"
	// 字符串"-42"转成-42 int

	var stringRead = "-42"

	i, _ := strconv.Atoi(stringRead)

	//将string浮点数转换成64位
	f, _ := strconv.ParseFloat("2.4523", 64)

	//将string转成int64位
	i64, _ := strconv.ParseInt("-42", 10, 64)

	fmt.Println(str1, str2, i, reflect.TypeOf(i))
	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(i64, reflect.TypeOf(i64))

}
