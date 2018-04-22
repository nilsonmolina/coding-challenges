package main

func main() {

}

func checkPalindrome(inputString string) bool {
	tail := len(inputString) - 1
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[tail-i] {
			return false
		}
	}
	return true
}

// --------------------------------------------------
// FIRST SOLUTION - NAIVE
//      Do not need to loop through entire array,
//      going through half is enough.
// --------------------------------------------------
// func checkPalindrome(inputString string) bool {
//     tail := len(inputString) - 1
//     for i := range inputString {
//         if inputString[i] != inputString[tail - i] {
//             return false
//         }
//     }
//     return true
// }
