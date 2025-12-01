package aoc

import "testing"

// RunPartsTest is a helper for day tests.
// - samplePath is a path relative to the day package (e.g. "data/sample.txt")
// - p1/p2 are the part functions (may be nil if not implemented)
// - want1/want2 are expected string outputs; when empty string means "no assertion".
func RunPartsTest(t *testing.T, samplePath string, p1 PartFunc, p2 PartFunc, want1 string, want2 string) {
	t.Helper()
	input, err := ReadInput(samplePath)
	if err != nil {
		t.Fatalf("read sample: %v", err)
	}
	if p1 != nil {
		got1, err := p1(input)
		if err != nil {
			t.Fatalf("part1: %v", err)
		}
		if want1 != "" && got1 != want1 {
			t.Fatalf("part1: got %s want %s", got1, want1)
		}
	} else {
		if want1 != "" {
			t.Fatalf("part1 not implemented but expected %s", want1)
		}
	}
	if p2 != nil {
		got2, err := p2(input)
		if err != nil {
			t.Fatalf("part2: %v", err)
		}
		if want2 != "" && got2 != want2 {
			t.Fatalf("part2: got %s want %s", got2, want2)
		}
	} else {
		if want2 != "" {
			t.Fatalf("part2 not implemented but expected %s", want2)
		}
	}
}
