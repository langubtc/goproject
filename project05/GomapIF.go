package main

import "fmt"

func main() {

	newMap := make(map[string]bool)

	newMap["new"] = true
	newMap["a"] = true

	// !只能用于bool类型判断
	if !newMap["b"] {
		fmt.Println("BBBBBB")
	}

}
