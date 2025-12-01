package aoc

// PartFunc is the signature for a part solver.
type PartFunc func(string) (string, error)

// registry holds per-year maps: year -> day -> part -> PartFunc
var registry = map[int]map[int]map[int]PartFunc{}

// RegisterPartYear registers a function for a specific year, day and part.
func RegisterPartYear(year int, day int, part int, fn PartFunc) {
	ym, ok := registry[year]
	if !ok {
		ym = map[int]map[int]PartFunc{}
		registry[year] = ym
	}
	dm, ok := ym[day]
	if !ok {
		dm = map[int]PartFunc{}
		ym[day] = dm
	}
	dm[part] = fn
}

// RegisterPart keeps backwards compatibility: register under year 0.
func RegisterPart(day int, part int, fn PartFunc) {
	RegisterPartYear(0, day, part, fn)
}

// GetPart returns the function registered for year/day/part, and whether it exists.
func GetPart(year int, day int, part int) (PartFunc, bool) {
	ym, ok := registry[year]
	if !ok {
		return nil, false
	}
	dm, ok := ym[day]
	if !ok {
		return nil, false
	}
	fn, ok := dm[part]
	return fn, ok
}

// GetPartsForYear is a convenience: returns Part1, Part2, and whether any entry exists for the day in the year.
func GetPartsForYear(year int, day int) (PartFunc, PartFunc, bool) {
	ym, ok := registry[year]
	if !ok {
		return nil, nil, false
	}
	dm, ok := ym[day]
	if !ok {
		return nil, nil, false
	}
	p1, _ := dm[1]
	p2, _ := dm[2]
	return p1, p2, true
}

// GetParts keeps backwards compatibility: get parts for year 0.
func GetParts(day int) (PartFunc, PartFunc, bool) {
	return GetPartsForYear(0, day)
}

// LastYear returns the highest registered year number or 0 if none.
func LastYear() int {
	max := 0
	for y := range registry {
		if y > max {
			max = y
		}
	}
	return max
}

// HasYear reports whether a given year is present in the registry.
func HasYear(year int) bool {
	_, ok := registry[year]
	return ok
}

// HasAnyYears reports whether any year is registered (including year 0).
func HasAnyYears() bool {
	return len(registry) > 0
}

// LastDay returns the highest registered day for a given year, or 0 if none.
func LastDay(year int) int {
	ym, ok := registry[year]
	if !ok {
		return 0
	}
	max := 0
	for d := range ym {
		if d > max {
			max = d
		}
	}
	return max
}

// LastPart returns the highest registered part id for a given year and day, or 0 if none.
func LastPart(year int, day int) int {
	ym, ok := registry[year]
	if !ok {
		return 0
	}
	dm, ok := ym[day]
	if !ok {
		return 0
	}
	max := 0
	for p := range dm {
		if p > max {
			max = p
		}
	}
	return max
}
