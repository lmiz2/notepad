package main

import (
	"sort"
)

// type item struct {
// 	num    int
// 	toCost int
// 	// index  int
// 	// bitmask int
// 	// sum   int
// 	// stack string
// }

// type priorityQueue []*item

// func (pq priorityQueue) Len() int { return len(pq) }
// func (pq priorityQueue) Less(i, j int) bool {
// 	return pq[i].toCost < pq[j].toCost
// }
// func (pq priorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	// pq[i].index = i
// 	// pq[j].index = j
// }
// func (pq *priorityQueue) Push(x interface{}) {
// 	// n := len(*pq)
// 	item := x.(*item)
// 	// item.index = n
// 	*pq = append(*pq, item)
// }
// func (pq *priorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil
// 	// item.index = -1
// 	*pq = old[0 : n-1]
// 	return item
// }
type line struct {
	by    int
	to    int
	value int
}

func solution(n int, costs [][]int) int {
	// maxBit := (1 << n) - 1
	costMap := make([][]int, n)
	nodes := make([]int, n)
	lineArr := []line{}
	for i := range costMap {
		costMap[i] = make([]int, n)
	}
	for i := range nodes {
		nodes[i] = i
	}
	for _, arr := range costs {
		costMap[arr[0]][arr[1]] = arr[2]
		costMap[arr[1]][arr[0]] = arr[2]
		lineArr = append(lineArr, line{arr[0], arr[1], arr[2]})
	}
	sort.Slice(lineArr, func(i, j int) bool {
		return lineArr[i].value < lineArr[j].value
	})

	return -1
}
