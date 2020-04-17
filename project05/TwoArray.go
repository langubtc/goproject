package main

import "fmt"

func main() {

	//定义二维数组
	var oneList [3][3]int

	fmt.Println(oneList)
	//定义二维数组并赋值
	a := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(a)
	// 修改数组中的某个值
	a[0][1] = 100

	fmt.Println(len(a), len(a[0]))
	//循环遍历数组
	for _, v := range a {

		for _, c := range v {
			fmt.Println(c)
		}
	}
}
