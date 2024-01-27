package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
// 	head.Next.Next.Next.Next.Next.Next.Next = head

// 	printList(oddEvenList(head), 14)
// }

func main() {
	// arr := []int{1, 5, 1, 1, 6, 4}
	// arr := []int{1, 5, 5, 3, 2, 5, 5, 4, 5, 5}
	arr := []int{12, 41345, 12319, 135, 32423, 3513468, 9765798, 435, 1245342, 123, 12, 4}
	wiggleSort(arr)
	fmt.Println(arr)
}

func wiggleSort(nums []int) {
	tNums := make([]int, len(nums))
	copy(tNums, nums)
	sort.Ints(tNums)
	i := len(tNums) - 1
	for j := 1; j < len(nums); j += 2 {
		nums[j] = tNums[i]
		i--
	}
	for j := 0; j < len(nums); j += 2 {
		nums[j] = tNums[i]
		i--
	}
}
