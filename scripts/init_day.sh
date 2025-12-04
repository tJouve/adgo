#!/usr/bin/env bash
set -euo pipefail

# Usage: ./scripts/init_day.sh [YEAR] [DAY]
# If YEAR omitted, choose the largest 4-digit directory in the repo root (if any) or 2025.
# If DAY omitted, choose the next available dayNN in the YEAR directory.

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
YEAR_ARG="${1:-}"
DAY_ARG="${2:-}"

# Find default year
if [ -z "$YEAR_ARG" ]; then
  # find directories named with 4 digits, choose max
  found_years=()
  while IFS= read -r -d $'\0' d; do
    base=$(basename "$d")
    if [[ $base =~ ^[0-9]{4}$ ]]; then
      found_years+=("$base")
    fi
  done < <(find "$ROOT_DIR" -maxdepth 1 -type d -print0)

  if [ ${#found_years[@]} -gt 0 ]; then
    max=0
    for y in "${found_years[@]}"; do
      (( y > max )) && max=$y
    done
    YEAR=$max
  else
    YEAR=2025
  fi
else
  YEAR=$YEAR_ARG
fi

# Compute day
if [ -n "$DAY_ARG" ]; then
  DAY=$DAY_ARG
else
  # inspect existing dayNN directories in YEAR
  YEAR_DIR="$ROOT_DIR/$YEAR"
  max=0
  if [ -d "$YEAR_DIR" ]; then
    for d in "$YEAR_DIR"/day*; do
      [ -d "$d" ] || continue
      base=$(basename "$d")
      num=${base#day}
      if [[ "$num" =~ ^[0-9]+$ ]]; then
        (( num > max )) && max=$num
      fi
    done
  fi
  DAY=$((max + 1))
fi

DAYP=$(printf "%02d" "$DAY")
PKG="day${DAYP}"
DIR="$ROOT_DIR/$YEAR/$PKG"

if [ -e "$DIR" ]; then
  echo "Error: directory '$DIR' already exists" >&2
  exit 1
fi

mkdir -p "$DIR/data"

# create day.go
cat > "$DIR/day.go" <<EOF
package ${PKG}

import (
	"fmt"
	"strings"

	"adgo/internal/aoc"
)

func init() {
	year := ${YEAR}
	day := ${DAY}
	// register parts for this day
	aoc.RegisterPartYear(year, day, 1, Part1)
	//aoc.RegisterPartYear(year, day, 2, Part2)
}

// Part1: stub
func Part1(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	_ = lines
	return fmt.Sprintf("Part 1 : %d", 0), nil
}

// Part2: stub
func Part2(input string) (string, error) {
	lines := aoc.Lines(strings.TrimSpace(input))
	_ = lines
	return fmt.Sprintf("Part 2 : %d", 0), nil
}
EOF

# create test file with the shared TestPartsExample
cat > "$DIR/day_test.go" <<EOF
package ${PKG}

import (
	"testing"

	"adgo/internal/aoc"
)

func TestPartsExample(t *testing.T) {
	aoc.RunPartsTest(t, "data/sample.txt", Part1, Part2, "6", "3")
}
EOF

# create data files
: > "$DIR/data/sample.txt"
: > "$DIR/data/myinput.txt"

# add import into cmd/aoc/main.go if missing
MAIN="$ROOT_DIR/cmd/aoc/main.go"
IMPORT_LINE="_ \"adgo/${YEAR}/${PKG}\""

if ! grep -qF "$IMPORT_LINE" "$MAIN"; then
  # insert into the import block
  awk -v import_line="$IMPORT_LINE" '
    /^import \($/ { in_import=1; print; next }
    in_import && /^$/ { print import_line; in_import=0 }
    { print }
  ' "$MAIN" > "$MAIN.tmp" && mv "$MAIN.tmp" "$MAIN"
  # format the main.go file
  gofmt -w "$MAIN"
fi
cat <<MSG
Created: $DIR
 - $DIR/day.go
 - $DIR/day_test.go
 - $DIR/data/sample.txt
 - $DIR/data/myinput.txt

Edited : $MAIN (to add import)

MSG

