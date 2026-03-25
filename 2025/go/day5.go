
package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
	"sort"
)

type Range struct {
	Start int64
	End   int64
}

func main() {
	file, err := os.Open("../inputs/day5.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var ranges []Range
	var ingredients []int64

	scanner := bufio.NewScanner(file)
	ingredientIds := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ingredientIds = true
			continue
		}

		if ! ingredientIds {
			s := strings.TrimSpace(string(line))
			bounds := strings.Split(s, "-")
			start, _ := strconv.ParseInt(bounds[0], 10, 64)
			end, _ := strconv.ParseInt(bounds[1], 10, 64)
			ranges = append(ranges, Range{start, end})
		} else {
			id, _ := strconv.ParseInt(string(line), 10, 64)
			ingredients = append(ingredients, id)
		}
	}

	ranges = removeOverlaps(ranges)

	fresh := 0
	for _, i := range ingredients {
		if isFresh(i, ranges) {
			fresh++
		}
	}

	numIngredients := int64(0)
	for _, r := range ranges {
		numIngredients += r.End - r.Start + 1
	}

	fmt.Printf("Part One: %d\n", fresh)
	fmt.Printf("Part Two: %d\n", numIngredients)
}

func isFresh(id int64, ranges []Range) bool {
	for _, r := range ranges {
		if r.Start <= id && r.End >= id {
			return true
		}
	}
	return false
}

func removeOverlaps(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}

	for _, r := range ranges[1:] {
		previous := &merged[len(merged) - 1]
	
		if r.Start <= previous.End {
			if r.End > previous.End {
				previous.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}
