package main

import (
	"reflect"
	"testing"
)

func Test_parseImg(t *testing.T) {
	tests := []struct {
		name string
		imgSrc string
		width, height int
		want []layer
	}{
		{"simple", "123456789012", 3, 2, []layer{{[]int{1,2,3}, []int{4,5,6}}, {[]int{7,8,9}, []int{0,1,2}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseImg(tt.imgSrc, tt.width, tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseImg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renderImg(t *testing.T) {
	tests := []struct {
		name string
		imgSrc string
		width, height int
		want [][]int
	}{
		{"simple", "0222112222120000", 2, 2, [][]int{{0,1},{1,0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderImg(tt.imgSrc, tt.width, tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderImg() = %v, want %v", got, tt.want)
			}
		})
	}
}
