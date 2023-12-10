package day8part2

import (
	"aoc2023/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

type Node struct {
	Left  string
	Right string
}

type Instruction string

const (
	L Instruction = "L"
	R Instruction = "R"
)

type Maps struct {
	StartingPoints []string
	Instructions   []Instruction
	Nodes          map[string]Node
}

var nodeRegex = regexp.MustCompile("^(?P<Id>[A-Z0-9]{3}) = \\((?P<Left>[A-Z0-9]{3}), (?P<Right>[A-Z0-9]{3})\\)")

func parseMaps(input string) Maps {
	lines := utils.GetLines(input)
	var instructions []Instruction
	for _, inst := range strings.Split(lines[0], "") {
		instructions = append(instructions, Instruction(inst))
	}
	nodes := map[string]Node{}
	var startingPoints []string
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		matches := nodeRegex.FindStringSubmatch(lines[i])
		id := matches[nodeRegex.SubexpIndex("Id")]
		left := matches[nodeRegex.SubexpIndex("Left")]
		right := matches[nodeRegex.SubexpIndex("Right")]
		if id[len(id)-1:] == "A" {
			startingPoints = append(startingPoints, id)
		}
		nodes[id] = Node{Left: left, Right: right}
	}
	return Maps{
		StartingPoints: startingPoints,
		Instructions:   instructions,
		Nodes:          nodes,
	}
}

func doStep(stepInst StepAndInst, maps Maps) (string, bool) {
	node := maps.Nodes[stepInst.step]
	var step string
	switch stepInst.inst {
	case L:
		step = node.Left
	case R:
		step = node.Right
	}
	return step, stepInst.step[len(stepInst.step)-1:] == "Z"
}

type StepAndInst struct {
	step string
	inst Instruction
}

func alreadyDone(stepInsts []StepAndInst, stepInst StepAndInst) bool {
	for _, si := range stepInsts {
		if si == stepInst {
			return true
		}
	}
	return false
}

func getStepsToZ(route []StepAndInst) int {
	for i := len(route) - 1; i >= 0; i-- {
		step := route[i].step
		if step[len(step)-1:] == "Z" {
			return i
		}
	}
	panic("Couldn't find Z")
}

func (Solver) Solve(input string) string {
	maps := parseMaps(input)
	steps := maps.StartingPoints
	log.Printf("%-v", maps.StartingPoints)
	routes := make([][]StepAndInst, len(maps.StartingPoints))
	loop := make([]bool, len(maps.StartingPoints))
	zMet := make([]bool, len(maps.StartingPoints))
	for i := 0; true; i++ {
		inst := maps.Instructions[i%len(maps.Instructions)]
		for idx, step := range steps {
			stepInst := StepAndInst{step: step, inst: inst}
			if !loop[idx] {
				if zMet[idx] && alreadyDone(routes[idx], stepInst) {
					loop[idx] = true
					continue
				}
				s, z := doStep(stepInst, maps)
				if z {
					zMet[idx] = true
				}
				routes[idx] = append(routes[idx], stepInst)
				steps[idx] = s
			}
		}
		shouldEnd := true
		for _, e := range loop {
			if !e {
				shouldEnd = false
				break
			}
		}
		if shouldEnd {
			break
		}
	}
	stepsToZ := make([]int, len(maps.StartingPoints))
	for i, route := range routes {
		stepsToZ[i] = getStepsToZ(route)
		log.Printf("Steps to Z %d: %d", i, stepsToZ[i])
	}
	return strconv.Itoa(utils.LCM(stepsToZ))
}

//
//    The sandstorm is upon you and you aren't any closer to escaping the
//    wasteland. You had the camel follow the instructions, but you've barely
//    left your starting position. It's going to take significantly more
//    steps to escape!
//
//    What if the map isn't for people - what if the map is for ghosts? Are
//    ghosts even bound by the laws of spacetime? Only one way to find out.
//
//    After examining the maps a bit longer, your attention is drawn to a
//    curious fact: the number of nodes with names ending in A is equal to
//    the number ending in Z! If you were a ghost, you'd probably just start
//    at every node that ends with A and follow all of the paths at the same
//    time until they all simultaneously end up at nodes that end with Z.
//
//    For example:
// LR
//
// 11A = (11B, XXX)
// 11B = (XXX, 11Z)
// 11Z = (11B, XXX)
// 22A = (22B, XXX)
// 22B = (22C, 22C)
// 22C = (22Z, 22Z)
// 22Z = (22B, 22B)
// XXX = (XXX, XXX)
//
//    Here, there are two starting nodes, 11A and 22A (because they both end
//    with A). As you follow each left/right instruction, use that
//    instruction to simultaneously navigate away from both nodes you're
//    currently on. Repeat this process until all of the nodes you're
//    currently on end with Z. (If only some of the nodes you're on end with
//    Z, they act like any other node and you continue as normal.) In this
//    example, you would proceed as follows:
//      * Step 0: You are at 11A and 22A.
//      * Step 1: You choose all of the left paths, leading you to 11B and
//        22B.
//      * Step 2: You choose all of the right paths, leading you to 11Z and
//        22C.
//      * Step 3: You choose all of the left paths, leading you to 11B and
//        22Z.
//      * Step 4: You choose all of the right paths, leading you to 11Z and
//        22B.
//      * Step 5: You choose all of the left paths, leading you to 11B and
//        22C.
//      * Step 6: You choose all of the right paths, leading you to 11Z and
//        22Z.
//
//    So, in this example, you end up entirely on nodes that end in Z after 6
//    steps.
//
//    Simultaneously start on every node that ends with A. How many steps
//    does it take before you're only on nodes that end with Z?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
