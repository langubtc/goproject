package main

import "fmt"

// if else   if else if  else;


func main()  {
	a :=1
	b :=1

	if a >b {
		fmt.Println("a>b")
	}else if a<b{
		fmt.Println("a<b")
	}else{
		fmt.Println(">>end<<")
		if b == 1{
			fmt.Println("B==1")
		}
	}
}
