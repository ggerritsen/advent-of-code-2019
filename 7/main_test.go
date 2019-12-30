package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name   string
		i      []int
		input  []int
		result []int
		output int
	}{
		{name: "simple", i: []int{1, 0, 0, 0, 99}, result: []int{2, 0, 0, 0, 99}},
		{name: "extended", i: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, result: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{name: "extended 2", i: []int{1002, 4, 3, 4, 33}, result: []int{1002, 4, 3, 4, 99}},
		{name: "negative", i: []int{1101, 100, -1, 4, 0}, result: []int{1101, 100, -1, 4, 99}},
		{name: "jump-if-true", i: []int{1105, 55, 7, 1101, 0, 0, 8, 99, 1}, result: []int{1105, 55, 7, 1101, 0, 0, 8, 99, 1}},
		{name: "jump-if-true2", i: []int{1105, 0, 7, 1101, 0, 0, 8, 99, 1}, result: []int{1105, 0, 7, 1101, 0, 0, 8, 99, 0}},
		{name: "jump-if-false", i: []int{6, 8, 11, 1, 9, 10, 9, 99, 55, 0, 1, 7}, result: []int{6, 8, 11, 1, 9, 10, 9, 99, 55, 1, 1, 7}},
		{name: "jump-if-false2", i: []int{6, 8, 11, 1, 9, 10, 9, 99, 0, 0, 1, 7}, result: []int{6, 8, 11, 1, 9, 10, 9, 99, 0, 0, 1, 7}},
		{name: "less-than", i: []int{7, 5, 6, 5, 99, 55, 8}, result: []int{7, 5, 6, 5, 99, 0, 8}},
		{name: "less-than-2", i: []int{7, 5, 6, 5, 99, 0, 8}, result: []int{7, 5, 6, 5, 99, 1, 8}},
		{name: "less-than-3", i: []int{1107, 55, 8, 1, 99}, result: []int{1107, 0, 8, 1, 99}},
		{name: "less-than-4", i: []int{1107, 0, 8, 1, 99}, result: []int{1107, 1, 8, 1, 99}},
		{name: "equals", i: []int{8, 5, 6, 5, 99, 55, 8}, result: []int{8, 5, 6, 5, 99, 0, 8}},
		{name: "equals-2", i: []int{8, 5, 6, 5, 99, 8, 8}, result: []int{8, 5, 6, 5, 99, 1, 8}},
		{name: "equals-3", i: []int{1108, 55, 8, 1, 99}, result: []int{1108, 0, 8, 1, 99}},
		{name: "equals-4", i: []int{1108, 8, 8, 1, 99}, result: []int{1108, 1, 8, 1, 99}},
		{name: "i/o", i: []int{3, 0, 4, 0, 99}, input: []int{55}, result: []int{55, 0, 4, 0, 99}, output: 55},
		{name: "i/o multiple inputs", i: []int{3, 0, 3, 3, 4, 3, 99}, input: []int{1, 55}, result: []int{1, 0, 3, 55, 4, 3, 99}, output: 55},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got1, got2 := runIntcode(test.i, test.input)
			want1, want2 := test.result, test.output
			if !reflect.DeepEqual(got1, want1) {
				t.Errorf("got %v want %v", got1, want1)
			}
			if got2 != want2 {
				t.Errorf("got %v want %v", got2, want2)
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

func Test_runAmplifiers(t *testing.T) {
	type args struct {
		i      []int
		phases []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{4, 3, 2, 1, 0}}, 43210},
		{"simple-2", args{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, []int{0, 1, 2, 3, 4}}, 54321},
		{"simple-3", args{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []int{1, 0, 4, 3, 2}}, 65210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runAmplifiers(tt.args.i, tt.args.phases); got != tt.want {
				t.Errorf("runAmplifiers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOptimalPhases(t *testing.T) {
	tests := []struct {
		name     string
		intCodes []int
		want     int
	}{
		{"simple", []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
		{"simple-2", []int{3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0}, 54321},
		{"simple-3", []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33, 1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}, 65210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOptimalOutput(tt.intCodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOptimalOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePermutations(t *testing.T) {
	tests := []struct {
		name string
		k    int
		a    []int
		want [][]int
	}{
		{"simple", 1, []int{1}, [][]int{{1}}},
		{"simple-2", 2, []int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{"simple-3", 3, []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePermutations(tt.k, tt.a)
			if len(got) != len(tt.want) {
				t.Errorf("len = %d, want %d", len(got), len(tt.want))
			}
		outer:
			for _, v := range tt.want {
				for _, v2 := range got {
					if reflect.DeepEqual(v, v2) {
						continue outer
					}
				}
				t.Errorf("did not find %v, (got %v)", v, got)
			}
		})
	}
}
