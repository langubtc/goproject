package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int    `json:"student_id"`
	Name string `json:"name,omitempty"` //omitempty如果没有序列化就不解析

}

func main() {

	var s Student

	s.Id = 1
	s.Name = "roddy"

	var b Student
	b.Id = 2

	buf, _ := json.Marshal(s)

	/**
		MarshalIndent 格式化输出
		{
			"student_id": 1,
			"name": "roddy"
		}

	**/

	buf_2, _ := json.MarshalIndent(b, "", "\t")
	fmt.Println(string(buf))
	fmt.Println(string(buf_2))

}
