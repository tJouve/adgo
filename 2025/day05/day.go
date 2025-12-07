package day05

import (
	"fmt"
	"slices"
	"sort"
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
	return strconv.Itoa(len(all)), nil
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
	sort.SliceStable(interval, func(i, j int) bool {
		d := interval[i]
		o := interval[j]
		if d.Start < o.Start {
			return true
		}
		if d.Start > o.Start {
			return false
		}
		if d.End < o.End {
			return true
		}
		if d.End > o.End {
			return false
		}
		return false
	})

	total := 0
	for i, r := range interval {
		_len := r.LenInclusive()
		total += _len
		var overlaps []int
		for _, other := range interval[i+1:] {
			overlap := r.OverlapLenInclusive(other)
			overlaps = append(overlaps, overlap)

		}
		slices.Sort(overlaps)
		if len(overlaps) > 0 {
			total -= overlaps[len(overlaps)-1]
		}

		fmt.Printf("Current %s Total: %d\n", r, total)
	}
	return fmt.Sprintf("%d", total), nil
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
