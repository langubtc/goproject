package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := student{
		"roddy",
		90,
	}

	//通过反射获取结构体中的所有字段信息

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	//遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		fileObj := t.Field(i)
		//fmt.Println(fileObj.Name,fileObj.PkgPath,fileObj.Tag,fileObj.Type)

		fmt.Println(fileObj.Tag.Get("json"), fileObj.Tag.Get("ini"))

	}

	//根据名字去获取结构体字段

	filedObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Println(filedObj2.Tag.Get("json"))

	}

	var serverName = "web服务器"

	fmt.Println(serverName)

	reflectType := reflect.TypeOf(serverName)
	reflectValue := reflect.ValueOf(serverName)
	fmt.Println(reflectType, reflectValue)

}
