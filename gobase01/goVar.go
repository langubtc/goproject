package main

import (
	"fmt"
	"goproject/gobase01/btvar"
	"reflect"
	"unsafe"
)

//变量的声明格式 var 变量名  类型
//变量声明及赋值 var 变量名  类型  值
//分组声明格式：
// var (
// 		i int
// 		j float32
// )


//冒号赋值只能用于函数体内

//特殊变量下划线_  可以把_理解为垃圾桶

//类型转换只能发生在两种兼容类型之间
//类型转换格式: 变量名  [:] = 目标类型
//大写字母开头的变量是可导出的，也可以在其它包读取，
//小写字母是私有变量，不可以导出。也就是不可以在其它包读取

// const 常量定义形式和变量类型范围
// const 定义可以分为显示和隐式:
	// const identifier [type] =value (显式)
	// const identifier = value  （隐式、也叫无类型常量）
// iota 特殊常量的使用


const imooc string = "moke"

const name  = "q我的名字"


const (
	cat string = "mao"
	doge string = "gou"
	fly = "fei"
)

// iota在const关键字出现时将被重置为0,const中每新增一行常量声明将使iota计数一次
// iota 常见使用法：1.跳值使用法，2.插队使用法  3.表达式隐式使用法，4.单行使用法

const ab1 = iota


//每新增加一行常量自增1，从0开始
const (
	abc = iota
	abd = iota
	_    //跳值使用法，直接使用下划线 下面abf就会变为3
	abf = iota

	abe = iota * 7    //自动隐式继承
	abm = iota *10
	abp
)


var a int

var b int = 10

var (
	c int
	f float32
	d bool
)


//单行赋多个值, 可以不声明类型，直接使用等于即可，go会自动判断变量类型
var i,j,l =1,2,3

var Name = "Roddy"

var school,age,_="cs school",23,"error" //下划线就相当于丢入垃圾桶,其它地方无法调用

func main(){
	//:简写声明必须在函数体中
	m := "cccc"

	// 类型转换
	var mcool int = 3

	cool := float32(mcool)

	//默认值为0
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(d)
	fmt.Println(b)
	fmt.Println(i,j,l)
	fmt.Println(Name)
	fmt.Println(m)

	fmt.Println(school,age,cool)

	fmt.Println(unsafe.Sizeof(cool),reflect.TypeOf(cool))

	fmt.Println(btvar.School)
	fmt.Println(imooc)
	fmt.Println(doge,fly,cat)
	fmt.Println(ab1)
	fmt.Println(abc,abd,abf,abe,abm,abp)


}