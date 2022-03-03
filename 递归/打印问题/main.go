package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Printf("n=%d", n)

}
func main() {
	test(10)
}
