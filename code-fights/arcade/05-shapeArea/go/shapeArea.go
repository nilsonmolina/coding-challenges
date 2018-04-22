package main

func main() {}

func shapeArea(n int) int {
	// n^2 + (n-1)^2
	return (n * n) + ((n - 1) * (n - 1))
}

// --------------------------------------------------
// FIRST SOLUTION - NAIVE
//      O(n) execution, much slower for larger
//      numbers. Above solution is O(1).
// --------------------------------------------------
// func shapeArea(n int) int {
//   area := 1
//   for i := 1; i < n; i++ {
//     area += 4 * i
//   }
//   return area
// }
