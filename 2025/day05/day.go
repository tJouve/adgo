package day05

import (
	"fmt"
	"strconv"
	"strings"

	"adgo/commons/range"
	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 5
	// register parts for this day
	aoc.RegisterPartYear(year, day, 1, Part1)
	aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1: stub
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	var interval []*_range.Range
	ligneVide := 0
	for i, line := range lines {
		if line == "" {
			ligneVide = i
			break
		}
		r := _range.Parse(line)
		interval = append(interval, r)
	}
	var ingredients []int
	for _, line := range lines[ligneVide+1:] {
		val, _ := strconv.Atoi(line)
		ingredients = append(ingredients, val)
	}
	total := 0
	for _, ing := range ingredients {
		for _, r := range interval {
			if r.IsInInclusive(ing) {
				total++
				break
			}
		}
	}

	return fmt.Sprintf("Part 1 : %d", total), nil
}

// Part2: stub
func Part2Brute(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	var interval []*_range.Range
	for _, line := range lines {
		if line == "" {
			break
		}
		r := _range.Parse(line)
		interval = append(interval, r)
	}

	var all []int
	for _, r := range interval {
		ex := r.Extend()
		all = append(all, ex...)
	}
	all = removeDuplicateInt(all)
	return fmt.Sprintf("Part 1 : %d", len(all)), nil
}

func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	var interval []*_range.Range
	for _, line := range lines {
		if line == "" {
			break
		}
		r := _range.Parse(line)
		interval = append(interval, r)
	}

	total := 0
	for i, r := range interval {
		_len := r.LenInclusive()
		total += _len
		for _, other := range interval[i+1:] {
			overlap := r.OverlapLenInclusive(other)
			total -= overlap
		}
	}
	return fmt.Sprintf("Part 2 : %d", total), nil
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
