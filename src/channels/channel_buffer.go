package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	//变量c为缓冲channel，在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的
	c <- 10
	fmt.Println("sent", 10)
}
