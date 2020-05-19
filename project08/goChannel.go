package main

import "fmt"

func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {

		ch1 <- i
	}

	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	//从通道中取值1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		} else {
			ch2 <- tmp * tmp
		}
	}
	close(ch2)

}

func main() {
	// var 变量  chan  元素类型
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	//从通道中取值2

	for ret := range ch2 {
		fmt.Println(ret)
	}

}
