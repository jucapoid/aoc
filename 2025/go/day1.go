package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs/day1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentPos := 50
	part1Count := 0
	part2Count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing distance: %v\n", err)
			continue
		}

		if direction == 'R' {
			for i := 1; i <= distance; i++ {
				currentPos = (currentPos + 1) % 100
				if currentPos == 0 {
					part2Count++
				}
			}
		} else if direction == 'L' {
			for i := 1; i <= distance; i++ {
				currentPos = (currentPos - 1 + 100) % 100
				if currentPos == 0 {
					part2Count++
				}
			}
		}

		if currentPos == 0 {
			part1Count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

	fmt.Printf("Part One: %d\n", part1Count)
	fmt.Printf("Part Two: %d\n", part2Count)
}
