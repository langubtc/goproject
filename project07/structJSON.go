package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name string
	Sn   string
}

//book的构造函数

func NewBook(name, sn string) Book {
	return Book{
		Name: name,
		Sn:   sn,
	}
}

func main() {

	book_list := make([]Book, 0, 20)
	for i := 0; i < 10; i++ {

		b1 := NewBook(fmt.Sprintf("django%v", i), fmt.Sprintf("232344422%v", i))
		book_list = append(book_list, b1)
	}

	//json 序列化数据

	data, err := json.Marshal(book_list)

	if err != nil {
		fmt.Printf("json 格式化错误,err:%v\n", err)
		return
	}

	fmt.Println(string(data))

	//json反序列化

	Jsostr := `{"Name":"百世go","Sn":"SN2323232323232323"}`
	var c2 Book
	err = json.Unmarshal([]byte(Jsostr), &c2)
	if err != nil {
		fmt.Printf("json 解析失败 err:%v\n", err)
		return
	}

	fmt.Println(c2)

}
