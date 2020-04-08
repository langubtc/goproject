package main

import "fmt"

func main()  {

	//for {
	//	fmt.Println("我是循环<<<")
	//	time.Sleep(1*time.Second) //每秒打印
	//}

	for i:=1;i<10;i++{
		fmt.Println("循环10次")
	}

	//定义一个字符串序列
	a:= [] string{"香蕉","苹果","雪梨"}
	for key,value:=range a{
		fmt.Println(key,value)
	}

	//定义一个int序列,用下滑线代替key
	b:= [] int{1,2,3,4,5}
	for _,value:=range b{
		fmt.Println(value)
	}

	goto One
	fmt.Println("中间代码块")

	One:
		fmt.Println("这里是代码块1")
	//goto One 无限循环

	for i:=1;i<=9;i++{
		for c:=1;c<=i;c++{
			fmt.Printf("%d * %d=%d  ",c,i,i*c)
		}
		fmt.Println("")
	}
}
