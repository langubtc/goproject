package main

import (
	"fmt"
	"time"
)

func PrintServer(s string) {
	for i := 0; i <= 100; i++ {

		time.Sleep(time.Second * 2)
		fmt.Println("server:", s)
	}
}

func main() {
	go PrintServer("hello")
	go PrintServer("sw")
	PrintServer("world")

}
