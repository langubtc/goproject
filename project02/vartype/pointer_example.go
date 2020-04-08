package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// 基本数据类型在内存的布局
	var i int = 10


	var ptr *int  = &i
	// ptr是指针变量
	// ptr 的类型的*int
	// ptr 本身的值是&i

	blkBits_,ok:=strconv.ParseUint("387067068",16,32)

	fmt.Println(blkBits_,ok)

	fmt.Println("变量:",&i)
	fmt.Println("指针变量:",ptr)
	fmt.Printf("指针地址本身的值:%v\n", &ptr)
	fmt.Printf("指针指向的值:%v\n", *ptr)

	// 案例：写一个程序，获取一个int变量num的地址，并显示到终端将num的地址赋值给指针ptr,并通过ptr去修改num的值

	var num = 100
	fmt.Println("num的地址是:",&num)

	var ptm *int
	ptm  = &num
	*ptm = 10
	fmt.Println("num = ", num)



}
