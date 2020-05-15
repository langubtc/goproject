package main

//golang 打开文件\读取文件
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readAll() {

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
			fmt.Println(string(tmp[:n]))
			return
		}

		if err != nil {
			fmt.Printf("read  from file failed, err:%v\n", err)
		}

		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp))
	}

}

func readByBufio() {
	fileObj, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)

	for {

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Printf(" read file by bufio failed,err:%v\n", err)
		}

		fmt.Print(line)
	}

}

func readByIoutil() {

	//ioutil不能读取大文件,否则会占用内存
	content, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed,err:%v\n", err)
		return
	}

	fmt.Println(string(content))

}

func main() {

	//readAll()
	//readByBufio()
	readByIoutil()
}
