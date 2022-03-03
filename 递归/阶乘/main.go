package main

import "fmt"

func f(i int) int {

	if i == 1 {
		return 1
	}
	return i * f(i-1)
}
func main() {
	res := f(3)
	fmt.Println(res)
}
