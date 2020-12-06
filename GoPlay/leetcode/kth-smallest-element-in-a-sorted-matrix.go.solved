package main

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	th := 0
	thNum := 0
	firstIdxs := make([]int, len(matrix))
	for th < k {
		for i, v := range matrix {
			idx := firstIdxs[i]
			if idx < n {
				tmpVal := v[idx]
				isChanged := false
				minVal := tmpVal
				minValNx := -1
				for nx := i + 1; nx < n; nx++ {
					if firstIdxs[nx] < n && matrix[nx][firstIdxs[nx]] < minVal {
						minVal = matrix[nx][firstIdxs[nx]]
						minValNx = nx
						isChanged = true
					}
				}
				if !isChanged {
					th++
					thNum = tmpVal
					firstIdxs[i]++
				} else {
					thNum = minVal
					th++
					firstIdxs[minValNx]++
				}
				break
			}
		}
	}
	return thNum
}
