package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fileObj, err := os.Open("./xx.txt") //相对目录打开文件

	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	defer fileObj.Close() //关闭文件

	for {

		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)

		if err == io.EOF {
			fmt.Print(string(tmp[:n]))
			return
		}

		if err != nil {
			fmt.Printf("read  from file failed, err:%v\n", err)
		}

		fmt.Print(string(tmp))
	}

}
