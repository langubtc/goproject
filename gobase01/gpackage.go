package main

import (
	"fmt"
	//前面bt是包的别名
	bt "goproject/gobase01/btfunction"
	//只执行btinit中的init函数
	_ "goproject/gobase01/btinit"
)
//package 是最基本的分发单位和工程管理中依赖关系的体现
//每个go源码文件开头都要拥有一个package声明，表示源码文件所属代码包
//要生成可执行程序，必须要有一个main的package包，包下有main()函数
//同一个路径下只能存在一个package ，一个package可以拆成多个源文件

//import 有两种形式 1. import "xx"  2. import ("xxx","xxxxx")
//如果依赖包包含B，首先导入B包，然后初始化B包中的常量和变量，最后如果B包中有init，会自动执行init()
//所有包导入完成后才会对main中常量和变量进行初始化
//重复导入包，改包只会被导入一次

// 1.import 别名
// 2. 使用.点来作为别名，如果以. 来导入该包，那就不需要使用包名+函数的方式，而是直接调用函数方法。不建议使用
// 3. 下划线_ 表示导入该包，但不导入整个包，而是执行该包中的init函数，因此无法通过包名来调用包中的其它函数。


func  main()  {
	//导入必须是以GOPATH路径来定义 导入后就可以使用包下面的函数
	var result = bt.Add(1,2)
	var result2 = bt.Subtraction(1,2)
	var result3 = bt.Division(1,2)

	vHostName := bt.HostNmae()
	fmt.Println(result,result2,result3)

	fmt.Print("主机名:",vHostName)

}


