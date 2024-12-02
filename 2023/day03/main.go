package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X     int
	Y     int
	Steps int
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}

func manhattenDistance(a, b Coordinate) int {
	return abs(b.X-a.Y) + abs(b.Y-a.Y)
}

func visitedBy(start Coordinate, movement []string) []Coordinate {
	var visited []Coordinate
	pos := start
	for _, cmd := range movement {
		direction := cmd[0]
		distance, _ := strconv.Atoi(cmd[1:])
		switch direction {
		case 'U':
			for i := 0; i < distance; i++ {
				pos.Y += 1
				pos.Steps += 1
				visited = append(visited, pos)
			}
		case 'D':
			for i := 0; i < distance; i++ {
				pos.Y -= 1
				pos.Steps += 1
				visited = append(visited, pos)
			}
		case 'L':
			for i := 0; i < distance; i++ {
				pos.X -= 1
				pos.Steps += 1
				visited = append(visited, pos)
			}
		case 'R':
			for i := 0; i < distance; i++ {
				pos.X += 1
				pos.Steps += 1
				visited = append(visited, pos)
			}
		}

	}

	return visited
}

func same(a, b Coordinate) bool {
	return a.X == b.X && a.Y == b.Y
}

func findDuplicates(left, right []Coordinate) []Coordinate {
	var dups []Coordinate
	for _, a := range left {
		for _, b := range right {
			if same(a, b) {
				dups = append(dups, a)
			}
		}
	}
	return dups
}

func findDuplicates2(left, right []Coordinate) []int {
	var dups []int
	for _, a := range left {
		for _, b := range right {
			if same(a, b) {
				dups = append(dups, a.Steps+b.Steps)
			}
		}
	}
	return dups
}

func calculateDistance(wire1 []string, wire2 []string) int {
	var occupied1, occupied2 []Coordinate
	start := Coordinate{X: 0, Y: 0}

	// occupied by wire1
	occupied1 = visitedBy(start, wire1)

	// occupied by wire2
	occupied2 = visitedBy(start, wire2)

	// find duplicates
	crosses := findDuplicates(occupied1, occupied2)

	// distance for all duplicates
	var distances []int
	for _, x := range crosses {
		distances = append(distances, manhattenDistance(start, x))
	}

	// return smallest
	sort.Ints(distances)
	return distances[0]
}

func calculateSteps(wire1 []string, wire2 []string) int {
	var occupied1, occupied2 []Coordinate
	start := Coordinate{X: 0, Y: 0}

	// occupied by wire1
	occupied1 = visitedBy(start, wire1)

	// occupied by wire2
	occupied2 = visitedBy(start, wire2)

	// find stepCount for all duplicates
	crosses := findDuplicates2(occupied1, occupied2)

	// return smallest
	sort.Ints(crosses)
	return crosses[0]
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	wires := strings.Split(string(file), "\n")[0:2]
	wire1 := strings.Split(wires[0], ",")
	wire2 := strings.Split(wires[1], ",")

	distance := calculateDistance(wire1, wire2)
	fmt.Printf("Solution to part 1 is %d\n", distance)

	steps := calculateSteps(wire1, wire2)
	fmt.Printf("Solution to part 2 is %d\n", steps)

}
