// Package mathx holds small numeric helpers built on the standard library.
package mathx

// MinMax returns the smallest and largest of the given numbers.
// It is variadic, so it accepts any number of ints.
func MinMax(nums ...int) (min, max int) {
	// TODO: handle the empty case, then loop over nums and track the
	// smallest and largest value with the smaller/larger helpers you
	// will add as unexported functions in this file.
	_ = nums
	return 0, 0
}

// Average returns the mean of the given numbers, or an error when no
// numbers are supplied (you cannot divide by zero).
func Average(nums ...int) (float64, error) {
	// TODO: return an error when len(nums) == 0, otherwise sum the
	// numbers and divide by the count.
	_ = nums
	return 0, nil
}
