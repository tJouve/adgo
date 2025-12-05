package day05

import (
	"testing"

	"adgo/internal/aoc"
)

func TestPartsExample(t *testing.T) {
	aoc.RunPartsTest(t, "data/sample2.txt", nil, Part2, "", "15")
}
