package main

import "fmt"

// bool 占用1个字节，只允许true 和false
// 主要用于if  and for 做流程和条件判断
// 适用于逻辑运算

func main() {
	var b = false
	if b {
		fmt.Println("tst")
	}

}
