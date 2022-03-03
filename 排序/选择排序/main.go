package main

import "fmt"

func main() {
	arr := []int{10, 34, 109, 80, 100}
	// fmt.Println(arr)
	SelectSort(arr)
}

func SelectSort(arr []int) {
	for j := 0; j < len(arr); j++ {
		max := arr[j]
		maxIndex :=j
		for i := 0; i < len(arr); i++ {
			if max > arr[i] {
				max = arr[i]
				maxIndex = i
			}
			if maxIndex+1 != j {
				arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
