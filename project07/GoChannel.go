package main

import "fmt"

//通道有数据进，还有出
//通道有缓冲区

func main() {
	// 创建通道
	ch := make(chan int)
	defer close(ch)
	go func() {
		ch <- 3 + 2 //向通道发送数据
		ch <- 3 + 4 //向通道发送数据
	}()
	//接收通道数据
	val := <-ch
	val2 := <-ch
	fmt.Println(val2, val)

}
