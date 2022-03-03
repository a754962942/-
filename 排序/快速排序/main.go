package main

import "fmt"

func QuickSort(arr []int) []int {
	//必须判断arr.len
	if len(arr) <= 1 {
		return arr
	}
	temp := arr[0]
	low := make([]int, 0)
	mid := make([]int, 0)
	high := make([]int, 0)
	mid = append(mid, temp)
	for i := 1; i < len(arr); i++ {
		if arr[i] < temp {
			low = append(low, arr[i])
		} else if arr[i] > temp {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	myarry := append(append(low, mid...), high...)
	return myarry
}
func main() {
	arr := []int{5, 2, 56, 29, 89, 9, 45, 34, 32, 3, 8}
	fmt.Println(QuickSort(arr))

}
