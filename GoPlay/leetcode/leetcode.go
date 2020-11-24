package main

import "fmt"

func main() {
	// arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	arr := []int{0, 1, 0, 3, 2, 3}
	// arr := []int{-77, 0, 77, -1, 3, -56, 123, 22, 5481, 124, 56}
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	// arr := []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}
	// arr := []int{1, 2, 3, 4}
	fmt.Println(lengthOfLIS(arr))
}
