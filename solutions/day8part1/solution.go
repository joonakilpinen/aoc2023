package day8part1

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
	StartingPoint string
	EndingPoint   string
	Instructions  []Instruction
	Nodes         map[string]Node
}

var nodeRegex = regexp.MustCompile("^(?P<Id>[A-Z]{3}) = \\((?P<Left>[A-Z]{3}), (?P<Right>[A-Z]{3})\\)")

func parseMaps(input string) Maps {
	lines := utils.GetLines(input)
	var instructions []Instruction
	for _, inst := range strings.Split(lines[0], "") {
		instructions = append(instructions, Instruction(inst))
	}
	nodes := map[string]Node{}
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		matches := nodeRegex.FindStringSubmatch(lines[i])
		id := matches[nodeRegex.SubexpIndex("Id")]
		left := matches[nodeRegex.SubexpIndex("Left")]
		right := matches[nodeRegex.SubexpIndex("Right")]
		nodes[id] = Node{Left: left, Right: right}
	}
	return Maps{
		StartingPoint: "AAA",
		EndingPoint:   "ZZZ",
		Instructions:  instructions,
		Nodes:         nodes,
	}
}

func (Solver) Solve(input string) string {
	maps := parseMaps(input)
	step := maps.StartingPoint
	stepsNeeded := 0
	for i := 0; true; i++ {
		if step == maps.EndingPoint {
			break
		}
		node := maps.Nodes[step]
		inst := maps.Instructions[i%len(maps.Instructions)]
		switch inst {
		case L:
			step = node.Left
		case R:
			step = node.Right
		}
		stepsNeeded += 1
	}
	log.Printf("Steps: %d", stepsNeeded)
	return strconv.Itoa(stepsNeeded)
}

//
//    You're still riding a camel across Desert Island when you spot a
//    sandstorm quickly approaching. When you turn to warn the Elf, she
//    disappears before your eyes! To be fair, she had just finished warning
//    you about ghosts a few minutes ago.
//
//    One of the camel's pouches is labeled "maps" - sure enough, it's full
//    of documents (your puzzle input) about how to navigate the desert. At
//    least, you're pretty sure that's what they are; one of the documents
//    contains a list of left/right instructions, and the rest of the
//    documents seem to describe some kind of network of labeled nodes.
//
//    It seems like you're meant to use the left/right instructions to
//    navigate the network. Perhaps if you have the camel follow the same
//    instructions, you can escape the haunted wasteland!
//
//    After examining the maps for a bit, two nodes stick out: AAA and ZZZ.
//    You feel like AAA is where you are now, and you have to follow the
//    left/right instructions until you reach ZZZ.
//
//    This format defines each node of the network individually. For example:
// RL
//
// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
//
//    Starting with AAA, you need to look up the next element based on the
//    next left/right instruction in your input. In this example, start with
//    AAA and go right (R) by choosing the right element of AAA, CCC. Then, L
//    means to choose the left element of CCC, ZZZ. By following the
//    left/right instructions, you reach ZZZ in 2 steps.
//
//    Of course, you might not find ZZZ right away. If you run out of
//    left/right instructions, repeat the whole sequence of instructions as
//    necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example,
//    here is a situation that takes 6 steps to reach ZZZ:
// LLR
//
// AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)
//
//    Starting at AAA, follow the left/right instructions. How many steps are
//    required to reach ZZZ?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
