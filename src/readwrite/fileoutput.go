package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 可以看到，OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符“|”连接），使用的文件权限。
	// 我们通常会用到以下标志：
	// os.O_RDONLY：只读
	// os.O_WRONLY：只写
	// os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	// os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	// 如果写入的东西很简单，我们可以使用 fmt.Fprintf(outputFile, “Some test data.\n”) 直接将内容写入文件。fmt 包里的 F 开头的 Print 函数可以直接写入任何 io.Writer，
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
