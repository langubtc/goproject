package main

import "fmt"

type IDC struct {
	name    string
	address string
}

type CloudServer struct {
	name string
	ip   string
	cpu  int8
	disk int32
	mem  int8
	net  float32
	IDC  //匿名嵌套另外的结构体
}

func main() {
	s1 := CloudServer{
		"webserver",
		"10.10.1.1",
		4,
		500,
		8,
		8,
		IDC{
			"北京1区",
			"北京xxxxx路",
		},
	}

	//fmt.Println(s1.idc.name,s1)
	fmt.Println(s1.address)

}
