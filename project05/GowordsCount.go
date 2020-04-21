package main

import (
	"fmt"
	"strings"
)

func main() {
	input_words := "one two   two   three two one  CCC KKK"
	map_string := make(map[string]int)
	//将字符串转成slice 并且以空格做区分
	string_slice := strings.Fields(input_words)

	for _, v := range string_slice {
		_, ok := map_string[v]
		if ok {
			map_string[v] += 1
		} else {
			map_string[v] = 1
		}
	}

	for k, v := range map_string {

		fmt.Printf("%v-->%d\n", k, v)
	}

}
