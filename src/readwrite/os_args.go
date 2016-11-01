package main

import (
	"fmt"
	"os"
	"strings"
)

//执行go build生成os_args命令后,在控制台执行: ./os_args fish
func main() {
	who := ""
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
