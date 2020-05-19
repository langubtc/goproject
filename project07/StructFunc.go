package main

import "fmt"

type Person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) Person {

	//返回person结构体
	return Person{
		name: name,
		city: city,
		age:  age,
	}

}

func newPerson01(name, city string, age int8) *Person {

	//返回person结构体指针程序优化
	return &Person{
		name: name,
		city: city,
		age:  age,
	}

}

func main() {

	a := newPerson("roddy", "成都", 29)
	av := newPerson01("roddy01", "成都2", 29)
	fmt.Printf("%#v\n", a)
	fmt.Println(av)

}
