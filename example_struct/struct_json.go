package main

import "fmt"

type StringValue struct {
	Key   int
	Value int
}

func main() {

	//
	a := StringValue{1, 2}
	fmt.Println(a)

}
