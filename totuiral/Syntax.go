package main

import (
	"errors"
	"fmt"
	"math"
)

const Pi = 3.14

func add(a int, b int) int {
	return a + b
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다.")
	}
	return a / b, nil
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var a int = 10
	var b float64 = 20.5
	var c string = "Hello"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(Pi)

	d := 10
	if d > 5 {
		fmt.Println("a는 5보다 큽니다.")
	} else {
		fmt.Println("a는 5보다 작거나 같습니다.")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	result := add(5, 3)
	fmt.Println(result)

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("오류:", err)
	} else {
		fmt.Println("결과:", result)
	}

	fmt.Println(math.Sqrt(16))
}
