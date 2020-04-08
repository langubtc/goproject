package main

import "fmt"

func site(){
	a := byte(0)
	b := byte(1)

	fmt.Println(a&b)
	fmt.Println(a>>b)

}

func start(){
	a :=true
	b :=true

	fmt.Println(a&&b)
	fmt.Println(a||b)

}



func  main()  {
	site()
	start()
	a:=1
	b:=2
	fmt.Println("A>b",a>b)
	fmt.Println("A<b",a<b)
	fmt.Println("A==b",a==b)
	fmt.Println("A!=b",a!=b)
	fmt.Println("A>=b",a>=b)
	// ++ 和--只能有这种用法
	a++
	b--
	c := a
	fmt.Println(a+b)
	fmt.Println(a*b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}