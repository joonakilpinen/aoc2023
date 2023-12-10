package utils

func GCD(a int, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func LCM(ints []int) int {
	l := ints[0]
	for i := 1; i < len(ints); i++ {
		g := GCD(l, ints[i])
		l = l * ints[i] / g
	}
	return l
}
