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
			if got := calculateNumOrbits(tt.orbitMap); got != tt.want {
				t.Errorf("calculateNumOrbits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOrbitTransfers(t *testing.T) {
	tests := []struct {
		name string
		orbitMap []string
		want int
	}{
		{"same orbit", []string{"COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L","I)YOU","I)SAN"}, 0},
		{"1 step", []string{"COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L","D)YOU","I)SAN"}, 1},
		{"simple", []string{"COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L","K)YOU","I)SAN"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOrbitTransfers(tt.orbitMap); got != tt.want {
				t.Errorf("minOrbitTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
