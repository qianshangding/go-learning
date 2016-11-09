package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(1 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	//变量c为非缓冲channel，会堵塞
	c <- 10
	fmt.Println("sent", 10)
	time.Sleep(1 * 1e9)
}
