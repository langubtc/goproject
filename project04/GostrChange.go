package main

import (
	"fmt"
)

func Conver(s string) string {

	//先转成rune数组
	buf := []rune(s)

	//得到每个字符的ascii编码对应的数字，然后修改数组值后再返回
	for k, v := range buf {
		//if v >= 97 && v<= 122 {
		if buf[k] >= 'a' && buf[k] <= 'z' {
			upperValue := v - 32
			//fmt.Println(string(v),reflect.TypeOf(v),string(upperValue))
			//fmt.Println(k)
			buf[k] = upperValue
		}

	}

	return string(buf)

}

func main() {

	//字符串修改将第一个字符修改成T
	str1 := "hello中国"
	//把字符串先转成一个rune数组，对数组进行修改后再赋值
	// 由于rune可表示的范围更大，所以能处理一切字符，当然也包括中文字符。在平时计算中文字符，可用rune。
	a := []rune(str1)
	ab := []byte(str1)

	str1 = string(a)

	for k, v := range a {
		fmt.Println(k, string(v), v)
	}

	fmt.Println(str1, string(ab))

	fmt.Println(Conver(str1))

}
