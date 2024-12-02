package main

import (
	"fmt"
	"strconv"
)

func checkRealPair(pos int, value byte, pwString string) bool {
	var left, right byte

	if pos+2 == len(pwString) {
		// we are at the right bound of pwString
		left = pwString[pos-1]
		right = 'x'
	} else if pos == 0 {
		// we are at the left bound of pwString
		left = 'x'
		right = pwString[pos+2]
	} else {
		// we are in the middle
		left = pwString[pos-1]
		right = pwString[pos+2]
	}

	if (left == value) || (right == value) {
		return false
	}

	return true
}

func isValid(password int, extendet bool) bool {
	pwString := strconv.Itoa(password)

	//Length criteria
	if len(pwString) != 6 {
		return false
	}

	// adjavent digits criteria
	pair := 0
	realPair := 0
	for i := 0; i < len(pwString)-1; i++ {
		left := pwString[i]
		right := pwString[i+1]
		if left == right {
			pair += 1
			if checkRealPair(i, left, pwString) {
				realPair += 1
			}
		}
	}
	if pair == 0 {
		return false
	}

	// incresing numbers criteria
	for i := 0; i < len(pwString)-1; i++ {
		left, _ := strconv.Atoi(string(pwString[i]))
		right, _ := strconv.Atoi(string(pwString[i+1]))
		if left > right {
			return false
		}
	}

	// check extra rule for part 2
	if extendet && realPair == 0 {
		return false
	}

	// All check passed
	return true
}

func findValidPasswords(from int, to int, extendet bool) int {
	counter := 0
	for i := from; i <= to; i++ {
		if isValid(i, extendet) {
			counter++
		}
	}

	return counter
}

func main() {
	from := 359282
	to := 820401

	count := findValidPasswords(from, to, false)
	fmt.Printf("Result for part 1 is %d\n", count)

	count2 := findValidPasswords(from, to, true)
	fmt.Printf("Result for part 2 is %d\n", count2)
}
