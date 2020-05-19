package main

import (
	"fmt"
	"os"
)

func showMenu() {

	fmt.Println("欢迎使用学生管理系统")
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生")
	fmt.Println("3.展示所有学生")
	fmt.Println("4.退出系统")
}

func main() {
	for {
		showMenu()
		var input int
		fmt.Scanf("%d\n", &input) //扫描用户输入

		switch input {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			os.Exit(0)
		default:
			showMenu()
		}
	}

}
