package day10part1

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type Pipe string

const (
	Vertical      Pipe = "|"
	Horizontal    Pipe = "-"
	NorthEastBend Pipe = "L"
	NorthWestBend Pipe = "J"
	SouthWestBend Pipe = "7"
	SouthEastBend Pipe = "F"
	NoPipe        Pipe = "."
	Start         Pipe = "S"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func PossibleDirections(current Pipe) []Direction {
	switch current {
	case NoPipe:
		return []Direction{}
	case Start:
		return []Direction{North, South, East, West}
	case Vertical:
		return []Direction{North, South}
	case Horizontal:
		return []Direction{East, West}
	case NorthEastBend:
		return []Direction{North, East}
	case NorthWestBend:
		return []Direction{North, West}
	case SouthEastBend:
		return []Direction{South, East}
	case SouthWestBend:
		return []Direction{South, West}
	}
	panic(fmt.Sprintf("Unknown pipe: %-v", current))
}

type Coordinates struct {
	X int
	Y int
}

func ParseMap(input string) [][]Pipe {
	lines := utils.GetLines(input)
	var m [][]Pipe
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		m = append(m, []Pipe{})
		for _, pipe := range strings.Split(line, "") {
			m[len(m)-1] = append(m[len(m)-1], Pipe(pipe))
		}
	}
	return m
}

func GetStartingPoint(m [][]Pipe) Coordinates {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == Start {
				return Coordinates{Y: y, X: x}
			}
		}
	}
	panic("Couldn't find Start")
}

func getAccessibleNeighbors(m [][]Pipe, coords Coordinates) []Coordinates {
	var neighbors []Coordinates
	currentPipe := m[coords.Y][coords.X]
	for _, dir := range PossibleDirections(currentPipe) {
		switch dir {
		case North:
			if coords.Y > 0 {
				next := m[coords.Y-1][coords.X]
				if next == Vertical || next == SouthEastBend || next == SouthWestBend {
					neighbors = append(neighbors, Coordinates{Y: coords.Y - 1, X: coords.X})
				}
			}
		case South:
			if coords.Y < len(m)-1 {
				next := m[coords.Y+1][coords.X]
				if next == Vertical || next == NorthEastBend || next == NorthWestBend {
					neighbors = append(neighbors, Coordinates{Y: coords.Y + 1, X: coords.X})
				}
			}
		case East:
			if coords.X < len(m[coords.Y])-1 {
				next := m[coords.Y][coords.X+1]
				if next == Horizontal || next == NorthWestBend || next == SouthWestBend {
					neighbors = append(neighbors, Coordinates{Y: coords.Y, X: coords.X + 1})
				}
			}
		case West:
			if coords.X > 0 {
				next := m[coords.Y][coords.X-1]
				if next == Horizontal || next == NorthEastBend || next == SouthEastBend {
					neighbors = append(neighbors, Coordinates{Y: coords.Y, X: coords.X - 1})
				}
			}
		}
	}
	return neighbors
}

type Distance struct {
	Coords   Coordinates
	Distance int
}

func GetDistances(m [][]Pipe, startingPoint Coordinates) []Distance {
	var alreadyVisited []Distance
	coords := []Distance{{Distance: 0, Coords: startingPoint}}
	for len(coords) > 0 {
		current := coords[0]
		neighbors := getAccessibleNeighbors(m, current.Coords)
		for _, n := range neighbors {
			skip := false
			for _, av := range alreadyVisited {
				if av.Coords == n {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			coords = append(coords, Distance{Distance: current.Distance + 1, Coords: n})
		}
		alreadyVisited = append(alreadyVisited, current)
		coords = coords[1:]
	}
	return alreadyVisited

}

func (Solver) Solve(input string) string {
	m := ParseMap(input)
	start := GetStartingPoint(m)
	distances := GetDistances(m, start)
	maxDistance := 0
	for _, dist := range distances {
		if dist.Distance > maxDistance {
			maxDistance = dist.Distance
		}
	}
	return strconv.Itoa(maxDistance)
}

//
//    You use the hang glider to ride the hot air from Desert Island all the
//    way up to the floating metal island. This island is surprisingly cold
//    and there definitely aren't any thermals to glide on, so you leave your
//    hang glider behind.
//
//    You wander around for a while, but you don't find any people or
//    animals. However, you do occasionally find signposts labeled "Hot
//    Springs" pointing in a seemingly consistent direction; maybe you can
//    find someone at the hot springs and ask them where the desert-machine
//    parts are made.
//
//    The landscape here is alien; even the flowers and trees are made of
//    metal. As you stop to admire some metal grass, you notice something
//    metallic scurry away in your peripheral vision and jump into a big
//    pipe! It didn't look like any animal you've ever seen; if you want a
//    better look, you'll need to get ahead of it.
//
//    Scanning the area, you discover that the entire field you're standing
//    on is densely packed with pipes; it was hard to tell at first because
//    they're the same metallic silver color as the "ground". You make a
//    quick sketch of all of the surface pipes you can see (your puzzle
//    input).
//
//    The pipes are arranged in a two-dimensional grid of tiles:
//      * | is a vertical pipe connecting north and south.
//      * - is a horizontal pipe connecting east and west.
//      * L is a 90-degree bend connecting north and east.
//      * J is a 90-degree bend connecting north and west.
//      * 7 is a 90-degree bend connecting south and west.
//      * F is a 90-degree bend connecting south and east.
//      * . is ground; there is no pipe in this tile.
//      * S is the starting position of the animal; there is a pipe on this
//        tile, but your sketch doesn't show what shape the pipe has.
//
//    Based on the acoustics of the animal's scurrying, you're confident the
//    pipe that contains the animal is one large, continuous loop.
//
//    For example, here is a square loop of pipe:
// .....
// .F-7.
// .|.|.
// .L-J.
// .....
//
//    If the animal had entered this loop in the northwest corner, the sketch
//    would instead look like this:
// .....
// .S-7.
// .|.|.
// .L-J.
// .....
//
//    In the above diagram, the S tile is still a 90-degree F bend: you can
//    tell because of how the adjacent pipes connect to it.
//
//    Unfortunately, there are also many pipes that aren't connected to the
//    loop! This sketch shows the same loop as above:
// -L|F7
// 7S-7|
// L|7||
// -L-J|
// L|-JF
//
//    In the above diagram, you can still figure out which pipes form the
//    main loop: they're the ones connected to S, pipes those pipes connect
//    to, pipes those pipes connect to, and so on. Every pipe in the main
//    loop connects to its two neighbors (including S, which will have
//    exactly two pipes connecting to it, and which is assumed to connect
//    back to those two pipes).
//
//    Here is a sketch that contains a slightly more complex main loop:
// ..F7.
// .FJ|.
// SJ.L7
// |F--J
// LJ...
//
//    Here's the same example sketch with the extra, non-main-loop pipe tiles
//    also shown:
// 7-F7-
// .FJ|7
// SJLL7
// |F--J
// LJ.LJ
//
//    If you want to get out ahead of the animal, you should find the tile in
//    the loop that is farthest from the starting position. Because the
//    animal is in the pipe, it doesn't make sense to measure this by direct
//    distance. Instead, you need to find the tile that would take the
//    longest number of steps along the loop to reach from the starting point
//    - regardless of which way around the loop the animal went.
//
//    In the first example with the square loop:
// .....
// .S-7.
// .|.|.
// .L-J.
// .....
//
//    You can count the distance each tile in the loop is from the starting
//    point like this:
// .....
// .012.
// .1.3.
// .234.
// .....
//
//    In this example, the farthest point from the start is 4 steps away.
//
//    Here's the more complex loop again:
// ..F7.
// .FJ|.
// SJ.L7
// |F--J
// LJ...
//
//    Here are the distances for each tile on that loop:
// ..45.
// .236.
// 01.78
// 14567
// 23...
//
//    Find the single giant loop starting at S. How many steps along the loop
//    does it take to get from the starting position to the point farthest
//    from the starting position?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
