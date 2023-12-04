package utils

import (
	"strconv"
	"strings"
)

func GetLines(input string) []string {
	return strings.Split(input, "\n")
}

func ToIntSlice(input string, separator string) []int {
	var ret []int
	for _, str := range strings.Split(input, separator) {
		if str == "" {
			continue
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ret = append(ret, n)
	}
	return ret
}
