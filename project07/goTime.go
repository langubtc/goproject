package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	now := time.Now() //时间对象
	//fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

	//
	////时间戳 1970.1.1 到现在的秒数

	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(timeStamp1) //当前世间戳
	fmt.Println(timeStamp2) //当前世间戳微秒，纳秒
	//
	////将时间戳转换为具体的时间格式
	//
	t := time.Unix(1589514825, 0)
	fmt.Println(t)
	fmt.Printf("%v年%v月%v日\n", t.Year(), t.Month(), t.Day())

	//
	////时间间隔 以纳秒为单位
	//
	//time.Sleep(time.Duration(1) * time.Second) // duration是一种类型
	//
	//time.Sleep(time.Second * 1)  //沉睡1秒
	//
	//fmt.Println()
	//
	//
	////获取3小时之后的时间
	//
	//t2 := now.Add(time.Hour * 3)
	//fmt.Println(t2)
	//
	//t3 := t2.Sub(now)
	//fmt.Println(t3)
	//
	//
	////time设置定时器
	//
	//for tmp  := range  time.Tick(time.Second * 3){
	//	fmt.Println(tmp)  //每隔3秒打印
	//
	//}

	//tmp:= time.Tick(time.Second * 3)
	//
	//for{
	//
	//	select {
	//	case <-tmp:
	//		fmt.Println("执行任务")
	//	}
	//}

	//
	//
	////时间格式化
	////不是使用y-m-d
	////    Y    m  d  H  M  S
	////使用2006 01 02 15 04 05  口诀 2006 1 2 15 4 5
	//fmt.Println(now.Format("2006-01-02 15:04:05")) //格式化后就是2020-05-15 13:52:44
	//fmt.Println(now.Format("2006-01-02 15:04:05.000")) //格式化后就是2020-05-15 13:52:44.234  //具体到毫秒
	//fmt.Println(now.Format("2006-01-02 03:04:05 PM")) //格式化后就是2020-05-15 01:56:04 PM
	//
	//
	//
	////解析字符串格式的时间
	//// 1.拿到时区
	//loc,err := time.LoadLocation("Asia/Shanghai")
	//if err != nil{
	//
	//	fmt.Println(err)
	//	return
	//}
	//
	////2.根据时区去解析一个字符串格式的时间
	//
	//timeStr := "2020/05/15 13:58:00"
	//
	//t4,err:=time.ParseInLocation("2006/01/02 15:04:05",timeStr,loc)
	//if err != nil{
	//	fmt.Printf("parse time failed err:%v\n",err)
	//	return
	//
	//}
	//fmt.Println(t4,t4.Year())
	//
	//currentTime := time.Now()
	//oldTime := currentTime.AddDate(0, 0, 5)
	//
	//newTime := currentTime.Add(time.Hour * 5)
	//t8:=newTime.Sub(currentTime)
	//fmt.Println(t8)
	//
	//fmt.Println(oldTime)
	//fmt.Println(newTime)

	// 获得时区

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeStr := "2020-03-25 05:02:00"

	strToTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	fmt.Println(strToTime, reflect.TypeOf(strToTime))

	utcTime := "2020-05-14 02:08:23"

	tm, _ := time.Parse("2006-01-02 15:04:05", utcTime)

	timeStamp3 := tm.Unix()
	tLocal := time.Unix(timeStamp3, 0)
	fmt.Println(tLocal)

}
