package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write() {
	//os.O_TRUNC 会清空原来的文件
	//os.O_APPEND  在原来文件中追加
	fileObj, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	str := "成都"
	fileObj.Write([]byte(str))
	fileObj.WriteString("\n")
	fileObj.WriteString("哈哈哈")
	fileObj.WriteString("\n")

	defer fileObj.Close()
}

func writeByBufio() {
	fileObj, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("xxxxxx") //将内容写入到缓冲区
	writer.Flush()               //将缓存区内容写入磁盘

}

func writeByIoutil() {
	str := "测试123554434"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {

		fmt.Printf("write file failed,err:%v\n", err)
		return
	}
}

func main() {

	//write()
	//writeByBufio()
	writeByIoutil()

}
