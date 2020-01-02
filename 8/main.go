package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("/Users/ggerritsen/dev/personal/advent-of-code-2019/8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start part 1")
	img := parseImg(string(b), 25, 6)
	l := findLayerWithFewestZeros(img)
	log.Printf("Found: %d", l.countOf(1) * l.countOf(2))
	log.Printf("End part 1")

	log.Printf("Start part 2")
	renderedImg := renderImg(string(b), 25, 6)
	for _, row := range renderedImg {
		s := ""
		for _, cell := range row {
			x := " "
			if cell == 1 {
				x = "*"
			}
			s = fmt.Sprintf("%s%s", s, x)
		}
		log.Printf("%s", s)
	}
	log.Printf("End part 2")
}

func renderImg(imgSrc string, width, height int) [][]int {
	src := parseImg(imgSrc, width, height)
	result := make([][]int, height)
	for i:=0; i<height; i++ {
		result[i] = make([]int, width)
	}

	// prefill with -1
	for _, row := range result {
		for y, _ := range row {
			row[y] = -1
		}
	}

	for i:=0; i<len(src);  i++ {
		for x, row := range src[i] {
			for y, v := range row {
				if v == 2 {
					continue
				}
				cell := result[x][y]
				if cell == 0 || cell == 1 {
					continue
				}
				result[x][y] = v
			}
		}
	}

	return result
}

func findLayerWithFewestZeros(img []layer) layer {
	minZeros := 99999999
	var result layer
	for _, l := range img {
		numZeros := l.countOf(0)
		if numZeros < minZeros {
			minZeros = numZeros
			result = l
		}
	}
	return result
}

func parseImg(imgSrc string, width, height int) []layer {
	layerSize := width * height
	numLayers := len(imgSrc) / layerSize
	result := make([]layer, numLayers)
	src := strings.Split(imgSrc, "")

	for i := 0; i < numLayers; i++ {
		l := make([][]int, height)
		for y := 0; y < height; y++ {
			l[y] = make([]int, width)
			for x := 0; x < width; x++ {
				s := src[x + y*width + i*layerSize]
				v, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				l[y][x] = v
			}
		}
		result[i] = l
	}

	return result
}

type layer [][]int

func (l layer) countOf(i int) int {
	numOfI := 0
	for _, row := range l {
		for _, cell := range row {
			if cell == i {
				numOfI++
			}
		}
	}
	return numOfI
}
