package main

import (
	"fmt"
	"unsafe"
)

//数据类型，数值类型:整数类型、浮点类型、字符byte(保存单个字母字符)、布尔、字符串
//派生/复杂类型：
//1.指针2.数组  3.结构体 4.管道  5 函数 6.切片  7.接口，8.map



func IntType(){
	//有符号整数类型用于存放整数值  int8 占用1字节  int16 2字节   int32  4字节  int64  8字节

	var i int = 1
	fmt.Println("i =",i)

	//测试int8的范围
	var j int8 = -128
	fmt.Println(j)

	//无符号的类型 uint8  范围是0-255   uint16 是0-2的16次方-1


	//查看变量类型
	fmt.Printf("type:%T\n",j)

	//查看变量占用字节数
	fmt.Printf("占用字节数:%d",unsafe.Sizeof(j))
}

func FloatFunc(){
	//小数类型使用
	var price float32 = 89.123
	fmt.Println("price=",price)

	var num1 float32 = -0.000892323
	var num2 float64 = -7832323.2323

	fmt.Println("num1:",num1,"num2:",num2)

	//精度丢失,float64精度要比float32大，如果要保存一个精度高的浮点数
	//应该选择float64
	var s1 float32 = -123.0000021
	var s2 float64 = -123.000000021
	fmt.Println("s1:",s1,"s2:",s2)

	//golang的浮点型默认是float64

	startNum :=0.2323232
	fmt.Println("start_num:",startNum,)
	fmt.Printf("T:%T\n",startNum)

	//浮点数可以使用.1234 等价于0.1234

	endNum := .2345
	fmt.Println(endNum)



	//科学技术法,使用大E和小e都表示科学计数法
	num8 := 5.1234e2  // 5.1234 * 10 的2次方
	num9 := 5.1234E2  // 5.1234 * 10 的2次方
	num10 := 5.1234E-2  // 5.1234 / 10 的2次方 = 0.051234
	fmt.Println(num8,num9,num10)

}


func main()  {
	//IntType()
	FloatFunc()

}
