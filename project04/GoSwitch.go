package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Args Error")
		os.Exit(1)
	}

	selectType := os.Args[1]

	switch selectType {
	case "hostname", "host":
		fmt.Println("HOST")
	case "disk":
		fmt.Println("DISK")
	case "mem":
		fmt.Println("MEM")
	case "net":
		fmt.Println("NET")
	default:
		fmt.Println("nothing")

	}

}
