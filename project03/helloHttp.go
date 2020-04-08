package main

import (
	"fmt"
	"net/http"
)

//定义handler方法  w 是writer, r 就是request
func Handler(w http.ResponseWriter,r *http.Request) {
	//格式化并且输出`Fprintf` 来格式化并输出到 `io.Writers` 并不是os.Stdout
	fmt.Fprintf(w,"Hello,world! %s",r.URL.Path)
}

func UserInfo(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"{'username':'roddy'}")
}


func main(){

	http.HandleFunc("/",Handler)          //注册到handler方法
	http.HandleFunc("/userInfo",UserInfo) //注册到UserInfo handler方法
	_ = http.ListenAndServe(":8080", nil) //监听到8080端口
}


