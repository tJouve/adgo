package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 3
	aoc.RegisterPartYear(year, day, 1, Part1)
	aoc.RegisterPartYear(year, day, 2, Part2)
}

func searchMax(input string) int {
	bank := strings.TrimSpace(input)
	max := 0
	for f := 0; f < len(bank); f++ {
		ten, _ := strconv.Atoi(string(bank[f]))
		ten *= 10
		for i := f + 1; i < len(bank); i++ {
			current, _ := strconv.Atoi(string(bank[i]))
			cInt := ten + current
			if cInt > max {
				max = cInt
			}
		}
	}
	return max
}

func searchMaxX(input string, batLen int) int {
	bank := strings.TrimSpace(input)
	start := 0
	var found []string
	//Create batLen batt bank
	for i := 0; i < batLen; i++ {
		//remaining cells to find
		rem := batLen - (i + 1)
		//rem := batLen - (len(found) + 1)
		limit := len(bank) - rem
		//Test current index
		index := start
		//Index Value
		d := string(bank[start])
		dv, _ := strconv.Atoi(d)
		// Find Bigger number than me after
		for j := start + 1; j < limit; j++ {
			jv, _ := strconv.Atoi(string(bank[j]))
			if jv > dv {
				//Yes
				//Index Value
				d = string(bank[j])
				dv, _ = strconv.Atoi(d)
				//Index is the new Start
				index = j
				if dv == 9 { //Quit if 9, can't be bigger
					break
				}
			}
		}
		found = append(found, d)
		start = index + 1
	}
	_max, _ := strconv.Atoi(strings.Join(found, ""))
	fmt.Fprintf(os.Stdout, "max for %s is %d\n", bank, _max)
	return _max
}

// Part1:
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	total := 0
	for _, r := range lines {
		total += searchMax(r)
	}

	return fmt.Sprintf("Part 1 : %d", total), nil
}

// Part2: count clicks
func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	total := 0
	for _, r := range lines {
		total += searchMaxX(r, 12)
	}

	return fmt.Sprintf("Part 2 : %d", total), nil
}
