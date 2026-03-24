package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

func main() {
	data, err := ioutil.ReadFile("../inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(data))
	parts := strings.Split(s, ",")
	var ranges []Range

	for _, p := range parts {
		if p == "" {
			continue
		}

		bounds := strings.Split(p, "-")
		start, _ := strconv.ParseInt(bounds[0], 10, 64)
		end, _ := strconv.ParseInt(bounds[1], 10, 64)
		ranges = append(ranges, Range{start, end})
	}

	sumPart1 := int64(0)
	sumPart2 := int64(0)

	for _, r := range	ranges {
		for i := r.Start; i <= r.End; i++ {
			if (isInvalidPart1(i)) {
				sumPart1 += i
			}
			if (isInvalidPart2(i)) {
				sumPart2 += i
			}
		}
	}

	fmt.Printf("Part One: %d\n", sumPart1)
	fmt.Printf("Part Two: %d\n", sumPart2)
}

func isInvalidPart1(id int64) bool {
	str := strconv.FormatInt(id, 10)
	l := len(str)
	if (l % 2 == 0 && str[0:l/2] == str[l/2:]) {
		return true
	}
	return false
}

func isInvalidPart2(id int64) bool {
	str := strconv.FormatInt(id, 10)
	for j := len(str); j >= 2; j-- {
		if (len(str) % j == 0) {
			parts := divideString(str, j)
			if (allPartsEqual(parts)) {
				return true
			}
		}
	}
	return false
}

func divideString(mystr string, size int) []string {
   var parts []string
   partSize := len(mystr) / size
   for i := 0; i < size; i++ {
      start := i * partSize
      end := start + partSize
      if i == size-1 {
         end = len(mystr)
      }
      parts = append(parts, mystr[start:end])
   }
   return parts
}

func allPartsEqual(parts []string) bool {
	for i := 0; i < len(parts) - 1; i++ {
		for j := 1; j < len(parts); j++ {
			if (parts[i] != parts[j]) {
				return false
			}
		}
	}
	return true
}
