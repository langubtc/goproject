package main

import (
	"fmt"
	"reflect"
)

type interfaceName interface {
	// 方法列表
	GetName() string
	GetAge() int
}

type Name struct {
	name string
	age  int
}

func (s Name) GetName() string {
	return s.name
}

func (s Name) GetAge() int {
	return s.age
}

type DefaultI interface {
	GetName()
}

func main() {

	s := Name{"roddy", 28}
	a := s.GetName()
	fmt.Println(a, s.GetAge(), s.age)

	var allData interface{}
	var allList interface{}

	allData = 234

	fmt.Println(allData, reflect.TypeOf(allData))

	m := 123
	ms := "check all"

	allList = m
	allList = ms //可覆盖此前的变量值
	fmt.Println(allList)

	var mk DefaultI
	fmt.Println(mk)

}
