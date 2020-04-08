package main

import "fmt"

//golang匿名变量特点是一个下划线"_",称为空白标识符,可以用于声明变量或者赋值

func GetUserInfo()(int,string){

	UserName := "roddy"
	UserAge := 27
	return UserAge,UserName
}


func main(){
	_, name := GetUserInfo()
	age, _ := GetUserInfo()

	fmt.Println("My name is:",name,"age:",age)

}
