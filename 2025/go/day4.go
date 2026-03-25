package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

func main() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var grid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	grid = replaceAccessible(grid)
	part1 := countReplaced(grid)

	part2 := part1
	temp := 0
	for {
		if (part2 == temp) {
			break
		}
		part2 = temp
		grid = replaceAccessible(grid)
		temp = countReplaced(grid)
	}

	fmt.Printf("Part One: %d\n", part1)
	fmt.Printf("Part Two: %d\n", part2)
}

func countReplaced(grid [][]string) int {
	sum := 0
	for line := 0; line < len(grid); line++ {
		for row := 0; row < len(grid[line]); row++ {
			if (grid[line][row] == "x") {
				sum++
			}
		}
	}
	return sum
}

func replaceAccessible(grid [][]string) [][]string {
	for line := 0; line < len(grid); line++ {
		for row := 0; row < len(grid[line]); row++ {
			if (string(grid[line][row]) == "@" && canBeAccessed(line, row, grid)) {
				grid[line][row] = "x"
			}
		}
	}
	return grid
}

func canBeAccessed(line int, row int, grid [][]string) bool {
	adjacent := 0
	if (line > 0 && string(grid[line - 1][row]) == "@") {
		adjacent++
	}
	if (line < len(grid) - 1 && string(grid[line + 1][row]) == "@") {
		adjacent++
	}
	if (row > 0 && string(grid[line][row - 1]) == "@") {
		adjacent++
	}
	if (row < len(grid[line]) - 1 && string(grid[line][row + 1]) == "@") {
		adjacent++
	}
	if (line > 0 && row > 0 && string(grid[line - 1][row - 1]) == "@") {
		adjacent++
	}
	if (line > 0 && row < len(grid[line]) - 1 && string(grid[line - 1][row + 1]) == "@") {
		adjacent++
	}
	if (line < len(grid) - 1 && row > 0 && string(grid[line + 1][row - 1]) == "@") {
		adjacent++
	}
	if (line < len(grid) - 1 && row < len(grid[line]) - 1 && string(grid[line + 1][row + 1]) == "@") {
		adjacent++
	}
	return adjacent < 4
}

