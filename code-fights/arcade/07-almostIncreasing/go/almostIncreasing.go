package main

import (
	"fmt"
)

func main() {
	fmt.Println(almostIncreasingSequence([]int{1, 3, 2}))
}

func almostIncreasingSequence(sequence []int) bool {
	remove := 0
	for i := 1; i < len(sequence)-1; i++ {
		if sequence[i-1] >= sequence[i] {
			if remove == 1 || sequence[i-1] >= sequence[i+1] {
				return false
			}
			remove = 1
		}
	}
	return true
}
