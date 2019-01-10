// --------------------------------------------------
// FIRST SOLUTION - NAIVE
//      Do not need 'head' variable
// --------------------------------------------------
// function checkPalindrome(inputString) {
//     var head = 0;
//     var tail = inputString.length - 1;
//     while (head < tail) {
//         if (inputString[head] != inputString[tail]) {
//             return false
//         }
//         head++;
//         tail--;
//     }
//     return true;
// }

// // --------------------------------------------------
// COMMUNITY SOLUTION - SLOWER
//      After benchmarking, this solution proves to
//      be slightly slower.
// --------------------------------------------------
// function checkPalindrome(inputString) {
//     return inputString == inputString.split('').reverse().join('');
// }

function checkPalindrome(inputString) {
  let tail = inputString.length - 1;
  for (let i = 0; i < tail; i++) {
    if (inputString[i] !== inputString[tail]) {
      return false;
    }
    tail--;
  }
  return true;
}


checkPalindrome('racecar');
