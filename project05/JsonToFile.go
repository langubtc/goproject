package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Id   int `json:"student_id"`
	Name string
}

//将json保存到文件，并且将文件内容读取出来

func main() {
	c := Student{Id: 1, Name: "roddy"}
	buf, _ := json.Marshal(c)

	f, _ := os.OpenFile("go.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	f.WriteString(string(buf))
	f.WriteString("\n")
	f.Close()

	filename, _ := os.Open("go.json")

	//将文件名交给带有缓存的IO Reader
	r := bufio.NewReader(filename)

	for {
		line, err := r.ReadString('\n') //指定分割符
		//判断是否是文件结尾
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	filename.Close()

}
