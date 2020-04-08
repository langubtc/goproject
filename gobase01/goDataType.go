package main



import (
	"fmt"
	"reflect"
	"unsafe"
)

//golang数据类型
//不同数据类型所占用大小不一样，所以需要制定不同的数据类型，这样更加优化内存的占用

//1.整型  2.浮点  3.复数  4.byte  5.rune  int32  6.uint  无符号  7.int 有符号整型 8.bool类型
//字符串string，统一为utf-8

//布尔值默认值为false


func other(){
	var i int32
	var j float32
	var b bool
	var d complex64
	var s string
	type vs string
	var name  vs = "vssdsd"

	fmt.Println("init32 默认值为:",i,"数据类型:",reflect.TypeOf(i))
	fmt.Println("float32 默认值为:",j)
	fmt.Println("bool 默认值为:",b)
	fmt.Println("complex64 默认值为:",d)
	fmt.Println("string 默认值为:",s,"默认值应该为空")
	fmt.Println("name别名值为:",name,"数据类型:",reflect.TypeOf(name))

}



func main()  {
	other()
	var i uint32 = 1

	//uinit32是4个字节，init8是一个字节

	//unsafe所占内存大小

	//如果不加 ，是根据操作来变化大小，32位就是4 64位的操作系统就是8
	var b int
	fmt.Println(unsafe.Sizeof(b))

	fmt.Println(i,unsafe.Sizeof(i))


	//bool

	var vCheck bool
	fmt.Println(vCheck)

	//rune 就是init32

	var c rune = 23

	fmt.Println(c,unsafe.Sizeof(c))

}
