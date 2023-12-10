package main

import (
	"aoc2023/solutions"
	"aoc2023/solutions/day10part1"
	"aoc2023/solutions/day10part2"
	"aoc2023/solutions/day1part1"
	"aoc2023/solutions/day1part2"
	"aoc2023/solutions/day2part1"
	"aoc2023/solutions/day2part2"
	"aoc2023/solutions/day3part1"
	"aoc2023/solutions/day3part2"
	"aoc2023/solutions/day4part1"
	"aoc2023/solutions/day4part2"
	"aoc2023/solutions/day5part1"
	"aoc2023/solutions/day5part2"
	"aoc2023/solutions/day6part1"
	"aoc2023/solutions/day6part2"
	"aoc2023/solutions/day7part1"
	"aoc2023/solutions/day7part2"
	"aoc2023/solutions/day8part1"
	"aoc2023/solutions/day8part2"
	"aoc2023/solutions/day9part1"
	"aoc2023/solutions/day9part2"
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

func TestDay5Part1SolveTest(t *testing.T) {
	Assert(t, &day5part1.Solver{}, GetTestInput(5), "35")
}

func TestDay5Part1Solve(t *testing.T) {
	Solve(&day5part1.Solver{}, GetInput(5))
}

func TestDay5Part2SolveTest(t *testing.T) {
	Assert(t, &day5part2.Solver{}, GetTestInput(5), "46")
}

func TestDay5Part2Solve(t *testing.T) {
	Solve(&day5part2.Solver{}, GetInput(5))
}

func TestDay6Part1SolveTest(t *testing.T) {
	Assert(t, &day6part1.Solver{}, GetTestInput(6), "288")
}

func TestDay6Part1Solve(t *testing.T) {
	Solve(&day6part1.Solver{}, GetInput(6))
}

func TestDay6Part2SolveTest(t *testing.T) {
	Assert(t, &day6part2.Solver{}, GetTestInput(6), "71503")
}

func TestDay6Part2Solve(t *testing.T) {
	Solve(&day6part2.Solver{}, GetInput(6))
}

func TestDay7Part1SolveTest(t *testing.T) {
	Assert(t, &day7part1.Solver{}, GetTestInput(7), "6440")
}

func TestDay7Part1Solve(t *testing.T) {
	Solve(&day7part1.Solver{}, GetInput(7))
}

func TestDay7Part2SolveTest(t *testing.T) {
	Assert(t, &day7part2.Solver{}, GetTestInput(7), "5905")
}

func TestDay7Part2Solve(t *testing.T) {
	Solve(&day7part2.Solver{}, GetInput(7))
}

func TestDay8Part1SolveTest(t *testing.T) {
	Assert(t, &day8part1.Solver{}, GetTestInput(8), "2")
}

func TestDay8Part1Solve(t *testing.T) {
	Solve(&day8part1.Solver{}, GetInput(8))
}

func TestDay8Part2SolveTest(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`
	Assert(t, &day8part2.Solver{}, input, "6")
}

func TestDay8Part2Solve(t *testing.T) {
	Solve(&day8part2.Solver{}, GetInput(8))
}

func TestDay9Part1SolveTest(t *testing.T) {
	Assert(t, &day9part1.Solver{}, GetTestInput(9), "114")
}

func TestDay9Part1Solve(t *testing.T) {
	Solve(&day9part1.Solver{}, GetInput(9))
}

func TestDay9Part2SolveTest(t *testing.T) {
	Assert(t, &day9part2.Solver{}, GetTestInput(9), "2")
}

func TestDay9Part2Solve(t *testing.T) {
	Solve(&day9part2.Solver{}, GetInput(9))
}

func TestDay10Part1SolveTest(t *testing.T) {
	Assert(t, &day10part1.Solver{}, GetTestInput(10), "8")
}

func TestDay10Part1Solve(t *testing.T) {
	Solve(&day10part1.Solver{}, GetInput(10))
}

func TestDay10Part2SolveTest(t *testing.T) {
	input := `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`
	Assert(t, &day10part2.Solver{}, input, "10")
}

func TestDay10Part2Solve(t *testing.T) {
	Solve(&day10part2.Solver{}, GetInput(10))
}
