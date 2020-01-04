package main

import (
	"reflect"
	"testing"
)

func Test_testBoost(t *testing.T) {
	tests := []struct {
		name string
		opCodes []int
		want []int
	}{
		{"simple", []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}, []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}},
		{"simple-2", []int{1102,34915192,34915192,7,4,7,99,0}, []int{1219070632396864}},
		{"simple-3", []int{104,1125899906842624,99}, []int{1125899906842624}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runBoostWithInput(tt.opCodes, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runBoostWithInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
