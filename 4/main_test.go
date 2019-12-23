package main

import (
	"reflect"
	"testing"
)

func Test_convertToIntSlice(t *testing.T) {
	tests := []struct {
		name string
		input int
		want []int
	}{
		{ "simple", 223450, []int{2, 2, 3, 4, 5, 0} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToIntSlice(tt.input, []int{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDecreasing(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{"simple", []int{2, 2, 3, 4, 5, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDecreasing(tt.input); got != tt.want {
				t.Errorf("isDecreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_containsAdjacentEquals(t *testing.T) {
	tests := []struct {
		name string
		input []int
		want bool
	}{
		{ "simple", []int{2, 2, 3, 4, 5, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsAdjacentEquals(tt.input); got != tt.want {
				t.Errorf("containsAdjacentEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}
