package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	_ "adgo/2025/day01"
	_ "adgo/2025/day02"
	"adgo/internal/aoc"
)

func main() {
	year := flag.Int("year", 2025, "year to run (0=last)")
	day := flag.Int("day", 0, "day to run (0=last)")
	part := flag.Int("part", -1, "part to run (-1=last,0=both,1,2)")
	test := flag.Bool("test", true, "use sample input if true, real input if false. This option is override by -input.")
	inputPath := flag.String("input", "", "path to input file (defaults to days/<year>/dayNN/data/myinput.txt or stdin)")
	flag.Parse()

	yearVal := *year
	if yearVal == 0 {
		lastYear := aoc.LastYear()
		if lastYear == 0 {
			// No positive year registered. If there are entries for year 0 (compat), accept year 0.
			if aoc.HasYear(0) {
				yearVal = 0
			} else {
				fmt.Fprintln(os.Stderr, "no years registered")
				os.Exit(2)
			}
		} else {
			yearVal = lastYear
		}
	}

	dayVal := *day
	if dayVal == 0 {
		last := aoc.LastDay(yearVal)
		if last == 0 {
			fmt.Fprintf(os.Stderr, "no days registered for year %d\n", yearVal)
			os.Exit(2)
		}
		dayVal = last
	}

	// If user didn't specify a part (part == -1), pick the last registered part for this year/day.
	partVal := *part
	if partVal == -1 {
		lastPart := aoc.LastPart(yearVal, dayVal)
		if lastPart == 0 {
			// no parts registered? default to part 1 behavior
			partVal = 1
		} else {
			partVal = lastPart
		}
	}

	p1fn, p2fn, ok := aoc.GetPartsForYear(yearVal, dayVal)
	if !ok {
		fmt.Fprintf(os.Stderr, "no solver registered for year %d day %d\n", yearVal, dayVal)
		os.Exit(2)
	}

	input := *inputPath
	if input == "" {
		if *test {
			// try sample path with year: days/<year>/dayNN/data/sample.txt
			input = filepath.Join(fmt.Sprintf("%d", yearVal), fmt.Sprintf("day%02d", dayVal), "data", "sample.txt")
		} else {
			input = filepath.Join(fmt.Sprintf("%d", yearVal), fmt.Sprintf("day%02d", dayVal), "data", "myinput.txt")
		}
	}
	if _, err := os.Stat(input); err == nil {
		fmt.Printf("Trying sample path: %s\n", input)
	} else {
		fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
		os.Exit(2)
	}
	inStr, err := aoc.ReadInput(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
		os.Exit(2)
	}

	// Execute requested parts independently. Part 2 can run without part 1.
	switch partVal {
	case 0:
		if p1fn != nil {
			p1, err := p1fn(inStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "part1 failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Part 1: %s\n", p1)
		} else {
			fmt.Println("Part 1: <not implemented>")
		}
		if p2fn != nil {
			p2, err := p2fn(inStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "part2 failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Part 2: %s\n", p2)
		} else {
			fmt.Println("Part 2: <not implemented>")
		}
	case 1:
		if p1fn == nil {
			fmt.Fprintln(os.Stderr, "part 1 not implemented for this day")
			os.Exit(2)
		}
		p1, err := p1fn(inStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "part1 failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(p1)
	case 2:
		if p2fn == nil {
			fmt.Fprintln(os.Stderr, "part 2 not implemented for this day")
			os.Exit(2)
		}
		p2, err := p2fn(inStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "part2 failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(p2)
	default:
		fmt.Fprintln(os.Stderr, "invalid part, must be 0,1,2 or -1 for last")
		os.Exit(2)
	}
}
