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

func sumUpMul(line string) int {
	i := 0
	sum := 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	for {
		line = line[i:]
		match := re.FindStringIndex(line)
		if match == nil {
			break
		}
		instruction := line[match[0]:match[1]]
		sum += calculate(instruction)
		i = match[1]
	}
	return sum
}

func cleanUp(line string) string {
	var clean string
	for {
		re_do := regexp.MustCompile(`do\(\)`)
		re_dont := regexp.MustCompile(`don\'t\(\)`)
		dont := re_dont.FindStringIndex(line[0:])
		if dont == nil {
			return line
		}
		fmt.Printf("DONT: %d,%d\n", dont[0], dont[1])
		do := re_do.FindStringIndex(line[dont[1]:])
		if do == nil {
			return line[0:dont[0]]
		}
		fmt.Printf("  DO: %d,%d\n", do[0], do[1])
		clean = line[0:dont[0]] + "----" + line[dont[1]+do[1]:len(line)]
		//clean = line[0:dont[0]] + "----" + line[dont[1]+do[1]:len(line)]
		return cleanUp(clean)
	}

}

func main() {
	file, _ := os.Open("input.txt")
	//file, _ := os.Open("test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1 := 0
	part2 := 0
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}
	part1 += sumUpMul(text)
	clean := cleanUp(text)
	fmt.Println(clean)
	part2 += sumUpMul(clean)

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
