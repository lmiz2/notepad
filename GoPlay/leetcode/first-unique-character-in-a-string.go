package main

func firstUniqChar(s string) int {
	rs := -1
	mem := [26]int{}
	for _, v := range s {
		mem[v-97] = mem[v-97] + 1

	}
	for i, v := range s {
		if mem[v-97] == 1 {
			return i
		}

	}
	return rs
}
