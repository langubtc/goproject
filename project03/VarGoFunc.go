package main

import "fmt"

func main() {

	var a int = 1

	var Check = true

	var c_back, b_back, d_back = 1, 2, 3

	Disk, Mem, Cpu := "500G", "2G", "8核"

	if Check == true {
		fmt.Println("检查OK<<<")
	}

	fmt.Println(a, c_back, b_back, d_back)

	fmt.Println("Disk:", Disk, "Mem:", Mem, "Cpu:", Cpu)

}
