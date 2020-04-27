package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for c := 1; c <= i; c++ {
			fmt.Printf("%d * %d = %-2d ", i, c, i*c)
		}
		fmt.Println()
	}
}
