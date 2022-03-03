package main

import "fmt"

func main() {
	arr := []int{10, 19, 29, 12, 56, 34}
	InsertSort(arr)
}
func InsertSort(arr []int) {
	for j := 0; j < len(arr); j++ {
		insertVal := arr[j]
		insertIndex := j - 1
		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		arr[insertIndex+1] = insertVal
	}
	fmt.Println(arr)
}
