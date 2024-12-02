package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func format(progString string) []int {
	var prog []int
	for _, s := range strings.Split(progString, ",") {
		cmd, _ := strconv.Atoi(s)
		prog = append(prog, cmd)
	}
	return prog
}

func execute(prog []int) ([]int, error) {
	pos := 0
	for {
		cmd := prog[pos]
		switch cmd {
		case 1:
			left := prog[pos+1]
			right := prog[pos+2]
			target := prog[pos+3]
			prog[target] = prog[left] + prog[right]
			pos += 4
		case 2:
			left := prog[pos+1]
			right := prog[pos+2]
			target := prog[pos+3]
			prog[target] = prog[left] * prog[right]
			pos += 4
		case 99:
			return prog, nil
		default:
			return nil, fmt.Errorf("Unknown command")
		}
	}

}

func solve(verb, noun int, prog []int) (int, error) {
	prog[1] = verb
	prog[2] = noun
	result, err := execute(prog)
	if err != nil {
		return 0, err
	}
	return result[0], nil
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var progString string

	for scanner.Scan() {
		progString = scanner.Text()
	}

	// Solve part 1
	result, _ := solve(12, 2, format(progString))
	fmt.Printf("Result for part 1 is %d\n", result)

	// Solve part 2
	for verb := 0; verb <= 99; verb++ {
		for noun := 0; noun <= 99; noun++ {
			prog := format(progString)
			prog[1] = verb
			prog[2] = noun
			//fmt.Printf("%d,%d,%d,%d\n", prog[0], prog[1], prog[2], prog[3])
			result, err := solve(verb, noun, prog)
			if err == nil && result == 19690720 {
				fmt.Printf("Result for part 2 is 100 * %d + %d = %d\n", verb, noun, 100*verb+noun)
				os.Exit(0)
			}
		}
	}

}
