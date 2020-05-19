package main

import "fmt"

func main() {

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {

		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println(i, "232323")

		default:
			fmt.Println("不做操作")
		}
	}
}
