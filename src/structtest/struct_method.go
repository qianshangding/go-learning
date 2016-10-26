package main

import (
	"fmt"
	"pack"
)

type TwoInts struct {
	a int
	b int
}

type IntVector []int

type employee struct {
	salary int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	fmt.Println(IntVector{1, 2, 3}.Sum())

	salary := &employee{1000}
	fmt.Println(salary.giveRaise(100))

	p := new(pack.Person)
	p.SetFirstName("zhengyu")
	fmt.Println(p)
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func (emp *employee) giveRaise(rate int) (toSalary int) {
	toSalary = emp.salary + rate
	return toSalary
}
