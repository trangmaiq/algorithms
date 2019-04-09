package algorithms

import "math"

// Given a 32-bit signed integer, reverse digits of an integer.
func ReverseInteger(x int) (rev int){
	for x != 0 {
		var temp int = x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 / 10 && temp > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && temp < - 8) {
			return 0
		}
		rev = rev * 10 + temp
	}
	return rev
}

