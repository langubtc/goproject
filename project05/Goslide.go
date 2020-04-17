package main

import "fmt"

func main() {
	newPrice := [...]string{"77", "99", "vsdsd", "99"}
	//切片从0开始，不包括最后一位
	fmt.Println(newPrice[:4])

	price := [6]int{23, 45, 235, 67, 21}

	//切片从0开始,不包括最后一位
	fmt.Println(price[1:4])
	fmt.Print(price)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2] //John  Paul
	b := names[1:3] //Paul Georage

	fmt.Println(a, b)

	b[0] = "xxx" //	XXX GEORAGE 修改一个后两个序列都会改,因为用的是同一个数组

	fmt.Println(a, b)
	fmt.Println(names)

}
