package day05

import (
	"fmt"
	"strings"

	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 5
	// register parts for this day
	aoc.RegisterPartYear(year, day, 1, Part1)
	//aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1: stub
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	_ = lines
	return fmt.Sprintf("Part 1 : %d", 0), nil
}

// Part2: stub
func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	_ = lines
	return fmt.Sprintf("Part 2 : %d", 0), nil
}
