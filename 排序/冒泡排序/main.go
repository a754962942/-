package main

import "fmt"

func main() {
	arr := []int{2, 1, 3, 5, 4, 9, 6, 7, 8, 0}
	fmt.Println(arr)
	BubbleSort(arr)
}
func BubbleSort(arr []int) {
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
