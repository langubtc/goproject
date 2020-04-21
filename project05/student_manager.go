package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

var students = make(map[string]Student)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Sscan(line, &cmd)

		switch cmd {
		case "add":
			var id int
			var name string
			var age int

			fmt.Sscan(line, &cmd, &id, &name, &age)
			students[name] = Student{Id: id, Name: name, Age: age}
			fmt.Println(students)

		case "list":
			var name string
			fmt.Sscan(line, &cmd, &name)
			values, ok := students[name]
			//如果存在记录就解析
			if ok {
				fmt.Printf("ID\tNAME\tAGE\n")

				fmt.Println(values.Id, values.Name, values.Age)

			} else {
				fmt.Printf("ID\tNAME\tAGE\n")
				for _, v := range students {
					fmt.Println(v.Id, v.Name, v.Age)
				}
			}

		case "save":
			f, _ := os.OpenFile("gostudents.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
			for _, v := range students {
				buf, _ := json.Marshal(v)
				fmt.Println(string(buf))
				f.WriteString(string(buf))
				f.WriteString("\n")

			}
			f.Close()

		case "load":
			var s3 Student
			filename, _ := os.Open("gostudents.json")

			//将文件名交给带有缓存的IO Reader
			r := bufio.NewReader(filename)

			for {
				line, err := r.ReadString('\n') //指定分割符
				//判断是否是文件结尾
				if err == io.EOF {
					break
				}
				json.Unmarshal([]byte(line), &s3)
				students[s3.Name] = s3
			}
			filename.Close()
		case "exit":
			os.Exit(0)
		}
	}

}
