package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("类型正确，类型是:: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("类型正确，类型是: %T\n", u)
	} else {
		fmt.Println("不是Circle类型")
	}

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Square类型： %T -> %v\n", t, t)
	case *Circle:
		fmt.Printf("Circle类型： %T -> %v\n", t, t)
	case nil:
		fmt.Printf("无效类型\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	//验证是什么类型的
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("参数 #%d 是 一个 bool\n", i)
		case float64:
			fmt.Printf("参数 #%d 是 一个 float64\n", i)
		case int, int64:
			fmt.Printf("参数 #%d 是 一个 int\n", i)
		case nil:
			fmt.Printf("参数 #%d 是 一个 nil\n", i)
		case string:
			fmt.Printf("参数 #%d 是 一个 string\n", i)
		default:
			fmt.Printf("参数 #%d 是 一个无效类型\n", i)
		}
	}
}
