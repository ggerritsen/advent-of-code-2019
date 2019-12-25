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
		{"negative", []int{1101,100,-1,4,0}, []int{1101,100,-1,4,99}},
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
		{"simple", 1002, 2, []int{0,1,0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOperand(tt.input)
			if got != tt.want {
				t.Errorf("parseOperand() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseOperand() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
