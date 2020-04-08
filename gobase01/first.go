package main

//导入所依赖的包
import "fmt"



//25个关键字 36个标识符
//1.单行注释和多行注释
//这里是单行注释
/**
	这里是行1
	这里是行2
**/

// const 来定义一个全局的常量

const NAME = "roddy"

const AGE int  = 28

//一般类型声明
type imoocInt int


//结构体声明

type Learn struct {

}

//声明接口
type ILearn interface {

}

//声明函数

func  star()  {

	var YouAge imoocInt = 0
	fmt.Println("开始启动......")
	fmt.Println("Yage......",YouAge +1)
}


func main(){
	//fmt.Printf("golang <<< ")
	star()
	fmt.Printf("My name:%s Age:%d\n",NAME,AGE)
	fmt.Print("My name:",NAME)
}


