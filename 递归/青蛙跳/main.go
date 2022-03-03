package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		temp := b
		a, b = b, a+b
		return temp
	}
}

func main() {
	s := make([]int, 0)
	f := fib()
	for i := 0; i < 20; i++ {
		s = append(s, f())
	}
	fmt.Println(s)
}
