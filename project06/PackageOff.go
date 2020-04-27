package main

import "fmt"

func addUpper() func(int) int {
	n := 10
	n -= 1
	c := 20

	return func(x int) int {
		c -= 2
		fmt.Println(c)

		n = n + x
		return n
	}
}

func main() {

	t := addUpper()

	fmt.Println(t(5))

}
