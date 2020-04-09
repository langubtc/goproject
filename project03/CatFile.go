package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PrintFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}

func main() {

	filename := os.Args[1]
	PrintFile(filename)

}
