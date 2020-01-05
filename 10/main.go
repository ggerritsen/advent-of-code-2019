package main

import (
	"log"
	"math"
	"strings"
)

func main() {

}

func run(s string) (point, int) {
	// parse s
	rows := strings.Split(s, "\n")
	coords := make([][]string, len(rows))
	for i, row := range rows {
		trimmed := strings.TrimSpace(row)
		coords[i] = make([]string, len(trimmed))
	}
	for i, row := range rows {
		trimmed := strings.TrimSpace(row)
		for j, c := range trimmed {
			coords[i][j] = string(c)
		}
	}
	coords = transpose(coords)
	log.Printf("coords: %v", coords)

	var asteroids []point
	for i, row := range coords {
		for j, v := range row {
			if v == "#" {
				asteroids = append(asteroids, point{i, j})
			}
		}
	}

	maxAsteroidsInSight := 0
	maxLocation := point{-1, -1}
	for _, origin := range asteroids {

		angles := map[float64]bool{}
		for _, a := range asteroids {
			if origin == a {
				continue
			}

			angles[origin.angle(a)] = true
		}

		log.Printf("location %v can see %d asteroids", origin, len(angles))

		if maxAsteroidsInSight < len(angles) {
			maxAsteroidsInSight = len(angles)
			maxLocation = point{origin.x, origin.y}
		}
	}

	log.Printf("maxAsteroidsInSight %d, maxLocation %v", maxAsteroidsInSight, maxLocation)

	return maxLocation, maxAsteroidsInSight
}

type point struct {
	x, y int
}

func (p point) angle(other point) float64 {
	deltaX, deltaY := p.x-other.x, p.y-other.y
	return math.Atan2(float64(deltaX), float64(deltaY))
}

func transpose(input [][]string) [][]string {
	// prepare output matrix
	output := make([][]string, len(input[0]))
	for i := 0; i < len(output); i++ {
		output[i] = make([]string, len(input))
	}

	// fill output matrix
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			output[j][i] = input[i][j]
		}
	}

	return output
}
