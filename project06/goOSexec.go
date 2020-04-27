package main

import (
	"fmt"
	"log"
	"os/exec"
)

//golang 调用系统命令
//name 表示命令，arg 代表参数列表
//combinedoutput 输出stdout stdin stderr

func main() {
	cmd := exec.Command("ls", "-l")
	//cmd := exec.Command("dir")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

}
