package main

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left := arr[0:mid]
	right := arr[mid:]
	return Merge(MergeSort(left), MergeSort(right))

}

func Merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	} else {
		result = append(result, right...)
	}
	return result
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := MergeSort(arr)
	fmt.Println(result)
}
