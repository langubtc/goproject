package main

import "fmt"

//存放单个字母就用byte
//字符串是一串固定长度的字符连接起来，golang的字符中是由字节组成

func main()  {

	var a byte = 'a'
	var b byte = '0'

	//当直接输入byte值时，就是输出所对应的ascii码
	// a = 97 b = 48
	fmt.Println(a,b)

	//可以使用Printf  %c :以字符形式输出
	fmt.Printf("%c,%c\n",a,b)

	//golang 语言使用UTF8编码


	//字符类型可以进行运算，是按字符所对应的ASCII码来进行运算

	var n1   = 10 + 'a'  //=107 使用单引号,如果是字符串需要使用双引号
	var n2   = "10"+"a" //=10a
	fmt.Println(n1,n2)

	//反引号打印出多行输出
	var string2= `	var n1   = 10 + 'a'  //=107 使用单引号,如果是字符串需要使用双引号
	var n2   = "10"+"a" //=10a
	fmt.Println(n1,n2)`
	fmt.Println(string2)


	//字符串拼接

	var sm1 = "hello " + "news"
	fmt.Println(sm1)
}

