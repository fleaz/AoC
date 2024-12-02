package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	fuel := math.Floor(float64(mass)/3) - 2
	return int(fuel)
}

func complexCalculateFuel(mass int) int {
	totalFuel := 0
	for {
		f := calculateFuel(mass)
		if f > 0 {
			totalFuel += f
			mass = f
		} else {
			return totalFuel
		}

	}
}

func main() {
	totalFuel := 0
	totalComplexFuel := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		totalFuel += calculateFuel(m)
		totalComplexFuel += complexCalculateFuel(m)
	}

	fmt.Printf("Solution for part 1 is %d\n", totalFuel)
	fmt.Printf("Solution for part 2 is %d\n", totalComplexFuel)

}
