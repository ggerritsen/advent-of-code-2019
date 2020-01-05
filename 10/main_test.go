package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		wantLocation     []int
		wantNumAsteroids int
	}{
		{"1-dimension-x", `#.#.#`, []int{2, 0}, 2},
		{"1-dimension-y",
			`#
			 .
			#
			.
			#`, []int{0, 2}, 2},
		{"2-dimension",
			`..#..
				#.#.#
				..#..`, []int{2, 1}, 4},
		{"simple",
			`.#..#
			.....
			#####
			....#
			...##`, []int{3, 4}, 8},
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
			.#....####`, []int{5, 8}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocation, gotNumAsteroids := run(tt.input)
			if !reflect.DeepEqual(gotLocation, tt.wantLocation) {
				t.Errorf("run() location = %v, want %v", gotLocation, tt.wantLocation)
			}
			if gotNumAsteroids != tt.wantNumAsteroids {
				t.Errorf("run() numAsteroids = %d, want %d", gotNumAsteroids, tt.wantNumAsteroids)
			}
		})
	}
}

func Test_firstBlockingAsteroid(t *testing.T) {
	type args struct {
		a      point
		origin point
		coords string
	}
	tests := []struct {
		name string
		args args
		want point
	}{{
		// 9,7 should be blocked by 3,3
		"edge case", args{point{9,7}, point{0,1}, `......#.#.
			#..#.#....
			..#######.
			.#.#.###..
			.#..#.....
			..#....#.#
			#..#....#.
			.##.#..###
			##...#..#.
			.#....####`}, point{3,3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBlockingAsteroid(tt.args.a, tt.args.origin, tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("firstBlockingAsteroid() = %v, want %v", got, tt.want)
			}
		})
	}
}
