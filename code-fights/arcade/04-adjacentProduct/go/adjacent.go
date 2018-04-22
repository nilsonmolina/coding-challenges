package main

func main() {}

func adjacentElementsProduct(inputArray []int) int {
	var max int
	for i := 0; i+1 < len(inputArray); i++ {
		if i == 0 || inputArray[i]*inputArray[i+1] > max {
			max = inputArray[i] * inputArray[i+1]
		}
	}
	return max
}
