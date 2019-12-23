package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
		wire1 string
		wire2 string
		expected int
	}{
		{ "simple","R8,U5,L5,D3", "U7,R6,D4,L4", 6 },
		{ "extended1","R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159 },
		{ "extended2","R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135 },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, want := run(test.wire1, test.wire2), test.expected; got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}