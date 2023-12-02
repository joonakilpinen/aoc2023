#!/usr/bin/env bash

mkdir -p "solutions/day${1}part2"

get_assignment() {
  curl --cookie "$(< .cookie)" "https://adventofcode.com/2023/day/${1}" |
  lynx -dump -nolist -stdin |
  sed "0,/--- Part Two/d"
}

assignment=$(get_assignment "$1")
assignment_comment=$(echo "${assignment}" | sed 's|^|// |')

cat <<EOF > "solutions/day${1}part2/solution.go"
package day${1}part2

type Solver struct{}

func (Solver) Solve(input string) string {
	return ""
}

${assignment_comment}
EOF

echo "${assignment}"
read -p "Expected test result? "

cat <<EOF >> aoc2023_test.go

func TestDay${1}Part2SolveTest(t *testing.T) {
	Assert(t, &day${1}part2.Solver{}, GetTestInput(${1}), "${REPLY}")
}

func TestDay${1}Part2Solve(t *testing.T) {
	Solve(&day${1}part2.Solver{}, GetInput(${1}))
}
EOF