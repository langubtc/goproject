package main

//HTTP是运行在TCP之上的

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err) //Fatal打印日志并且调用os.Exit 方法
		}
		go handle(conn)

	}

}

// 进程、线程、协程

// 线程是为了更好的利用CPU多核
// 按重量区分 进程>线程>协程
// 出现历史  进程、协程、线程

//C10K 问题解决可采用IO多路复用  select epoll  poll方式
//1.select方式：使用fd_set结构体告诉内核同时监控那些文件句柄，使用逐个排查方式去检查是否有文件句柄就绪或者超时。该方式有以下缺点：文件句柄数量是有上线的，逐个检查吞吐量低，每次调用都要重复初始化fd_set。
//2.poll方式：该方式主要解决了select方式的2个缺点，文件句柄上限问题(链表方式存储)以及重复初始化问题(不同字段标注关注事件和发生事件)，但是逐个去检查文件句柄是否就绪的问题仍然没有解决。
//3.epoll方式：该方式可以说是C10K问题的killer，他不去轮询监听所有文件句柄是否已经就绪。epoll只对发生变化的文件句柄感兴趣。其工作机制是，使用"事件"的就绪通知方式，通过epoll_ctl注册文件描述符fd，一旦该fd就绪，内核就会采用类似callback的回调机制来激活该fd, epoll_wait便可以收到通知, 并通知应用程序。而且epoll使用一个文件描述符管理多个描述符,将用户进程的文件描述符的事件存放到内核的一个事件表中, 这样数据只需要从内核缓存空间拷贝一次到用户进程地址空间。而且epoll是通过内核与用户空间共享内存方式来实现事件就绪消息传递的，其效率非常高。但是epoll是依赖系统的(Linux)。
