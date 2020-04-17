package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//声明一个数组，里面有3个数，类型是int  默认是0
	var a [3]int

	var p [3]int = [3]int{1, 23}

	a = p
	fmt.Println(a == p) //a=p,如果下标和值一样，那就是相等的

	var q [3]string = [3]string{"1", "2", "2"}
	var o [3]string = [3]string{"1", "2"}

	fmt.Println(q == o)

	fmt.Println(p, a)

	//数组数据初始化
	var m [3]string = [3]string{"a", "v", "d"}
	fmt.Println(m)

	//多个值d短变量声明
	f := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(f)

	//只给其中某几个数据赋值
	l := [...]int{0: 1, 1: 2, 5: 6}
	fmt.Println(l)

	var b int32

	fmt.Println(a)

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))

	ipIist := [...]string{"172.16.2.3", "172.16.2.5", "172.16.3.2"}

	for _, v := range ipIist {
		//fmt.Println(v)

		if v == "172.16.2.3" {
			fmt.Println("DBserver")
		} else {
			fmt.Println("Simple")
		}
	}

	cint := "a" //a的acsii 97
	fmt.Printf("%v", []byte(cint))

}
