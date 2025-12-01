package aoc

import (
	"io"
	"os"
)

// ReadInput returns the content of the file at path. If path is empty, reads from stdin.
func ReadInput(path string) (string, error) {
	var r io.Reader
	if path == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(path)
		if err != nil {
			return "", err
		}
		defer func() { _ = f.Close() }()
		r = f
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
