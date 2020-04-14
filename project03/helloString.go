package main

import (
	"fmt"
	"github.com/langubtc/goproject/gobase01/btfunction"
	"unsafe"
)

func main() {
	cityName := "成都"

	fmt.Println(unsafe.Sizeof(cityName))
	fmt.Println("Hello world!")
	fmt.Println("Hello", cityName)

	//引用github上面的包并且使用
	result := btfunction.Add(2, 3)
	fmt.Printf("github pkg result:%d", result)

}
