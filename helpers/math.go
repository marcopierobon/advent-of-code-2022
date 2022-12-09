package helpers

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
