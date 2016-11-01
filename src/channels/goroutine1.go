package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始运行main()")
	// longWait()
	// shortWait()
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	//如果我们不在 main() 中等待，协程会随着程序的结束而消亡
	time.Sleep(4 * 1e9)
	fmt.Println("main()运行完成")
}

func longWait() {
	fmt.Println("longWait()开始")
	time.Sleep(3 * 1e9) // sleep for 5 seconds
	fmt.Println("longWait()结束")
}

func shortWait() {
	fmt.Println("shortWait()开始")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("shortWait()结束")
}
