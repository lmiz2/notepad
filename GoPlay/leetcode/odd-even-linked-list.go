package main

import "fmt"

//   Definition for singly-linked list.
//   type ListNode struct {
//       Val int
//       Next *ListNode
//   }
// func main() {
// 	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
// 	// head.Next.Next.Next.Next.Next = head

// 	printList(oddEvenList(head), 10)
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	var t *ListNode = head
	var to *ListNode
	var toTail *ListNode
	var te *ListNode
	var teTail *ListNode
	cnt := 0
	for {
		if t == nil {
			break
		}
		if cnt%2 == 0 {
			if to == nil {
				to = t
				toTail = t
			} else {
				toTail.Next = t
				if toTail.Next != nil {
					toTail = toTail.Next
				}
			}
		} else {
			if te == nil {
				te = t
				teTail = t
			} else {
				teTail.Next = t
				if teTail.Next != nil {
					teTail = teTail.Next
				}
			}
		}
		if t.Next == nil {
			break
		}
		cnt++
		d := t
		t = t.Next
		d.Next = nil
	}
	if toTail != nil {
		toTail.Next = te
	}
	return to
}

func printList(list *ListNode, size int) {
	if list == nil {
		fmt.Println("nil!!")
		return
	}
	t := list
	cnt := 0
	fmt.Print(*t)
	fmt.Print(" -> ")
	for t.Next != list && cnt < size && t.Next != nil {
		t = t.Next
		fmt.Print(*t)
		fmt.Print(" -> ")
		cnt++
	}
	fmt.Println()
}
