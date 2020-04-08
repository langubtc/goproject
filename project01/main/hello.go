//表示该文件hello.go 所在的包是main，go中每个文件都必须归属一个包
package main

//引入fmt包 格式化输出
import "fmt"


//打印hello world  func是关键字，表示后面是一个函数
//main是我们程序的入口

func main(){
	//使用fmt包中的Println函数 输入Hello World
	fmt.Println("Hello World")

}

//go build hello.go   生成hello.exe文件
//go run hello.go     类似执行脚本方式

//流程应该是先编译再执行，通常不建议使用go run 方式执行代码
