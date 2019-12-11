package rmath

// Abs returns the absolute value of the given int
func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
