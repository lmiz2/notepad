package main

//{3, 5, 4, 0, 1, 0, 1, 5, 2, 8, 9}

//[-1 -1 -1 5 -1 1 -1 -1 -1 -1 -1]

func avoidFlood(rains []int) []int {
	result := make([]int, len(rains))
	mem := make(map[int]bool)
	lastDryIdxs := make(map[int]int)
	for i, v := range rains {
		if v > 0 {
			if isFull := mem[v]; isFull {
				if lastDryIdx == -1 { //lastDryIdx 를 스택으로 구현..?
					return []int{}
				}
				if islastFull := mem[result[lastDryIdx]]; islastFull {
					return []int{}
				}
				delete(mem, v)
				mem[result[lastDryIdx]] = true
				result[lastDryIdx] = v
			} else {
				mem[v] = true
			}
			result[i] = -1
		} else {
			lastDryIdx = i
			ramIdx := -1
			for k, v := range mem {
				if v {
					ramIdx = k
					break
				}
			}
			if ramIdx == -1 {
				result[i] = 1
			} else {
				result[i] = ramIdx
				delete(mem, ramIdx)
			}
		}
	}
	return result
}
