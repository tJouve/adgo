package dial

// Dial represents a numeric dial with a current value and inclusive bounds Min...Max.
// It initializes Current to 50 by default (constructor ensures it's inside bounds).
// Rotations wrap around the inclusive range [Min, Max].

type Dial struct {
	Current int
	Min     int
	Max     int
	Click   int
}

// New creates a Dial with given min and max and Current initialized to 50.
// If min > max they are swapped. If 50 is outside the range Current is set to Min.
func New(min, max int) *Dial {
	if min > max {
		min, max = max, min
	}
	cur := 50
	if cur < min || cur > max {
		cur = min
	}
	return &Dial{Current: cur, Min: min, Max: max}
}

// RotateRight increases the dial by n steps, wrapping inside [Min,Max].
// It updates and returns the new Current value.
func (d *Dial) RotateRight(n int) int {
	return d.add(n)
}

// RotateLeft decreases the dial by n steps, wrapping inside [Min,Max].
// It updates and returns the new Current value.
func (d *Dial) RotateLeft(n int) int {
	return d.add(-n)
}

// add moves the current value by delta (maybe negative) and wraps inside [Min,Max].
func (d *Dial) add(delta int) int {
	// Ensure bounds are valid
	if d.Max < d.Min {
		d.Min, d.Max = d.Max, d.Min
	}
	// Range length (inclusive)
	r := d.Max - d.Min + 1
	if r <= 0 {
		// Degenerate: set to Min
		d.Current = d.Min
		return d.Current
	}
	newVal := d.Current + delta
	// Normalize into [0, r-1]
	offset := ((newVal-d.Min)%r + r) % r

	d.Current = d.Min + offset
	return d.Current
}
