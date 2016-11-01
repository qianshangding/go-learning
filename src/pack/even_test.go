package pack

import (
	"testing"
)

// 1）func (t *T) Fail()
//     标记测试函数为失败，然后继续执行（剩下的测试）。
// 2）func (t *T) FailNow()
//     标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
// 3）func (t *T) Log(args ...interface{})
//     args 被用默认的格式格式化并打印到错误日志中。
// 4）func (t *T) Fatal(args ...interface{})
//     结合 先执行 3），然后执行 2）的效果。

//go test even_test.go -v  通过该命令查看测试结果
func TestEven(t *testing.T) {
	if Even(10) {
		t.Log("Everything OK: 10 is even, just a test to see failed output!")
		t.Fail()
	} else {
		t.Log("is good")
	}

	if !Even(10) {
		t.Log(" 10 must be even!")
		t.Fail()
	} else {
		t.Log("is good!")
	}

	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}

}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
