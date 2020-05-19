package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() { // 开启一个主goroutine去执行main函数

	wg.Add(100) //计数
	// 启动100个线程
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i)
	}
	fmt.Println("start")
	wg.Wait() //等待所有线程都执行完才结束 是阻塞状态
}
