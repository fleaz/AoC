package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func countIn(x int, list []int) int {
	cnt := 0
	for _, i := range list {
		if x == i {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	var left, right []int
	f, _ := os.Open("input.txt")
	//f, _ := os.Open("test.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "   ")
		a, _ := strconv.Atoi(x[0])
		b, _ := strconv.Atoi(x[1])
		left = append(left, a)
		right = append(right, b)
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0.0
	for i, _ := range left {
		distance += math.Abs(float64((right[i] - left[i])))
	}

	similarity := 0
	for i, _ := range left {
		similarity += left[i] * countIn(left[i], right)
	}

	fmt.Printf("Part 1: %d\n", int(distance))
	fmt.Printf("Part 2: %d\n", int(similarity))
}
