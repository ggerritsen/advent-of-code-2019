package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/1/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		input io.Reader
		expected int
	}{
		{"simple", strings.NewReader("10"), 1},
		{"input.txt", f, 4822435},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, want := part2(test.input), test.expected; got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

