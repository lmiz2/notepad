package main

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		idx := findInsertPoint(dp, v)
		if len(dp) == idx {
			dp = append(dp, v)
		} else {
			dp[idx] = v
		}
	}
	return len(dp)
}

func findInsertPoint(arr []int, target int) int {
	for i, v := range arr {
		if v >= target {
			return i
		}
	}
	return len(arr)
}

// func lengthOfLIS(nums []int) int {
// 	mem := make(map[int]int)
// 	return recursive(nums, false, -10001, mem)
// }
// func recursive(nums []int, choosed bool, beforeNum int, mem map[int]int) int {
// 	rs := 0
// 	if choosed {
// 		rs++
// 	}
// 	if len(nums) == 0 {
// 		return rs
// 	}

// 	tKey := beforeNum + (len(nums) * 100000)
// 	memVal, hits := mem[tKey]
// 	if hits {
// 		return memVal + rs
// 	}
// 	t2 := 0
// 	if nums[0] > beforeNum {
// 		t2 = recursive(nums[1:], true, nums[0], mem)
// 	}
// 	t1 := recursive(nums[1:], false, beforeNum, mem)

// 	if t1 > t2 {
// 		mem[tKey] = t1
// 		rs += t1
// 	} else {
// 		mem[tKey] = t2
// 		rs += t2
// 	}
// 	return rs
// }
