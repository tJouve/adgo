package day01

import (
	"testing"

	"AD2025/internal/aoc"
)

func TestPartsExample(t *testing.T) {
	aoc.RunPartsTest(t, "data/sample.txt", Part1, Part2, "6", "3")
}
