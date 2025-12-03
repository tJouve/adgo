package _range

import "testing"

func TestDialRotate(t *testing.T) {
	d := New(0, 99) // range 0..99, start 50
	if d.Current != 50 {
		t.Fatalf("expected start 50, got %d", d.Current)
	}

	// rotate right by 10 -> 60
	if got := d.RotateRight(10); got != 60 {
		t.Fatalf("expected 60, got %d", got)
	}

	// rotate right by 50 -> should wrap: 60+50=110 -> wrap to 10
	if got := d.RotateRight(50); got != 10 {
		t.Fatalf("expected 10, got %d", got)
	}

	// rotate left by 20 -> 10-20=-10 -> wrap to 90
	if got := d.RotateLeft(20); got != 90 {
		t.Fatalf("expected 90, got %d", got)
	}

	// rotate left by 100 -> 90-100=-10 -> wraps to 90 again
	if got := d.RotateLeft(100); got != 90 {
		t.Fatalf("expected 90, got %d", got)
	}
}

func TestDialDegenerateRange(t *testing.T) {
	// if min>max, they should be swapped
	d := New(10, 5)
	if d.Min != 5 || d.Max != 10 {
		t.Fatalf("expected swapped bounds 5..10, got %d..%d", d.Min, d.Max)
	}

	// start current should be min because 50 is outside 5..10
	if d.Current != d.Min {
		t.Fatalf("expected start at min %d, got %d", d.Min, d.Current)
	}

	// add 3 -> min+3
	if got := d.RotateRight(3); got != d.Min+3 {
		t.Fatalf("expected %d, got %d", d.Min+3, got)
	}
}
