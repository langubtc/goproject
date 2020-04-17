package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	//data := []byte(buf)
	md5sum := md5.Sum(buf)
	//十六进制表示，字母形式为小写
	fmt.Printf("md5:%x,len:%d\n", md5sum, len(md5sum))

}
