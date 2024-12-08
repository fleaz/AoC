package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculate(input string) int {
	x := strings.Split(input[4:len(input)-1], ",")
	a, _ := strconv.Atoi(x[0])
	b, _ := strconv.Atoi(x[1])
	return a * b
}

func main() {
	file, _ := os.Open("input.txt")
	//file, _ := os.Open("test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		for {
			fmt.Printf("i: %d\n", i)
			line = line[i:]
			match := re.FindStringIndex(line)
			if match == nil {
				break
			}
			// fmt.Printf("0: %d\n", match[0])
			// fmt.Printf("1: %d\n", match[1])
			instruction := line[match[0]:match[1]]
			sum += calculate(instruction)
			i = match[1]
			//bufio.NewScanner(os.Stdin).Scan()
		}
	}

	fmt.Printf("Part1: %d\n", sum)
}
