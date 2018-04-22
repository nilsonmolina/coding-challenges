package main

func main() {

}

func centuryFromYear(year int) int {
	c := year / 100
	if (year % 100) != 0 {
		c++
	}
	return c
}

// --------------------------------------------------
// SHORTER SOLUTION
// --------------------------------------------------
// func centuryFromYear(year int) int {
//     return 1 + (year - 1) / 100
// }
