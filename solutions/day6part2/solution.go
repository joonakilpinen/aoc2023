package day6part2

import (
	"aoc2023/solutions/day6part1"
	"aoc2023/utils"
	"errors"
	"strconv"
	"strings"
)

type Solver struct{}

func ParseRace(lines []string) day6part1.Race {
	timeWithoutPrefix, ok := strings.CutPrefix(lines[0], "Time:")
	if !ok {
		panic(errors.New("parsing error"))
	}
	time := utils.RemoveWhitespace(timeWithoutPrefix)
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		panic(err)
	}
	distWithoutPrefix, ok := strings.CutPrefix(lines[1], "Distance:")
	if !ok {
		panic(errors.New("parsing error"))
	}
	dist := utils.RemoveWhitespace(distWithoutPrefix)
	distInt, err := strconv.Atoi(dist)
	if err != nil {
		panic(err)
	}
	return day6part1.Race{Time: timeInt, Distance: distInt}
}

func countWaysToBeatRecord(race day6part1.Race) int {
	count := 0
	for i := race.Time - 1; i > 0; i-- {
		timeLeft := race.Time - i
		distance := i * timeLeft
		if distance > race.Distance {
			count++
		} else if count > 0 {
			break
		}
	}
	return count
}

func (Solver) Solve(input string) string {
	race := ParseRace(utils.GetLines(input))

	return strconv.Itoa(countWaysToBeatRecord(race))
}

//
//    As the race is about to start, you realize the piece of paper with race
//    times and record distances you got earlier actually just has very bad
//    kerning. There's really only one race - ignore the spaces between the
//    numbers on each line.
//
//    So, the example from before:
// Time:      7  15   30
// Distance:  9  40  200
//
//    ...now instead means this:
// Time:      71530
// Distance:  940200
//
//    Now, you have to figure out how many ways there are to win this single
//    race. In this example, the race lasts for 71530 milliseconds and the
//    record distance you need to beat is 940200 millimeters. You could hold
//    the button anywhere from 14 to 71516 milliseconds and beat the record,
//    a total of 71503 ways!
//
//    How many ways can you beat the record in this one much longer race?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
