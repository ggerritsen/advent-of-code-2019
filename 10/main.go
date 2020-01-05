package main

import (
	"log"
	"strings"
)

func main() {

}

func run(s string) ([]int, int) {
	// parse s into [][]string
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

	log.Printf("coords: %v", coords)
	coords = transpose(coords)
	log.Printf("coords transposed: %v", coords)

	width, height := len(coords), len(coords[0])
	log.Printf("w %d, h %d", width, height)

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

		asteroidsInSight := 0
	outer:
		for _, a := range asteroids {
			if origin == a {
				continue
			}

			// Q1 (between 9 and 12 o'clock)
			if a.x < origin.x && a.y < origin.y {
				deltaX := origin.x - a.x
				deltaY := origin.y - a.y

				// correct for 'smaller' distances
				if deltaX > 0 && deltaY > 0 {
					if deltaX%deltaY == 0 {
						deltaX = deltaX / deltaY
						deltaY = 1
					}
					if deltaY%deltaX == 0 {
						deltaY = deltaY / deltaX
						deltaX = 1
					}
				}
				if deltaX == 0 {
					deltaY = 1
				}
				if deltaY == 0 {
					deltaX = 1
				}

				for k, l := a.x+deltaX, a.y+deltaY; k < origin.x && l < origin.y; k, l = k+deltaX, l+deltaY {
					if coords[k][l] == "#" {
						log.Printf("Q1 Found blocking asteroid, %v not visible to %v (blocked by %v)", a, origin, point{k, l})
						// found blocking asteroid, continue
						continue outer
					}
				}

			}

			// Q2 (between 12 and 3 o'clock)
			if a.x >= origin.x && a.y < origin.y {
				deltaX := a.x - origin.x
				deltaY := origin.y - a.y

				// correct for 'smaller' distances
				if deltaX > 0 && deltaY > 0 {
					if deltaX%deltaY == 0 {
						deltaX = deltaX / deltaY
						deltaY = 1
					}
					if deltaY%deltaX == 0 {
						deltaY = deltaY / deltaX
						deltaX = 1
					}
				}
				if deltaX == 0 {
					deltaY = 1
				}
				if deltaY == 0 {
					deltaX = 1
				}

				for k, l := a.x-deltaX, a.y+deltaY; k >= origin.x && l < origin.y; k, l = k-deltaX, l+deltaY {
					if coords[k][l] == "#" {
						// found blocking asteroid, continue
						log.Printf("Q2 Found blocking asteroid, %v not visible to %v (blocked by %v)", a, origin, point{k, l})
						continue outer
					}
				}
			}

			// Q3 (between 3 and 6 o'clock)
			if a.x >= origin.x && a.y >= origin.y {
				deltaX := a.x - origin.x
				deltaY := a.y - origin.y

				// correct for 'smaller' distances
				if deltaX > 0 && deltaY > 0 {
					if deltaX%deltaY == 0 {
						deltaX = deltaX / deltaY
						deltaY = 1
					}
					if deltaY%deltaX == 0 {
						deltaY = deltaY / deltaX
						deltaX = 1
					}
				}
				if deltaX == 0 {
					deltaY = 1
				}
				if deltaY == 0 {
					deltaX = 1
				}
				// 9,7 should be blocked by 3,3

				for k, l := a.x-deltaX, a.y-deltaY; k >= origin.x && l >= origin.y; k, l = k-deltaX, l-deltaY {
					if coords[k][l] == "#" {
						if k == origin.x && l == origin.y {
							break
						}
						// found blocking asteroid, continue
						log.Printf("Q3 Found blocking asteroid, %v not visible to %v (blocked by %v)", a, origin, point{k, l})
						continue outer
					}
				}
			}

			// Q4 (between 6 and 9 o'clock)
			if a.x < origin.x && a.y >= origin.y {
				deltaX := origin.x - a.x
				deltaY := a.y - origin.y

				// correct for 'smaller' distances
				if deltaX > 0 && deltaY > 0 {
					if deltaX%deltaY == 0 {
						deltaX = deltaX / deltaY
						deltaY = 1
					}
					if deltaY%deltaX == 0 {
						deltaY = deltaY / deltaX
						deltaX = 1
					}
				}
				if deltaX == 0 {
					deltaY = 1
				}
				if deltaY == 0 {
					deltaX = 1
				}

				for k, l := a.x+deltaX, a.y-deltaY; k < origin.x && l >= origin.y; k, l = k+deltaX, l-deltaY {
					if coords[k][l] == "#" {
						// found blocking asteroid, continue
						log.Printf("Q4 Found blocking asteroid, %v not visible to %v (blocked by %v)", a, origin, point{k, l})
						continue outer
					}
				}
			}

			//log.Printf("android %v visible by origin %v", a, origin)
			// no blocking asteroids, line of sight!
			asteroidsInSight++
		}

		log.Printf("location %v can see %d asteroids", origin, asteroidsInSight)

		if maxAsteroidsInSight < asteroidsInSight {
			maxAsteroidsInSight = asteroidsInSight
			maxLocation = point{origin.x, origin.y}
		}
	}

	log.Printf("maxAsteroidsInSight %d, maxLocation %v", maxAsteroidsInSight, maxLocation)

	return []int{maxLocation.x, maxLocation.y}, maxAsteroidsInSight
}

type point struct {
	x, y int
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


func firstBlockingAsteroid(a, origin point, s string) point {
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

	// Q3 (between 3 and 6 o'clock)
	if a.x >= origin.x && a.y >= origin.y {
		deltaX := a.x - origin.x
		deltaY := a.y - origin.y
		log.Printf("deltaX %d, deltaY %d", deltaX, deltaY)

		// correct for 'smaller' distances
		if deltaX > 0 && deltaY > 0 {
			if deltaX%deltaY == 0 {
				deltaX = deltaX / deltaY
				deltaY = 1
			}
			if deltaY%deltaX == 0 {
				deltaY = deltaY / deltaX
				deltaX = 1
			}
		}
		if deltaX == 0 {
			deltaY = 1
		}
		if deltaY == 0 {
			deltaX = 1
		}
		log.Printf("deltaX %d, deltaY %d", deltaX, deltaY)

		for k, l := a.x-deltaX, a.y-deltaY; k >= origin.x && l >= origin.y; k, l = k-deltaX, l-deltaY {
			if coords[k][l] == "#" {
				if k == origin.x && l == origin.y {
					break
				}
				// found blocking asteroid, continue
				return point{k, l}
			}
		}
	}

	return point{-1, -1}
}
