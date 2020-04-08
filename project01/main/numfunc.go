//表示该文件hello.go 所在的包是main，go中每个文件都必须归属一个包
package main

//引入fmt包 格式化输出
import "fmt"


// 函数使用 func 声明,参数必须声明 类型
// 返回值也需要指定返回类型

func add(n int) int {
	 n = n + 1
	 return n
}


func main(){
	result := add(5)
	fmt.Printf("%s:%d","执行结果",result)
}

