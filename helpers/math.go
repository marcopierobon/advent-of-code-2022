package helpers

import (
	"math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func Max(array []int) int {
	var max = math.MinInt
	for _, element := range array {
		if element > max && element != max {
			max = element
		}
	}
	return max
}
