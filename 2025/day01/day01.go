package day01

import (
	"fmt"
	"strconv"
	"strings"

	"adgo/2025/day01/dial"
	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 1
	// also register under year 0
	aoc.RegisterPartYear(year, day, 1, Part1)
	//aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1:
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	d := dial.New(0, 99)
	result := 0
	for _, line := range lines {
		s := strings.TrimSpace(line)
		if s == "" {
			continue
		}
		// detect direction
		hasL := strings.ContainsAny(s, "Ll")
		hasR := strings.ContainsAny(s, "Rr")
		if !hasL && !hasR {
			// nothing to do
			continue
		}
		// remove L/R characters to get the numeric part (default to 1)
		clean := strings.Map(func(r rune) rune {
			if r == 'L' || r == 'R' || r == 'l' || r == 'r' {
				return -1
			}
			return r
		}, s)
		clean = strings.TrimSpace(clean)
		steps := 1
		if clean != "" {
			if n, err := strconv.Atoi(clean); err == nil {
				steps = n
			}
		}
		if hasL {
			d.RotateLeft(steps)
		} else {
			d.RotateRight(steps)
		}
		if d.Current == 0 {
			result++
		}
	}
	return fmt.Sprintf("%d", result), nil
}

// Part2: count of ints
func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	ints, err := aoc.Ints(lines)
	if err != nil {
		return "", fmt.Errorf("parse ints: %w", err)
	}
	return fmt.Sprintf("%d", len(ints)), nil
}
