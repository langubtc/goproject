package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//defer 用于程序释放资源,用于延迟操作。

func closeFile() {

	f, _ := os.Open("Godefer.go")

	//将文件名交给带有缓存的IO Reader
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n') //指定分割符
		//判断是否是文件结尾
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	defer f.Close()

}

func deferFuncReturn() (result int) {
	i := 1
	//延迟调用result,修改result值
	defer func() {
		result++
	}()

	return i
}

func f() {
	//延迟调用，采用FILO(先进后出原则)
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return

}

func main() {
	f() //结果是1 后是0
	fmt.Println(deferFuncReturn())
	closeFile()
}
