package main

import (
	"aoc2023/solutions"
	"aoc2023/solutions/day1part1"
	"aoc2023/solutions/day1part2"
	"aoc2023/solutions/day2part1"
	"aoc2023/solutions/day2part2"
	"aoc2023/solutions/day3part1"
	"aoc2023/solutions/day3part2"
	"aoc2023/solutions/day4part1"
	"aoc2023/solutions/day4part2"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Solve(solver solutions.Solver, input string) string {
	solution := solver.Solve(input)
	log.Printf("Solution:\n%s", solution)
	return solution
}

func Assert(t *testing.T, solver solutions.Solver, input string, expectedResult string) {
	assert.Equal(t, expectedResult, Solve(solver, input))
}

func TestDay1Part1SolveTest(t *testing.T) {
	Assert(t, &day1part1.Solver{}, GetTestInput(1), "142")
}

func TestDay1Part1Solve(t *testing.T) {
	Solve(&day1part1.Solver{}, GetInput(1))
}

func TestDay1Part2SolveTest(t *testing.T) {
	input := `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
	Assert(t, &day1part2.Solver{}, input, "281")
}

func TestDay1Part2Solve(t *testing.T) {
	Solve(&day1part2.Solver{}, GetInput(1))
}

func TestDay2Part1SolveTest(t *testing.T) {
	Assert(t, &day2part1.Solver{}, GetTestInput(2), "8")
}

func TestDay2Part1Solve(t *testing.T) {
	Solve(&day2part1.Solver{}, GetInput(2))
}

func TestDay2Part2SolveTest(t *testing.T) {
	Assert(t, &day2part2.Solver{}, GetTestInput(2), "2286")
}

func TestDay2Part2Solve(t *testing.T) {
	Solve(&day2part2.Solver{}, GetInput(2))
}

func TestDay3Part1SolveTest(t *testing.T) {
	Assert(t, &day3part1.Solver{}, GetTestInput(3), "4361")
}

func TestDay3Part1Solve(t *testing.T) {
	Solve(&day3part1.Solver{}, GetInput(3))
}

func TestDay3Part2SolveTest(t *testing.T) {
	Assert(t, &day3part2.Solver{}, GetTestInput(3), "467835")
}

func TestDay3Part2Solve(t *testing.T) {
	Solve(&day3part2.Solver{}, GetInput(3))
}

func TestDay4Part1SolveTest(t *testing.T) {
	Assert(t, &day4part1.Solver{}, GetTestInput(4), "13")
}

func TestDay4Part1Solve(t *testing.T) {
	Solve(&day4part1.Solver{}, GetInput(4))
}

func TestDay4Part2SolveTest(t *testing.T) {
	Assert(t, &day4part2.Solver{}, GetTestInput(4), "30")
}

func TestDay4Part2Solve(t *testing.T) {
	Solve(&day4part2.Solver{}, GetInput(4))
}
