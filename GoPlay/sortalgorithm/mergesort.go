package sortalgorithm

func Mergesort(nums *[]int, left int, right int, rsArr *[]int) {
	if left < right {
		mid := (left + right) / 2
		Mergesort(nums, left, mid, rsArr)
		Mergesort(nums, mid+1, right, rsArr)
		merge(nums, left, right, mid, rsArr)
	}
}
func merge(nums *[]int, left int, right int, mid int, rsArr *[]int) {
	lIdx := left
	rIdx := mid + 1
	rsIdx := left

	for lIdx <= mid && rIdx <= right {
		if (*nums)[lIdx] >= (*nums)[rIdx] {
			(*rsArr)[rsIdx] = (*nums)[rIdx]
			rsIdx++
			rIdx++
		} else {
			(*rsArr)[rsIdx] = (*nums)[lIdx]
			rsIdx++
			lIdx++
		}
	}
	if lIdx <= mid {
		for lIdx <= mid {
			(*rsArr)[rsIdx] = (*nums)[lIdx]
			rsIdx++
			lIdx++
		}
	} else {
		for rIdx <= right {
			(*rsArr)[rsIdx] = (*nums)[rIdx]
			rsIdx++
			rIdx++
		}
	}

	for i := left; i <= right; i++ {
		(*nums)[i] = (*rsArr)[i]
	}
}
