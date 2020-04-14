package main

import "fmt"

//解题思路是先将string转成两个rune类型，然后循环遍历将newStrings每个字符进行修改

func main() {
	//反转字符串
	var string1 = "HELLO"

	buf := []rune(string1)
	newStrings := []rune(string1)

	for i := 0; i < len(buf); i++ {
		if i == 0 {
			newStrings[i] = buf[len(buf)-1]
		} else {
			newStrings[i] = buf[len(buf)-i-1]
			fmt.Println(len(buf) - i - 1)
		}

	}
	//输出OLLEH
	fmt.Println(string(newStrings))

}
