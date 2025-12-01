package aoc

import (
	"strconv"
	"strings"
)

// Lines splits input into lines, trimming a trailing newline.
func Lines(s string) []string {
	if s == "" {
		return nil
	}
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimRight(s, "\n")
	return strings.Split(s, "\n")
}

// Ints converts a slice of strings to ints; returns error on any parse failure.
func Ints(ss []string) ([]int, error) {
	out := make([]int, 0, len(ss))
	for _, line := range ss {
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

// SplitOnBlank splits lines into groups separated by blank lines.
func SplitOnBlank(lines []string) [][]string {
	var groups [][]string
	var cur []string
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			if len(cur) > 0 {
				groups = append(groups, cur)
				cur = nil
			}
			continue
		}
		cur = append(cur, l)
	}
	if len(cur) > 0 {
		groups = append(groups, cur)
	}
	return groups
}
