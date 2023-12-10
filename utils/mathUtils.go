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
	l := 0
	for i := 1; i < len(ints); i++ {
		if i == 1 {
			g := GCD(ints[0], ints[1])
			l = ints[0] * ints[1] / g
		} else {
			g := GCD(l, ints[i])
			l = l * ints[i] / g
		}
	}
	return l
}
