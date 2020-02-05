package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		name string
		input []string
		want []moon
	} {
		{"simple", []string{"<x=8, y=0, z=8>"}, []moon{{pos: position{8,0,8}}}},
		{"extended", []string{"<x=8, y=0, z=8>", "<x=0, y=-5, z=-10>"}, []moon{{pos: position{8,0,8}}, {pos:position{0, -5, -10}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := parse(tt.input), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("parse() = %v, want %v", got, want)
			}
		})
	}
}

func Test_iteration(t *testing.T) {
	tests := []struct {
		name string
		input []moon
		want []moon
	} {
		{"simple", []moon{
			newMoon(-1, 0, 2, 0, 0, 0),
			newMoon(2, -10, -7, 0, 0, 0),
			newMoon(4, -8, 8, 0, 0, 0),
			newMoon(3, 5, -1, 0, 0, 0),
		}, []moon{
			newMoon(2, -1, 1, 3,-1, -1),
			newMoon(3, -7, -4, 1,3, 3),
			newMoon(1, -7, 5, -3,1, -3),
			newMoon(2, 2, 0, -1,-3, 1),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := iterate(tt.input), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("iterate() = %v, want %v", got, want)
			}
		})
	}
}


func Test_calcEnergy(t *testing.T) {
	tests := []struct {
		name string
		input []moon
		want int
	} {
		{"simple", []moon{
			newMoon(2, -1, 1, 3,-1, -1),
			newMoon(3, -7, -4, 1,3, 3),
			newMoon(1, -7, 5, -3,1, -3),
			newMoon(2, 2, 0, -1,-3, 1),
		}, 20+98+91+20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := calcEnergy(tt.input), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("calcEnergy() = %v, want %v", got, want)
			}
		})
	}
}

