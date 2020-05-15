package main

import "fmt"

type SERVER struct {
	ServerName string
	CPU        int
	MEM        int
	DISK       int
	IP         string
}

func changeServer(server *SERVER) {
	if server.MEM > 4 {
		server.MEM = 8
	}
}

func GetServerInfo(server SERVER) {
	fmt.Printf("server ip:%s\n", server.IP)

}

func main() {
	//{HOSTNAME 4 8 500 200.2.14.2} 创建一个结构体，并赋予值
	fmt.Println(SERVER{"DBSERVER", 4, 8, 500, "200.2.14.2"})

	//也可以使用KEY->VALUE格式定义一个结构体
	serverInfo := SERVER{ServerName: "web server", CPU: 4, MEM: 8, DISK: 200, IP: "47.23.23.1"}

	fmt.Println(serverInfo.IP)

	//如果没有赋值，将使用缺省默认值
	webInfo := SERVER{ServerName: "web 2server", DISK: 200, IP: "47.23.23.2"}

	fmt.Println(webInfo)

	var newServer SERVER

	//赋值
	newServer.ServerName = "new server"
	newServer.DISK = 500
	newServer.MEM = 6
	newServer.CPU = 2
	newServer.IP = "2.2.3.2"

	GetServerInfo(newServer)
	changeServer(&newServer)

	//读取

	fmt.Printf("服务器名%s,磁盘:%d,内存:%d,CPU:%d,IP:%s", newServer.ServerName,
		newServer.DISK, newServer.MEM,
		newServer.CPU, newServer.IP,
	)

}
