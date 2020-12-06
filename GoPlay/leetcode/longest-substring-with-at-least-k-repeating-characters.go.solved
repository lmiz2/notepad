package main

import (
	"sort"
)

func longestSubstring(s string, k int) int {

	locates := make([][]int, 26) // char - 97 is index

	for i, v := range s {
		locates[v-97] = append(locates[v-97], i)
	}

	return findLongestStrings(s, 0, len(s)-1, k, locates)
}
func findLongestStrings(s string, start int, end int, k int, dict [][]int) int {
	bannedPoints := []int{}
	skipList := make(map[rune]bool)
	for _, v := range s[start : end+1] {
		if !skipList[v] {
			bannedPoints = append(bannedPoints, findBannedPointsForLetter(v, start, end, k, dict)...)
			skipList[v] = true
		}
	}
	if len(bannedPoints) == 0 {
		return len(s[start : end+1])
	}
	rs := 0
	sort.Ints(bannedPoints)
	lidx := 0
	ridx := 0
	for _, v := range bannedPoints {
		ridx = v - 1
		if ridx-lidx >= k-1 && lidx >= 0 && ridx <= end {
			ls := findLongestStrings(s, lidx, ridx, k, dict)
			if ls > rs {
				rs = ls
			}
		}
		lidx = v + 1
		ridx = v + 1
	}
	if ridx < end {
		ridx = end
		if ridx-lidx >= k-1 {
			ls := findLongestStrings(s, lidx, ridx, k, dict)
			if ls > rs {
				rs = ls
			}
		}
	}

	return rs
}
func findBannedPointsForLetter(char rune, start int, end int, k int, dict [][]int) []int {
	rs := []int{}

	for _, v := range dict[char-97] {
		if v >= start && v <= end {
			rs = append(rs, v)
		}
	}
	if len(rs) >= k {
		return []int{}
	}
	return rs
}
