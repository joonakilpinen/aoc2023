package day3part1

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
	return num, match[0], match[1]
}

var symbolsRegex = regexp.MustCompile("[^a-zA-Z\\d\\s.]")
var numberRegex = regexp.MustCompile("\\d+")

func SymbolInRange(line string, start int, end int) bool {
	if len(line) == 0 {
		return false
	}
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	matches := symbolsRegex.FindAllString(line[start:end], -1)
	return len(matches) > 0
}

func GetPartNumbers(previous *string, current *string, next *string) []int {
	var partNumbers []int
	matches := numberRegex.FindAllStringSubmatchIndex(*current, -1)
	for _, match := range matches {
		num, start, end := GetMatchNumAndIndices(*current, match)
		if SymbolInRange(*current, start-1, end+1) {
			partNumbers = append(partNumbers, num)
			continue
		}
		if previous != nil && SymbolInRange(*previous, start-1, end+1) {
			partNumbers = append(partNumbers, num)
			continue
		}
		if next != nil && SymbolInRange(*next, start-1, end+1) {
			partNumbers = append(partNumbers, num)
		}
	}
	return partNumbers
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
			respChan <- GetPartNumbers(previous, line, next)
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
//    You and the Elf eventually reach a gondola lift station; he says the
//    gondola lift will take you up to the water source, but this is as far
//    as he can bring you. You go inside.
//
//    It doesn't take long to find the gondolas, but there seems to be a
//    problem: they're not moving.
//
//    "Aaah!"
//
//    You turn around to see a slightly-greasy Elf with a wrench and a look
//    of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't
//    working right now; it'll still be a while before I can fix it." You
//    offer to help.
//
//    The engineer explains that an engine part seems to be missing from the
//    engine, but nobody can figure out which one. If you can add up all the
//    part numbers in the engine schematic, it should be easy to work out
//    which part is missing.
//
//    The engine schematic (your puzzle input) consists of a visual
//    representation of the engine. There are lots of numbers and symbols you
//    don't really understand, but apparently any number adjacent to a
//    symbol, even diagonally, is a "part number" and should be included in
//    your sum. (Periods (.) do not count as a symbol.)
//
//    Here is an example engine schematic:
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
//    In this schematic, two numbers are not part numbers because they are
//    not adjacent to a symbol: 114 (top right) and 58 (middle right). Every
//    other number is adjacent to a symbol and so is a part number; their sum
//    is 4361.
//
//    Of course, the actual engine schematic is much larger. What is the sum
//    of all of the part numbers in the engine schematic?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
