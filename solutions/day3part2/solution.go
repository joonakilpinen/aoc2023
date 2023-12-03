package day3part2

import (
	"aoc2023/utils"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type Solver struct{}

func GetMatchNumAndIndices(line string, match []int) (int, int, int) {
	text := line[utf8.RuneCountInString(line[:match[0]]):utf8.RuneCountInString(line[:match[1]])]
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num, match[0], match[1] - 1
}

var gearRegex = regexp.MustCompile("\\*")
var numberRegex = regexp.MustCompile("\\d+")

type PartWithIndices struct {
	num   int
	start int
	end   int
}

func GetPartsAndIndices(line *string) []PartWithIndices {
	var partsWithIndices []PartWithIndices
	matches := numberRegex.FindAllStringSubmatchIndex(*line, -1)
	for _, match := range matches {
		num, start, end := GetMatchNumAndIndices(*line, match)
		partsWithIndices = append(partsWithIndices, PartWithIndices{num, start, end})
	}
	return partsWithIndices
}

func FilterPartsByIndice(parts []PartWithIndices, index int) []int {
	var justParts []int
	for _, pi := range parts {
		if pi.start <= index && pi.end >= index {
			justParts = append(justParts, pi.num)
			continue
		}
		if pi.start == index+1 || pi.end == index-1 {
			justParts = append(justParts, pi.num)
			continue
		}
	}
	return justParts
}

func GetAdjacentParts(previous *string, current *string, next *string, index int) []int {
	var adjacentParts []int
	pis := GetPartsAndIndices(current)
	adjacentParts = append(adjacentParts, FilterPartsByIndice(pis, index)...)
	if previous != nil {
		pis := GetPartsAndIndices(previous)
		adjacentParts = append(adjacentParts, FilterPartsByIndice(pis, index)...)
	}
	if next != nil {
		pis := GetPartsAndIndices(next)
		adjacentParts = append(adjacentParts, FilterPartsByIndice(pis, index)...)
	}
	return adjacentParts
}

func GetGearRatios(previous *string, current *string, next *string) []int {
	var ratios []int
	matches := gearRegex.FindAllStringSubmatchIndex(*current, -1)
	for _, match := range matches {
		adjacentParts := GetAdjacentParts(previous, current, next, match[0])
		if len(adjacentParts) == 2 {
			ratios = append(ratios, adjacentParts[0]*adjacentParts[1])
		}
	}
	return ratios
}

func (Solver) Solve(input string) string {
	lines := utils.GetLines(input)
	var respChan = make(chan []int, len(lines))
	defer close(respChan)
	for i := range lines {
		var previous *string
		var next *string
		if i > 0 {
			previous = &lines[i-1]
		}
		if i < len(lines)-1 {
			next = &lines[i+1]
		}
		line := &lines[i]
		go func() {
			respChan <- GetGearRatios(previous, line, next)
		}()
	}
	var parts []int
	for range lines {
		parts = append(parts, <-respChan...)
	}
	sum := 0
	for _, part := range parts {
		sum += part
	}
	return strconv.Itoa(sum)
}

//
//    The engineer finds the missing part and installs it in the engine! As
//    the engine springs to life, you jump in the closest gondola, finally
//    ready to ascend to the water source.
//
//    You don't seem to be going very fast, though. Maybe something is still
//    wrong? Fortunately, the gondola has a phone labeled "help", so you pick
//    it up and the engineer answers.
//
//    Before you can explain the situation, she suggests that you look out
//    the window. There stands the engineer, holding a phone in one hand and
//    waving with the other. You're going so slowly that you haven't even
//    left the station. You exit the gondola.
//
//    The missing part wasn't the only issue - one of the gears in the engine
//    is wrong. A gear is any * symbol that is adjacent to exactly two part
//    numbers. Its gear ratio is the result of multiplying those two numbers
//    together.
//
//    This time, you need to find the gear ratio of every gear and add them
//    all up so that the engineer can figure out which gear needs to be
//    replaced.
//
//    Consider the same engine schematic again:
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
//
//    In this schematic, there are two gears. The first is in the top left;
//    it has part numbers 467 and 35, so its gear ratio is 16345. The second
//    gear is in the lower right; its gear ratio is 451490. (The * adjacent
//    to 617 is not a gear because it is only adjacent to one part number.)
//    Adding up all of the gear ratios produces 467835.
//
//    What is the sum of all of the gear ratios in your engine schematic?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
