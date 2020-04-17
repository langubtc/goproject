package main

import "fmt"

//删除数组中的第几个元素

func deleteArray(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	fmt.Println(s, s[i:], s[i+1:])
	return s[:len(s)-1]
}

func deleteString(s []string) []string {
	cvalue := s
	bvalue := s

	fmt.Println(cvalue == bvalue)
	//for k,v :=range(s){

	//}

}

func main() {

	mone := []int{1, 2, 3, 4, 5}
	fmt.Println(deleteArray(mone, 0))

	cstring := []string{"tom", "tom", "tom", "jack", "tom"}

	fmt.Println(deleteString(cstring))

}
