Advent of Code - Go template

Usage:

Run a day:

    make run DAY=1 PART=0

Flags:
- DAY: day number (0 = last registered day)
- PART: 0=both,1=2
- INPUT: optional path to input file; defaults to days/dayNN/data/myinput.txt if present

Conventions for day data files:
- Each day uses a `data/` directory.
- `data/sample.txt` contains the example used by unit tests (used by `go test`).
- `data/myinput.txt` contains your personal input; it's used by the runner when no `--input` is provided.

Add a new day:
- create `days/dayNN` with `dayNN.go` exposing `Part1(input string)(string,error)` and `Part2(input string)(string,error)` (nil is allowed for an unimplemented part).
- call `aoc.Register(N, Part1, Part2)` in `init()`.
- add `data/sample.txt` and `data/myinput.txt`.
- add tests in `dayNN_test.go` that read `data/sample.txt`.
