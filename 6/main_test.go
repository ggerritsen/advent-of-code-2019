package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		orbitMap []string
		want int
	}{
		{"direct orbit", []string{"COM)B"}, 1},
		{"indirect orbit", []string{"COM)B", "B)C"}, 3},
		{"extended", []string{"COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L"},42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.orbitMap); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
