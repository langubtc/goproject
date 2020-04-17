package main

import "fmt"

func main() {
	string1 := []string{"one", "two", "three"}

	string2 := [...]string{0: string1[0]}

	fmt.Println(string1[:], string2)

	array_list := [...]byte{'s', 2, 'v', 'e'}
	fmt.Println(array_list)

	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]

	//cap统计容量
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_a), cap(Slice_a), Slice_a)

	Slice_b := Slice_a[0:8]
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_b), cap(Slice_b), Slice_b)

}
