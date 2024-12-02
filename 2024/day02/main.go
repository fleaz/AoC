package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
func isSafe(list []int) bool {
	inc := list[0] < list[1]

	for i := 0; i < len(list)-1; i++ {
		d := diff(list[i], list[i+1])

		if inc != (list[i] < list[i+1]) ||
			(d < 1) || (d > 3) {
			return false
		}
	}
	return true
}

func ToIntSlice(s []string) []int {
	var x []int
	for _, e := range s {
		i, _ := strconv.Atoi(e)
		x = append(x, i)
	}
	return x
}

func main() {
	file, _ := os.Open("input.txt")
	//file, _ := os.Open("test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if isSafe(ToIntSlice(line)) {
			fmt.Println("Safe")
			count += 1
		} else {
			fmt.Println("Unsafe")
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}
