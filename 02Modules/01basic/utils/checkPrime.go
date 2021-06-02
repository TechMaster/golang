package utils

import "math"

func IsPrime(number int) bool {
	if number <= 2 {
		return true
	}
	sq_root := int(math.Sqrt(float64(number)))
	for i := 2; i <= sq_root; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
