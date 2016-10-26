package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time //匿名字段
}

func (t myTime) yearMonthDay() string {
	return t.Time.String()[0:10]
}

func main() {
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("当前时间:", m.String())
	// 调用myTime.yearMonthDay
	fmt.Println("年月日:", m.yearMonthDay())
}
