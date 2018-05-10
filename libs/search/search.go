package search

import "errors"

// lower_bound returns the first index of array "a".
// The index satisfied with a condition x <= a[index].
// The complexity is O(logN).
//
// Example:
//      a := []int{1, 11, 22, 33, 44, 55, 66}
//	idx, err := lower_bound(a, 36)
// Output:
//      idx is equals to 4
//
// If all elements of array "a" are smaller than x, lower_bound returns -1 and error.
//
// Example:
//      a := []int{1, 11, 22, 33, 44, 55, 66}
//	idx, err := lower_bound(a, 100)
// Output:
//      idx is equals to -1, and error is not nil
func lower_bound(a []int, x int) (int, error) {
	l, r := 0, len(a)-1
	if a[r] < x {
		e := errors.New("nothing")
		return -1, e
	}
	for l < r {
		mid := int((l + r) / 2)
		if x <= a[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l, nil
}

func UpperBound(a []int, x int) {

}
