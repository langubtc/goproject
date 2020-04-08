package main

import "fmt"

func main()  {


	switch "yes" {

	case "yes" :
		fmt.Println("判断为yes")
	case "no":
		fmt.Println("判断为no")
	default:
		fmt.Println("都不满足")

	}

	//判断类型
	var a interface{}
	a ="test"
	switch a.(type) {
	case int:
		fmt.Println("a的类型为Int")
	case float32:
		fmt.Println("a的类型为Float")
	default:
		fmt.Println("其它类型")

	}
	
	
	
}
