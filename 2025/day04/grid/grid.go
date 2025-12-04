package Grid

import (
	"strings"
)

type Grid struct {
	W    int
	H    int
	Cell []int
}

func Parse(lines []string) *Grid {
	w := len(lines[0])
	h := len(lines)
	cells := make([]int, 0)
	for _, l := range lines {
		for _, c := range strings.Split(l, "") {
			if c == "." {
				cells = append(cells, 0)
			} else {
				cells = append(cells, 1)
			}
		}
	}

	return &Grid{W: w, H: h, Cell: cells}
}

func (d *Grid) OutOfRange(current int, target int) bool {
	//top and bottom
	if target < 0 || target >= len(d.Cell) {
		return true
	}
	//left and right
	ccol := current % d.W
	tcol := target % d.W
	//should be at most 1 column apart
	if ccol-tcol > 1 || ccol-tcol < -1 {
		return true
	}
	return false
}

func (d *Grid) NeighborValue(index int) []int {
	ns := make([]int, 0)
	var ni = []int{index - d.W - 1, index - d.W, index - d.W + 1,
		index - 1, index + 1,
		index + d.W - 1, index + d.W, index + d.W + 1}
	for _, v := range ni {
		if !d.OutOfRange(index, v) {
			ns = append(ns, d.Cell[v])
		}
	}

	return ns
}

func (d *Grid) NeighborIndex(index int) []int {
	ns := make([]int, 0)
	var ni = []int{index - d.W - 1, index - d.W, index - d.W + 1,
		index - 1, index + 1,
		index + d.W - 1, index + d.W, index + d.W + 1}
	for _, v := range ni {
		if !d.OutOfRange(index, v) {
			ns = append(ns, v)
		}
	}

	return ns
}
