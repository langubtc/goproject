package main

//我们导入包的路径要从$GOPATH/src后面继续写

import (
	"fmt"
	ca "goproject/project08/calc"   //包的别名如果名太长或者有冲突可取一个别名
	"goproject/project08/package01" //匿名导入使用_就可以,不使用包的方法，主要使用init()函数
)

func main() {

	fmt.Println(ca.Add(1, 2))
	package01.Server("10.10.10.1")
}
