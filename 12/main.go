package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {

}

func parse(s []string) []*moon {
	result := make([]*moon, len(s))
	for i, ss := range s {
		result[i] = parseLine(ss)
	}

	return result
}

func parseLine(s string) *moon {
	// <x=8, y=0, z=8>
	s2 := strings.Trim(s, "<>")
	ss := strings.Split(s2, ",")

	xIdx := strings.Index(ss[0], "x=")
	xStr := ss[0][xIdx+2:]
	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Fatal(err)
	}

	yIdx := strings.Index(ss[1], "y=")
	yStr := ss[1][yIdx+2:]
	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Fatal(err)
	}

	zIdx := strings.Index(ss[2], "z=")
	zStr := ss[2][zIdx+2:]
	z, err := strconv.Atoi(zStr)
	if err != nil {
		log.Fatal(err)
	}

	return &moon{x, y, z}
}

type moon struct {
	x, y, z int
}

func run() {

}
