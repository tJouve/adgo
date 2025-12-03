package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"adgo/2025/day02/range"
	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 3
	aoc.RegisterPartYear(year, day, 1, Part1)
	//aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1:
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	total := 0
	for _, r := range lines {
		firstValue := 0
		//secondIndex := 0
		intArray := make([]int, 0)

		for i := 0; i < len(r); i++ {
			intvalue, _ := strconv.Atoi(string(r[i]))
			intArray = append(intArray, intvalue)
		}
		sortArray := make([]int, len(intArray))
		copy(sortArray, intArray)
		sort.Ints(sortArray[:])
		firstValue = sortArray[0]

	}

	return fmt.Sprintf("Part 1 : %d", total), nil
}

// Part2: count clicks
func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	ranges := strings.Split(lines[0], ",")
	total := 0
	for _, r := range ranges {
		currentRange := _range.Parse(r)
		dup := currentRange.Duplicate()
		for _, v := range dup {
			total += v
		}
	}
	return fmt.Sprintf("Part 2 : %d", total), nil
}
