package main

import (
	"fmt"
	"os"
)

func Info() {

	fmt.Printf("<<<<<帮助信息<<<<<\n1.start\t启动\n2.stop\t停止\n3.restart\t重启")

}

func Menu(Input string) {

	switch Input {

	case "start":
		fmt.Println("<<<<<<start<<<<<<")
	case "stop":
		fmt.Println("<<<<<<stop<<<<<<")
	case "restart":
		fmt.Println("<<<<<<stop<<<<<<")
		fmt.Println("<<<<<<start<<<<<<")
	case "help":
		Info()
	default:
		Info()
	}

}

func main() {
	// os.args 可以获取用户输入参数
	//for i:=0;i<len(os.Args);i++{
	//	fmt.Println(os.Args[i])
	//}

	//根据用户数据参数进行不同的方法，如果不输入则提示帮助信息
	if len(os.Args) <= 1 {
		fmt.Println("脚本：缺少参数!")
		os.Exit(1)
	} else {
		inputArgs := os.Args[1]

		Menu(inputArgs)

	}

}
