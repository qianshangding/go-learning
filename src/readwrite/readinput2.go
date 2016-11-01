package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入文字: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("您输入的是: %s\n", input)
	}
}
