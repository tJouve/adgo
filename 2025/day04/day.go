package day05

import (
	Grid "adgo/2025/day04/grid"
	"fmt"
	"os"
	"strings"

	_ "adgo/2025/day04/grid"
	"adgo/internal/aoc"
)

func init() {
	year := 2025
	day := 4
	aoc.RegisterPartYear(year, day, 1, Part1)
	aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1:
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	total := 0
	g := Grid.Parse(lines)
	amm := g.H * g.W
	for i := 0; i < amm; i++ {
		if g.Cell[i] == 0 {
			continue
		}
		var nb = 0
		val := g.NeighborValue(i)
		for _, v := range val {
			nb += v
		}
		if nb < 4 {
			total++
		}
	}
	return fmt.Sprintf("Part 1 : %d", total), nil
}

// Part2: count clicks
func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	total := 0
	g := Grid.Parse(lines)
	amm := g.H * g.W

	it := 0
	for {
		it++
		torm := make([]int, 0)
		for i := 0; i < amm; i++ {
			if g.Cell[i] == 0 {
				continue
			}
			var nb = 0
			val := g.NeighborValue(i)
			for _, v := range val {
				nb += v
			}
			if nb < 4 {
				total++
				//save roll to remove index
				torm = append(torm, i)
			}
		}
		//remove all marked
		for _, v := range torm {
			g.Cell[v] = 0
		}
		fmt.Fprintf(os.Stdout, "Turn %d => removing %d\n", it, len(torm))
		// if none to remove end of the game
		if len(torm) == 0 {
			break
		}
	}

	return fmt.Sprintf("Part 2 : %d", total), nil
}
