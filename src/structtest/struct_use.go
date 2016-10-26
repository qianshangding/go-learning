package main

import (
	"fmt"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	//各种初始化结构化
	var ms1 *struct1 = new(struct1)
	ms1.i1 = 10
	ms1.f1 = 15.5
	ms1.str = "Chris"

	fmt.Printf("The int is: %d\n", ms1.i1)
	fmt.Printf("The float is: %f\n", ms1.f1)
	fmt.Printf("The string is: %s\n", ms1.str)
	fmt.Println(ms1)

	fmt.Println("m2-----------------------------------------")

	ms2 := new(struct1)
	ms2.i1 = 10
	ms2.f1 = 15.5
	ms2.str = "Chris"

	fmt.Printf("The int is: %d\n", ms2.i1)
	fmt.Printf("The float is: %f\n", ms2.f1)
	fmt.Printf("The string is: %s\n", ms2.str)
	fmt.Println(ms2)

	fmt.Println("m3-----------------------------------------")

	ms3 := &struct1{10, 15.5, "Chris"}

	fmt.Printf("The int is: %d\n", ms3.i1)
	fmt.Printf("The float is: %f\n", ms3.f1)
	fmt.Printf("The string is: %s\n", ms3.str)
	fmt.Println(ms3)

	fmt.Println("m4-----------------------------------------")

	var ms4 struct1
	ms4 = struct1{10, 15.5, "Chris"}

	fmt.Printf("The int is: %d\n", ms4.i1)
	fmt.Printf("The float is: %f\n", ms4.f1)
	fmt.Printf("The string is: %s\n", ms4.str)
	fmt.Println(ms4)

	fmt.Println("m5-----------------------------------------")

	ms5 := &struct1{i1: 10, f1: 15.5, str: "Chris"}

	fmt.Printf("The int is: %d\n", ms5.i1)
	fmt.Printf("The float is: %f\n", ms5.f1)
	fmt.Printf("The string is: %s\n", ms5.str)
	fmt.Println(ms5)

}
