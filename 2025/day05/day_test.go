package day05

import (
	"testing"

	"adgo/internal/aoc"
)

func TestPartsExample(t *testing.T) {
	input, _ := aoc.ReadInput("data/sample2.txt")
	bf, _ := Part2Brute(input)
	aoc.RunPartsTest(t, "data/sample2.txt", nil, Part2, "", bf)
}
