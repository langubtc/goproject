package main

import "fmt"

// Server是一个结构体
type server struct {
	name string
	ip   string
}

// 构造一个Server的函数
func NewServer(name, ip string) *server {

	return &server{
		name: name,
		ip:   ip,
	}
}

//函数不属于任何类型,可以在任何地方调用
//方法属于具体类型

//定义方法
func (s server) Check() {
	fmt.Printf("%s\n", s.ip)
}

func (s server) Ping() {
	fmt.Printf("正在检查ping %s\n", s.ip)
}

//指针接收者
func (s *server) ChangeIp(ip string) {
	fmt.Println("修改前的IP:", s.ip)

	s.ip = ip
	fmt.Println("修改后的IP:", s.ip)
}

//值接收者
func (s server) ChangeValueIp(ip string) {
	fmt.Println("修改前的IP:", s.ip)

	s.ip = ip
	fmt.Println("修改后的IP:", s.ip)
}

func main() {

	server1 := NewServer("news", "192.168.1.2")
	server1.Check()
	//(*server1).Ping()
	server1.Ping()
	server1.ChangeIp("10.10.10.2")
	fmt.Println(server1.ip)

	server1.ChangeValueIp("10.10.10.4")
	fmt.Println(server1.ip)
}
