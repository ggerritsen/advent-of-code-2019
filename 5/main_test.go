package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"simple", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"extended", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{"extended 2", []int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}},
		{"negative", []int{1101, 100, -1, 4, 0}, []int{1101, 100, -1, 4, 99}},
		{"jump-if-true", []int{1105, 55, 7, 1101, 0, 0, 8, 99, 1}, []int{1105, 55, 7, 1101, 0, 0, 8, 99, 1}},
		{"jump-if-true2", []int{1105, 0, 7, 1101, 0, 0, 8, 99, 1}, []int{1105, 0, 7, 1101, 0, 0, 8, 99, 0}},
		{"jump-if-false", []int{6,10,13,1,11,12,11,4,11,99,55,0,1,7}, []int{6,10,13,1,11,12,11,4,11,99,55,1,1,7}},
		{"jump-if-false2", []int{6,10,13,1,11,12,11,4,11,99,0,0,1,7}, []int{6,10,13,1,11,12,11,4,11,99,0,0,1,7}},
		{"less-than", []int{4,0,7,9,10,9,4,9,99,55,8}, []int{4,0,7,9,10,9,4,9,99,0,8}},
		{"less-than-2", []int{4,0,7,9,10,9,4,9,99,0,8}, []int{4,0,7,9,10,9,4,9,99,1,8}},
		{"less-than-3", []int{4,0,1107,55,8,3,4,3,99}, []int{4,0,1107,0,8,3,4,3,99}},
		{"less-than-4", []int{4,0,1107,0,8,3,4,3,99}, []int{4,0,1107,1,8,3,4,3,99}},
		{"equals", []int{4,0,8,9,10,9,4,9,99,55,8}, []int{4,0,8,9,10,9,4,9,99,0,8}},
		{"equals-2", []int{4,0,8,9,10,9,4,9,99,8,8}, []int{4,0,8,9,10,9,4,9,99,1,8}},
		{"equals-3", []int{4,0,1108,55,8,3,4,3,99}, []int{4,0,1108,0,8,3,4,3,99}},
		{"equals-4", []int{4,0,1108,8,8,3,4,3,99}, []int{4,0,1108,1,8,3,4,3,99}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, want := run(test.input), test.output; !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func Test_parseOperand(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
		want1 []int
	}{
		{"simple", 1002, 2, []int{0, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOperator(tt.input)
			if got != tt.want {
				t.Errorf("parseOperator() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseOperator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
