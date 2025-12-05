package _range

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func Parse(text string) *Range {
	split := strings.Split(text, `-`)
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return &Range{Start: start, End: end}
}

func New(start, end int) *Range {
	return &Range{Start: start, End: end}
}

func (d *Range) Twice() []int {
	result := make([]int, 0)
	for i := d.Start; i <= d.End; i++ {
		t := strconv.Itoa(i)
		tLen := len(t)
		if tLen%2 == 0 {
			middleIndex := (tLen / 2)
			if t[:middleIndex] == t[middleIndex:] {
				result = append(result, i)
			}
		}
	}
	return result
}

func (d *Range) OverlapLenInclusive(o *Range) int {
	_len := 0
	if d.Start <= o.End && o.Start <= d.End {
		_len = min(o.End-d.Start, d.End-o.Start) + 1
		fmt.Printf("Overlap: %s & %s => %d\n", d, o, _len)
	}
	return _len
}

func (d *Range) LenInclusive() int {
	return d.End - d.Start + 1
}

func (d *Range) LenExclusive() int {
	return d.End - d.Start - 1
}

func (d *Range) Extend() []int {
	result := make([]int, 0)
	for i := d.Start; i <= d.End; i++ {
		result = append(result, i)
	}
	return result
}

func (d *Range) IsInInclusive(value int) bool {
	return value >= d.Start && value <= d.End
}

func (d *Range) IsInExclusive(value int) bool {
	return value > d.Start && value < d.End
}

func (d *Range) Duplicate() []int {
	result := make([]int, 0)

	for i := d.Start; i <= d.End; i++ {
		t := strconv.Itoa(i)
		tLen := len(t)
		alreadySearched := make([]string, 0)
		alreadyBad := make([]int, 0)
		if tLen > 1 {
			maxDupSize := tLen / 2
			for size := 1; size <= maxDupSize; size++ {
				for sindex := 0; sindex <= tLen-size; sindex++ {
					searched := t[sindex : sindex+size]
					if !slices.Contains(alreadySearched, searched) && !slices.Contains(alreadyBad, i) {
						alreadySearched = append(alreadySearched, searched)
						occ := strings.Count(t, searched)
						if occ >= tLen/size {
							if size*occ == tLen {
								_, err := fmt.Fprintf(os.Stdout, "%s is duplicate in %s\n", searched, t)
								if err != nil {
									return result
								}
								result = append(result, i)
								alreadyBad = append(alreadyBad, i)
							}
						}
					}
				}
			}
		}
	}
	return result
}
func (d *Range) String() string {
	return fmt.Sprintf("%d-%d", d.Start, d.End)
}
