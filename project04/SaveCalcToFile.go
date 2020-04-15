package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	r, err := os.OpenFile("calc.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {

			//fmt.Printf("%d * %d = %-2d  ",i,j,i*j)

			//v_first_string := strconv.Itoa(i)
			//v_two_string := strconv.Itoa(j)
			//v_result := strconv.Itoa(i*j)

			//result := v_first_string + " * " + v_two_string +" = " + v_result + " "

			//使用Fpringf将输出定向到文件
			fmt.Fprintf(r, "%d * %d = %-2d  ", i, j, i*j)

		}
		r.WriteString("\n")
	}

	f, _ := os.Open("calc.txt")

	//将文件名交给带有缓存的IO Reader
	v := bufio.NewReader(f)

	for {
		line, err := v.ReadString('\n') //指定分割符
		//判断是否是文件结尾
		if err == io.EOF {
			break
		}
		fmt.Print(line)

	}

	f.Close()

}
