package main

import "sort"

func main() {}

func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	count := 0
	for i := 0; i+1 < len(statues); i++ {
		if statues[i+1]-statues[i] > 1 {
			count += statues[i+1] - statues[i] - 1
		}
	}
	return count
}
