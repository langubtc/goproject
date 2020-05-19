package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("HELLO", i)
	wg.Done() // 通知wg减1
}

func main() { // 开启一个主goroutine去执行main函数

	wg.Add(100) //计数
	// 启动100个线程
	for i := 0; i < 100; i++ {
		go hello(i) //开启一个goroutine去执行hello函数 有可能不输出Hello
	}
	fmt.Println("start")
	//time.Sleep(time.Second * 1)
	wg.Wait() //等待所有线程都执行完才结束 是阻塞状态
}
