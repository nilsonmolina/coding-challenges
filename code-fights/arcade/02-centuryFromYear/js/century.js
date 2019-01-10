// // FIRST SOLUTION
// function centuryFromYear(year) {
//     return year % 100 !== 0 ? Math.floor(year / 100) + 1 : year / 100;
// }
//

function centuryFromYear(year) {
  return Math.ceil(year / 100);
}
