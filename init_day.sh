#!/usr/bin/env bash

mkdir -p "inputs/day${1}" "solutions/day${1}part1"

get_assignment() {
  curl --cookie "$(< .cookie)" "https://adventofcode.com/2023/day/${1}" |
  lynx -dump -nolist -stdin |
  sed "0,/--- Day ${1}/d"
}

assignment=$(get_assignment "$1")
assignment_comment=$(echo "${assignment}" | sed 's|^|// |')

curl --cookie "$(< .cookie)" "https://adventofcode.com/2023/day/${1}/input" > "inputs/day${1}/input.txt"

cat <<EOF > "solutions/day${1}part1/solution.go"
package day${1}part1

type Solver struct{}

func (Solver) Solve(input string) string {
	return ""
}

${assignment_comment}
EOF

echo "${assignment}"
echo "Paste test input and press Ctrl+D"
mapfile input

for line in "${input[@]}"; do
  printf "%s" "$line"
done > "inputs/day${1}/test.txt"

read -p "Expected test result? "

cat <<EOF >> aoc2023_test.go

func TestDay${1}Part1SolveTest(t *testing.T) {
	Assert(t, &day${1}part1.Solver{}, GetTestInput(${1}), "${REPLY}")
}

func TestDay${1}Part1Solve(t *testing.T) {
	Solve(&day${1}part1.Solver{}, GetInput(${1}))
}
EOF