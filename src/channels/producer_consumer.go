package main

import "fmt"

func producer(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- (i * 10)
	}
	close(c)
}

func consumer(c chan int, done chan bool) {
	for num := range c {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func main() {
	c := make(chan int)
	done := make(chan bool)
	go producer(c)
	go consumer(c, done)
	<-done
}
