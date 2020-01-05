package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		wantLocation     point
		wantNumAsteroids int
	}{
		{"1-dimension-x", `#.#.#`, point{2, 0}, 2},
		{"1-dimension-y",
			`#
			 .
			#
			.
			#`, point{0, 2}, 2},
		{"2-dimension",
			`..#..
				#.#.#
				..#..`, point{2, 1}, 4},
		{"simple",
			`.#..#
			.....
			#####
			....#
			...##`, point{3, 4}, 8},
		{"extended-1",
			`......#.#.
			#..#.#....
			..#######.
			.#.#.###..
			.#..#.....
			..#....#.#
			#..#....#.
			.##.#..###
			##...#..#.
			.#....####`, point{5, 8}, 33},
		{"extended-2",
			`#.#...#.#.
			.###....#.
			.#....#...
			##.#.#.#.#
			....#.#.#.
			.##..###.#
			..#...##..
			..##....##
			......#...
			.####.###.`, point{1, 2}, 35},
		{"extended-3",
			`.#..#..###
			####.###.#
			....###.#.
			..###.##.#
			##.##.#.#.
			....###..#
			..#.#..#.#
			#..#.#.###
			.##...##.#
			.....#.#..`, point{6, 3}, 41},
		{"extended-4",
			`.#..##.###...#######
			##.############..##.
			.#.######.########.#
			.###.#######.####.#.
			#####.##.#.##.###.##
			..#####..#.#########
			####################
			#.####....###.#.#.##
			##.#################
			#####.##.###..####..
			..######..##.#######
			####.##.####...##..#
			.#####..#.######.###
			##...#.##########...
			#.##########.#######
			.####.#.###.###.#.##
			....##.##.###..#####
			.#.#.###########.###
			#.#.#.#####.####.###
			###.##.####.##.#..##`, point{11, 13}, 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocation, gotNumAsteroids := findBestMonitoringLocation(tt.input)
			if !reflect.DeepEqual(gotLocation, tt.wantLocation) {
				t.Errorf("findBestMonitoringLocation() location = %v, want %v", gotLocation, tt.wantLocation)
			}
			if gotNumAsteroids != tt.wantNumAsteroids {
				t.Errorf("findBestMonitoringLocation() numAsteroids = %d, want %d", gotNumAsteroids, tt.wantNumAsteroids)
			}
		})
	}
}
