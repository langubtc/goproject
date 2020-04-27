package main

import "fmt"

func MinMaxCount(args ...int) (min int, max int) {
	min = args[0]
	max = args[0]

	for i := 0; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}

		if args[i] > max {
			max = args[i]
		}
	}
	return
}

func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {

	min, max := MinMaxCount(1, 2, 3, 4, 2, 5, 2, 7, 8)
	fmt.Println(min, max)

	fmt.Println(Factorial(2))

}
