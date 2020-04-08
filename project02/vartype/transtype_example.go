package main

import (
	"fmt"
	_ "unsafe"   //如果我们没有引用一个包，但是又不想去掉，这时候可以在前面加一个下划线表示忽略
)

//golang 类型转换需要显示转换，不能自动转换,数据本身没有转换，只是转换后的值赋予另外一个值
//如果从一个高精度的类型转换成低精度的类型，类型会按溢出处理
func main()  {

	var ains int64 = 999999
	var ains2 int8 = int8(ains)
	fmt.Println(ains2)

	//T(v) T就是int8 float32等  变量就是v
	var i int = 100
	var c float64 = float64(i)

	//低精度到高精度
	var d float32 = 0.2324
	var f float64 = float64(d)



	fmt.Println(c+0.234)
	fmt.Println("s:",f)
	fmt.Printf("%T\n",c)
	fmt.Printf("%T\n",f)
	//转换后我们查看下原有的值
	fmt.Printf("i原来的类型%T\n",i)

}

