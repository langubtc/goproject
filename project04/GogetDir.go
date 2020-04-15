package main

import (
	"fmt"
	"os"
)

func getPath(f string) {
	fmt.Println(f)
	filepath, _ := os.Open(f)
	infos, _ := filepath.Readdir(-1)
	for _, info := range infos {

		if info.IsDir() {
			fmt.Println("di")
			getPath("../" + info.Name())

		}
		fmt.Printf("%v %v\n", info.Name(), info.Size())
	}

	filepath.Close()

}

func main() {
	//使用os.Open 读取目录下的文件
	//filepath,_:=os.Open(os.Args[1])
	filename := os.Args[1]

	getPath(filename)
	//infos,_ := filepath.Readdir(-1)
	//for _,info := range infos {
	//
	//	fmt.Printf("%v %v\n",info.Name(),info.Size())
	//}
	//
	//filepath.Close()

}
