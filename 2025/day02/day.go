package day01

import (
	"adgo/commons/range"
	"fmt"
	"strings"

	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 2
	aoc.RegisterPartYear(year, day, 1, Part1)
	aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1:
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	ranges := strings.Split(lines[0], ",")
	total := 0
	for _, r := range ranges {
		currentRange := _range.Parse(r)
		dup := currentRange.Twice()
		for _, v := range dup {
			total += v
		}
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
