package main

import "fmt"

func MinMax(input ...int) (min, max int) {
	min = input[0]
	max = input[0]

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}

		if input[i] > max {
			max = input[i]
		}
	}
	return
}

func main() {
	min, max := MinMax(1, 2, 34, 5, 6, 7)
	fmt.Println("最小值为:", min, "最大值为:", max)
}
