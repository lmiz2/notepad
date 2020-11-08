package main

import "fmt"

func main() {

	// arr := [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}}
	// fmt.Println(solution(4, arr))

	// arr := {}{}int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}, {1, 4, 6}, {2, 5, 1}, {3, 5, 2}}
	// fmt.Printf("\n answer : %d", solution(6, arr))

	arr := [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 3}, {3, 1, 2}, {3, 0, 4}, {2, 4, 6}, {4, 0, 7}}
	fmt.Printf("\n answer : %d", solution(5, arr))
}
