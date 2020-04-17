package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)

}

func main() {
	s := [3]int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))
	s1 := s[:2]
	s2 := s[:2]

	a := 1
	c := 1
	fmt.Println(a == c)

	fmt.Println(s1, s2)

	//切片两者不能作比较
	//fmt.Println(s1 == s2) //slice can only be compared to

	//make 创建slice
	sss := make([]int, 5)
	printSlice("a", sss)

	//空数组才能进行append,如果要用append,必须使用make
	var vbsc []int

	vbsc = append(vbsc, 0)

	fmt.Println(vbsc)

	//切片原理，如果超过容量会申请2倍的空间
	mmm := make([]int, 1, 4)
	for i := 0; i < 5; i++ {
		mmm = append(mmm, i)
	}

	fmt.Println(mmm)

	mlist := []int{1, 2, 3, 5}

	for v := range mlist {
		fmt.Println(v)
	}

}
