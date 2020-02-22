package rmath

import "math"

// Min returns the minimum of all passed ints.
// if no ints are passed, then math.MaxInt64 is returned
func Min(n ...int) (min int) {
	min = math.MaxInt64
	for _, val := range n {
		if val < min {
			min = val
		}
	}
	return
}

// Max returns the maximum of all passed ints.
// if no ints are passed, then math.MinInt64 is returned
func Max(n ...int) (max int) {
	max = math.MinInt64
	for _, val := range n {
		if val > max {
			max = val
		}
	}
	return
}

// Sum returns the sum of all passed ints.
// if no ints are passed, then 0 is returned
func Sum(n ...int) (sum int) {
	sum = 0
	for _, val := range n {
		sum += val
	}
	return
}
