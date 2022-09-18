package main

import "fmt"

func BubbleSortAsc(arr []float32) []float32 {
	lenArr := len(arr)
	for i := 0; i < lenArr-1; i++ {
		for j := 0; j < lenArr-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func BubbleSortDesc(arr []float32) []float32 {
	lenArr := len(arr)
	for i := 0; i < lenArr-1; i++ {
		for j := 0; j < lenArr-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	arr := []float32{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	fmt.Println(arr)
	fmt.Println("Array secara Ascending", BubbleSortAsc(arr))
	fmt.Println("Array secara Descending" ,BubbleSortDesc(arr))
}