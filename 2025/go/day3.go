package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumPart1 := 0
	sumPart2 := 0

	for scanner.Scan() {
		bank := scanner.Text()
		if bank == "" {
			continue
		}

		val1, _ := strconv.Atoi(findNHighestBatteries(bank, 2))
		val2, _ := strconv.Atoi(findNHighestBatteries(bank, 12))

		sumPart1 += val1
		sumPart2 += val2
	}
	fmt.Printf("Part One: %d\n", sumPart1)
	fmt.Printf("Part Two: %d\n", sumPart2)
}

func findNHighestBatteries(bank string, n int) string {
	val := ""
	topIndex := -1 
	for i := 1; i <= n; i++ {
		topIndex = findIndexHighest(bank[topIndex + 1:len(bank) - (n - i)], topIndex + 1)
		top := string(bank[topIndex])
		val += top
	}
	return val
}

func findIndexHighest(bank string, start int) int {
	top := 0
	topIndex := 0
	for i := 0; i < len(bank); i++ {
		val, _ := strconv.Atoi(string(bank[i]))
		if (val > top) {
			top = val
			topIndex = i + start
		}
	}

	return topIndex
}
