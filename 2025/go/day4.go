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
	part1, grid := countReplaced(grid)

	part2 := part1
	temp := 0
	for {
		grid = replaceAccessible(grid)
		temp, grid = countReplaced(grid)
		if (temp == 0) {
			break
		}
		part2 += temp
	}

	fmt.Printf("Part One: %d\n", part1)
	fmt.Printf("Part Two: %d\n", part2)
}

func countReplaced(grid [][]string) (int, [][]string) {
	sum := 0
	for line := 0; line < len(grid); line++ {
		for row := 0; row < len(grid[line]); row++ {
			if (grid[line][row] == "x") {
				grid[line][row] = "."
				sum++
			}
		}
	}
	return sum, grid
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
	if (north(line, row, grid)) {
		adjacent++
	}
	if (south(line, row, grid)) {
		adjacent++
	}
	if (west(line, row, grid)) {
		adjacent++
	}
	if (east(line, row, grid)) {
		adjacent++
	}
	if (northwest(line, row, grid)) {
		adjacent++
	}
	if (northeast(line, row, grid)) {
		adjacent++
	}
	if (southwest(line, row, grid)) {
		adjacent++
	}
	if (southeast(line, row, grid)) {
		adjacent++
	}
	return adjacent < 4
}

func north(line int, row int, grid [][]string) bool {
	if (line > 0) {
		return string(grid[line - 1][row]) != "."
	}
	return false
}

func south(line int, row int, grid [][]string) bool {
	if (line < len(grid) - 1) {
		return string(grid[line + 1][row]) != "."
	}
	return false
}

func west(line int, row int, grid [][]string) bool {
	if (row > 0) {
		return string(grid[line][row - 1]) != "."
	}
	return false
}

func east(line int, row int, grid [][]string) bool {
	if (row < len(grid[line]) - 1) {
		return string(grid[line][row + 1]) != "."
	}
	return false
}

func northwest(line int, row int, grid [][]string) bool {
	if (line > 0 && row > 0) {
		return string(grid[line - 1][row - 1]) != "."
	}
	return false
}

func northeast(line int, row int, grid [][]string) bool {
	if (line > 0 && row < len(grid[line]) - 1) {
		return string(grid[line - 1][row + 1]) != "."
	}
	return false
}

func southwest(line int, row int, grid [][]string) bool {
	if (line < len(grid) - 1 && row > 0) {
		return string(grid[line + 1][row - 1]) != "."
	}
	return false
}

func southeast(line int, row int, grid [][]string) bool {
	if (line < len(grid) - 1 && row < len(grid[line]) - 1) {
		return string(grid[line + 1][row + 1]) != "."
	}
	return false
}
