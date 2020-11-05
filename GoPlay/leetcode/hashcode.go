package main

import "hash/fnv"

func hashCode(nums []int) int {
	result := nums[0]
	for _, v := range nums[1:] {
		result = 31*result + v
	}
	return result
}

func FNV32a(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}
