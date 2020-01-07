package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(dir + "/10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start part 1")
	p, i := findBestMonitoringLocation(string(b))
	log.Printf("Best location %v (numAsteroids %d)", p, i)
	log.Printf("End part 1")

	log.Printf("Start part 2")
	a := vaporizeAsteroids(string(b), point{29,28})
	log.Printf("200th asteroid to be vaporized: %v", a[199])
	log.Printf("End part 2")
}

func vaporizeAsteroids(s string, monitoringLocation point) ([]point) {
	// parse
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

	// find all asteroids
	angles := map[float64][]point{}
	var asteroids []point
	for i, row := range coords {
		for j, v := range row {
			if i == monitoringLocation.x && j == monitoringLocation.y {
				continue
			}

			if v == "#" {
				p := point{i, j}
				a := monitoringLocation.angle(p)
				angles[a] = append(angles[a], p)
				asteroids = append(asteroids, p)
			}
		}
	}

	for _, v := range angles {
		sort.Slice(v, func(i, j int) bool {
			distanceI := (v[i].x-monitoringLocation.x) + (v[i].y-monitoringLocation.y)
			if distanceI < 0 {
				distanceI = -distanceI
			}
			distanceJ := (v[j].x-monitoringLocation.x) + (v[j].y-monitoringLocation.y)
			if distanceJ < 0 {
				distanceJ = -distanceJ
			}

			return distanceI < distanceJ
		})
	}
	log.Printf("Angles after sorting: %v", angles)

	// collect all angles in order
	var anglesIdx []float64
	for k, _ := range angles {
		anglesIdx = append(anglesIdx, k)
	}
	sort.Slice(anglesIdx, func(i, j int) bool {
		return anglesIdx[i] < anglesIdx[j]
	})
	log.Printf("Angle idxs: %v", anglesIdx)

	start := -1
	for i := len(anglesIdx)-1; i>= 0; i-- {
		if anglesIdx[i] <= math.Pi/2 {
			start = i
			break
		}
	}
	log.Printf("start is %d: %v", start, anglesIdx[start])

	var removedAsteroidsInOrder []point

	// 1st pass (3/4 circle)
	for i := start; i>= 0; i-- {
		a := anglesIdx[i]
		if curAngles, ok := angles[a]; ok {
		log.Printf("ngle is %v", a)
			log.Printf("curAngles is %v", curAngles)
			removedAsteroidsInOrder = append(removedAsteroidsInOrder, curAngles[0])
			//log.Printf("removed asteroids is %v", removedAsteroidsInOrder)
			if len(curAngles) == 1 {
				delete(angles, a)
				continue
			}
			angles[a] = curAngles[1:] // remove first asteroid at this angle
		}
	}

	// rest of the passes
	for len(angles) > 0 {
		for i := len(anglesIdx) - 1; i >= 0; i-- {
			a := anglesIdx[i]
			if curAngles, ok := angles[a]; ok {
			log.Printf("ngle is %v", a)
				log.Printf("curAngles is %v", curAngles)
				removedAsteroidsInOrder = append(removedAsteroidsInOrder, curAngles[0])
				//log.Printf("removed asteroids is %v", removedAsteroidsInOrder)
				if len(curAngles) == 1 {
					delete(angles, a)
					continue
				}
				angles[a] = curAngles[1:] // remove first asteroid at this angle
			}
		}
	}

	log.Printf("map size %d", len(angles))
	log.Printf("Removed asteroids: %v", removedAsteroidsInOrder)
	return removedAsteroidsInOrder
}

func findBestMonitoringLocation(s string) (point, int) {
	// parse
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

	// find all asteroids
	var asteroids []point
	for i, row := range coords {
		for j, v := range row {
			if v == "#" {
				asteroids = append(asteroids, point{i, j})
			}
		}
	}

	// find max asteroids
	maxAsteroidsInSight := 0
	maxLocation := point{-1, -1}
	for _, origin := range asteroids {

		angles := map[float64]bool{}
		for _, a := range asteroids {
			if origin == a {
				continue
			}

			angle := origin.angle(a)
			log.Printf("DEBUG: angle %.2f", angle)
			angles[angle] = true
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
	return math.Atan2(float64(deltaY), float64(deltaX))
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
