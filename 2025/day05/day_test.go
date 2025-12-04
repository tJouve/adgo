package day05

import (
	"testing"

	"adgo/internal/aoc"
)

func TestPartsExample(t *testing.T) {
	aoc.RunPartsTest(t, "data/sample.txt", Part1, Part2, "6", "3")
}
