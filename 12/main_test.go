package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		name string
		input []string
		want []*moon
	} {
		{"simple", []string{"<x=8, y=0, z=8>"}, []*moon{{8,0,8}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := parse(tt.input), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("parse() = %v, want %v", got, want)
			}
		})
	}
}

