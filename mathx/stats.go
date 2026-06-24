// Package mathx holds small numeric helpers built on the standard library.
package mathx

import "errors"

// ErrNoNumbers is returned when a calculation needs at least one number
// but none were supplied.
var ErrNoNumbers = errors.New("mathx: no numbers given")

// MinMax returns the smallest and largest of the given numbers.
// It is variadic, so it accepts any number of ints, and it uses named
// returns (min, max) so the values it reports are clear from the signature.
func MinMax(nums ...int) (min, max int) {
	if len(nums) == 0 {
		return 0, 0
	}

	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		min = smaller(min, n)
		max = larger(max, n)
	}
	return min, max
}

// Average returns the mean of the given numbers. It returns an error when
// no numbers are supplied, because dividing by zero has no answer. This is
// the (value, error) pair you will see all over Go.
func Average(nums ...int) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrNoNumbers
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums)), nil
}

// smaller returns the smaller of two ints. It is unexported (lowercase),
// so only code inside package mathx can call it.
func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// larger returns the larger of two ints. It is unexported, too.
func larger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
