package package01

import (
	"fmt"
)

// init函数在包导入的时候自动执行 优先于main函数
// init函数没有参数也没有返回值

// 如果导入是嵌套的，优先执行最后引用的init函数  A-B-C  结果就是C-B-A的init执行流程

func init() {
	fmt.Println("我是初始化函数")
}

// 标识符大写首字母大写表示对外可见。
// Server 传递一个ip得到一个输出
func Server(ip string) {
	fmt.Println("我的server", ip)
}
