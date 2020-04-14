package main

import (
	"flag"
	"fmt"
)

var (
	serverIp   = flag.String("ip", "", "Server Ip address")
	serverPort = flag.Int("port", 0, "Default Port")
)

func main() {

	flag.Parse()
	//fmt.Println(*
	//*serverIp = "127.0.0.1"

	if flag.NFlag() <= 1 { //NFlag返回解析时进行了设置的flag的数量。
		flag.Usage() //当参数小于1个打印flag的默认提示

	} else {
		fmt.Println(*serverIp, *serverPort)
	}

}
