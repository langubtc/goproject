package main

import "fmt"

type Server struct {
	ID   int
	Name string
	Type string
	Cpu  int
	Disk int
}

//结构体构造函数
func ServerManager(id int, Name string, Type string, Cpu int, Disk int) *Server {
	return &Server{
		ID:   id,
		Name: Name,
		Type: Type,
		Cpu:  Cpu,
		Disk: Disk,
	}
}

func (r Server) Run() {
	fmt.Printf("%d服务器运行......\n", r.ID)
}

func (r Server) Stop() {
	fmt.Printf("%d服务器停止......\n", r.ID)

}

//修改值传结构体指针
func (r *Server) ChangeName() {

	r.Name = "dbserver"
}

func main() {
	s := ServerManager(1, "webServer", "云服务器", 4, 500)

	s.Run()
	s.Stop()

	s.ChangeName()
	fmt.Println(s.Name)

}
