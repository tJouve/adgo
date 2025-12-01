package day01

import (
	"fmt"
	"strings"

	"AD2025/internal/aoc"
)

func init() {
	// register each part separately (compat: year 0)
	aoc.RegisterPart(1, 1, Part1)
	aoc.RegisterPart(1, 2, Part2)
	// also register under year 2026
	aoc.RegisterPartYear(2026, 1, 1, Part1)
	aoc.RegisterPartYear(2026, 1, 2, Part2)
}

// Part1: sum of ints
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	ints, err := aoc.Ints(lines)
	if err != nil {
		return "", fmt.Errorf("parse ints: %w", err)
	}
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return fmt.Sprintf("%d", sum), nil
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
