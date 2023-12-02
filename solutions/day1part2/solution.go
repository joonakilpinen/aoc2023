package day1part2

import (
	"aoc2023/solutions/day1part1"
	"aoc2023/utils"
	"strconv"
	"strings"
)

type Solver struct{}

var nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func find(line string) (string, string, int, int) {
	tmp := 9000
	tmpKey := ""
	for key := range nums {
		index := strings.Index(line, key)
		if index != -1 && index < tmp {
			tmp = index
			tmpKey = key
		}
	}
	firstNum := tmpKey
	firstIdx := tmp
	tmp = -1
	tmpKey = ""
	for key := range nums {
		index := strings.LastIndex(line, key)
		if index > tmp {
			tmp = index
			tmpKey = key
		}
	}
	lastNum := tmpKey
	lastIdx := tmp
	return firstNum, lastNum, firstIdx, lastIdx
}

func (Solver) Solve(input string) string {
	lines := utils.GetLines(input)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		first, last, firstIdx, lastIdx := find(line)
		if first == "" {
			continue
		}
		if firstIdx == lastIdx {
			// First and last are the same substring
			lines[i] = strings.Replace(line, first, strconv.Itoa(nums[first]), 1)
			continue
		}
		if firstIdx+len(first) >= lastIdx {
			// Overlapping
			lines[i] = strings.Replace(line, first, strconv.Itoa(nums[first])+strconv.Itoa(nums[last]), 1)
			continue
		}
		// log.Printf("Line: %s, First: %s, Last: %s", line, first, last)
		line = strings.Replace(line, first, strconv.Itoa(nums[first]), 1)
		lastIndex := strings.LastIndex(line, last)
		lines[i] = line[:lastIndex] + strings.Replace(line[lastIndex:], last, strconv.Itoa(nums[last]), 1)
	}
	preprocessedInput := strings.Join(lines, "\n")
	// log.Println(preprocessedInput)
	return day1part1.Solver{}.Solve(preprocessedInput)
}

//
//    Your calculation isn't quite right. It looks like some of the digits
//    are actually spelled out with letters: one, two, three, four, five,
//    six, seven, eight, and nine also count as valid "digits".
//
//    Equipped with this new information, you now need to find the real first
//    and last digit on each line. For example:
// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
//
//    In this example, the calibration values are 29, 83, 13, 24, 42, 14, and
//    76. Adding these together produces 281.
//
//    What is the sum of all of the calibration values?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
