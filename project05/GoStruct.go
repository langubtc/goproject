package main

import "fmt"

//结构体定义
//三部分组成  type  名称  struct
//结构体中首字母大写，就是可导出，可引用

type Student struct {
	Id   int
	Name string
	Age  int
}

type Car struct {
	Name  string
	Model string
	Color string
	Price float64
}

func main() {
	var s Student //将S变量定义为Student结构体
	s.Name = "roddy"
	s.Id = 1
	s.Age = 28
	fmt.Println(s)

	var c Car

	c.Name = "宝马"
	c.Color = "白色"
	c.Model = "120i"
	c.Price = 21.04

	fmt.Println(c)

	//第二种方式声明
	audiA4L := Car{
		Name:  "奥迪",
		Color: "黑色",
		Model: "A4L",
		Price: 28.23,
	}
	fmt.Println(audiA4L)

	//通过指针修改结构体里面的值
	var m *Car
	m = &audiA4L
	m.Price = 29.32
	fmt.Println(audiA4L)
}
