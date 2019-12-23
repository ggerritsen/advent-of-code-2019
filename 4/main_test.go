package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_convertToIntSlice(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []int
	}{
		{"simple", 223450, []int{2, 2, 3, 4, 5, 0}},
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
		name  string
		input []int
		want  bool
	}{
		{"simple", []int{2, 2, 3, 4, 5, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsAdjacentEquals(tt.input); got != tt.want {
				t.Errorf("containsAdjacentEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsAdjacentEqualsAndNotMore(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{[]int{1, 1, 2, 2, 3, 3}, true},
		{[]int{2, 2, 3, 4, 5, 0}, true},
		{[]int{1, 1, 1, 1, 2, 2}, true},
		{[]int{1, 2, 3, 4, 4, 4}, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := containsAdjacentEqualsAndNotMore(tt.input); got != tt.want {
				t.Errorf("containsAdjacentEqualsAndNotMore() = %v, want %v", got, tt.want)
			}
		})
	}
}
