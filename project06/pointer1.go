package main

import "fmt"

func incr(p, b *int) {
	*p++
	*b = *b + 30
}

func main() {
	x := 1
	c := 2
	p := &x            //&获取x的地址
	fmt.Println(x, *p) //*p获取x所对应的值

	*p = 3 //修改变量的值

	fmt.Println(x, *p)

	//将x地址传递给incr函数，然后返回变量修改
	incr(&x, &c)
	fmt.Println(x, c)

	name := new(string)
	server := new(string)
	fmt.Println(name, server)

	mm := make(map[string]string)
	mm["name"] = "lc"
	fmt.Println(mm["name"], mm)

	a := new(int)

	*a = 2

	fmt.Println(*a)

}
