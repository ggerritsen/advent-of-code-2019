package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		//{"simple", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		//{"extended", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{"input", []int{3, 0, 4, 0, 99}, []int{3, 0, 4, 0, 99}},
		//{"extended 2", []int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, want := run(test.input), test.output; !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
