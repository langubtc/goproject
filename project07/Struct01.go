package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

func main() {
	p4 := person{
		name: "roddy",
		city: "chengdu",
		age:  28, //不声明值就是零值

	}

	fmt.Println(p4)
	fmt.Printf("%#v\n", p4)

	// 2. 值的列表进行初始化

	p6 := person{
		"roddy",
		"成都",
		28,
	}

	fmt.Println(p6)

}
