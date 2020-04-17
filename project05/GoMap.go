package main

import "fmt"

func main() {
	//定义字典方式1：使用make
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	//方法2
	school := map[string]string{
		"roddy": "xxxx",
		"bbb":   "ccccc",
	}

	fmt.Println(ages)
	for k, v := range ages {

		fmt.Println(k, v)
	}

	for k, v := range school {

		fmt.Println(k, v)
	}
	//获取其中一个key的值
	fmt.Println(school["roddy"])

	//获取时可以判断是否存在
	//value,true 或者false
	c, ok := school["roddssy"]
	if ok {
		fmt.Println(c, ok)
	} else {
		c = "defalut"
		fmt.Println(c)
	}

	//删除其中一个key的值
	delete(school, "roddy")
	fmt.Println(school)

}
