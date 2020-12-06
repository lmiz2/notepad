package main

import "sort"

// func main() {
// 	// arr := []int{1, 2, 5}
// 	// arr := []int{3}
// 	// arr := []int{186, 419, 83, 408} // 6249
// 	// arr := []int{3, 7, 405, 436} // 8839
// 	// arr := []int{284, 260, 393, 494} // 7066
// 	// arr := []int{429, 171, 485, 26, 381, 31, 290} // 8440
// 	// arr := []int{259, 78, 94, 130, 493, 4, 168, 149} // 4769
// 	arr := []int{288, 160, 10, 249, 40, 77, 314, 429} // 9208

// 	fmt.Println(coinChange(arr, 9208))
// }

// const maxUint = ^uint(0)
// const maxInt = int(maxUint >> 1)

// func coinChange(coins []int, amount int) int {
// 	memory := make([]int, amount+1)
// 	for i := range memory {
// 		memory[i] = -1
// 	}
// 	memory[0] = 0
// 	for i := range memory {
// 		calcVals := []int{}
// 		for _, cv := range coins {
// 			if i >= cv {
// 				calcVals = append(calcVals, memory[i-cv])
// 			}
// 		}
// 		prevAmount := min(calcVals)
// 		if prevAmount >= 0 {
// 			memory[i] = prevAmount + 1
// 		}
// 	}
// 	return memory[amount]
// }
// func min(arr []int) int {
// 	if arr == nil || len(arr) == 0 {
// 		return -1
// 	}
// 	min := -1
// 	for _, v := range arr {
// 		if min == -1 {
// 			min = v
// 		} else if min > v && v >= 0 {
// 			min = v
// 		}
// 	}
// 	return min
// }

//DP 이용하여 다시풀것
func coinChange(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	init := 100000
	currMin := &init //currMin 추가 : 현재보다 더 큰 연산수를 막기위해 추가됨

	return dv(0, amount, coins, 0, 0, currMin)
}
func dv(sum int, amount int, coins []int, coinIdx int, depth int, currMin *int) int {
	if sum == amount {
		if *currMin > depth {
			*currMin = depth
		}
		return depth
	} else if sum > amount || coinIdx >= len(coins) {
		return -1
	}
	target := (amount - sum)
	min := -1
	v := coins[coinIdx]
	t := -1
	d1 := target / v
	d2 := target % v
	if depth+d1 > *currMin {
		return -1
	}
	if target < v {
		return dv(sum, amount, coins, coinIdx+1, depth, currMin)
	}
	for d1 >= 0 {
		tt := dv(amount-d2, amount, coins, coinIdx+1, depth+d1, currMin)
		if t == -1 {
			t = tt
		} else if tt >= 0 && tt < t {
			t = tt
		}
		d1--
		d2 += v
	}
	if min == -1 {
		min = t
	} else if min > t && t >= 0 {
		min = t
	}
	return min
}
