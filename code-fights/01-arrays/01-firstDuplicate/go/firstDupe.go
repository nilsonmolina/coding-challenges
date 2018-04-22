func firstDuplicate(a []int) int {
	dupe := -1
	vals := map[int]int{} //map[number]index

	for i := range a {
		if _, ok := vals[a[i]]; ok {
			dupe = a[i]
			break
		} else {
			vals[a[i]] = i
		}
	}
	return dupe
}

// --------------------------------------------------
// THIS WAS WRONG
//      It assumed that I wanted the duplicate
//      that was closest to its second occurrence
// --------------------------------------------------
// func firstDuplicate(a []int) int {
//     dupe := struct {
// 		value    int
// 		distance int
// 	}{-1, len(a)}
//     vals := map[int]int{} //map[number]index

//     for i := range a {
//         if v, ok := vals[a[i]]; !ok {
//             vals[a[i]] = i
//         } else if (i - v) < dupe.distance {
//             dupe.value = a[i]
//             dupe.distance = i - v
//         }
//     }
//     return dupe.value
// }