package main

import (
	"fmt"
	"os"
)

func GetHostName() string {

	echo, _ := os.Hostname()
	return echo

}

func main() {

	var sd = GetHostName()
	var nameString *string = &sd

	a := 100
	//打印内存地址
	fmt.Println(&a)

	//设置一个变量,并且变量的值到&p
	var p = 100

	var c *int = &p

	//*& 显示变量值 *nameString就是显示所对应的值
	fmt.Println(p, *c, *&p, *nameString)

	//porint_name 指向name
	//通过*porint_name 获取值
	var name = "sxxxxxxx"

	var porint_name = &name

	*porint_name = "check" //修改原值，内存地址一样，指向的值会发生变化

	fmt.Println(*porint_name, name, porint_name)

}
