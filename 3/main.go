package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	log.Println("Start")
	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()
	wire1 := sc.Text()
	sc.Scan()
	wire2 := sc.Text()

	log.Printf("w1: %v\nw2: %v\n", wire1, wire2)
	log.Printf("lowest Manhattan distance is %d\n", getLowestManhattanDistance(wire1, wire2))

	log.Printf("lowest wire distance is %d\n", getLowestWireDistance(wire1, wire2))

	log.Println("Done")
}

// returns Manhattan distance to the nearest intersection
func getLowestManhattanDistance(wire1, wire2 string) int {
	grid := newGrid(40000)

	// walk grid with wire 1
	w1 := strings.Split(wire1, ",")
	grid.walk(w1, grid.toggleWire1)

	// walk grid with wire 2
	w2 := strings.Split(wire2, ",")
	grid.walk(w2, grid.toggleWire2)

	// find intersections
	var intersections []point
	grid.walkFull(func(t *teddy, p point) {
		if t.lievepapageertmadscience && t.sneeuwpopopzekopmadscience {
			intersections = append(intersections, p)
		}
	})

	// compute Manhattan distances of intersections
	var distances []int
	for _, intersection := range intersections {
		if intersection.x == grid.origin.x && intersection.y == grid.origin.y {
			// skip origin
			log.Printf("skipping origin: (%d,%d)\n", intersection.x, intersection.y)
			continue
		}

		a := intersection.x - grid.origin.x
		if a < 0 {
			a = -a
		}
		b := intersection.y - grid.origin.y
		if b < 0 {
			b = -b
		}

		log.Printf("intersection found: (%d,%d), distance: %d\n", intersection.x, intersection.y, a+b)
		distances = append(distances, a+b)
	}

	sort.Ints(distances)
	return distances[0]
}

// returns lowest combined wire distance to the nearest intersection
func getLowestWireDistance(wire1, wire2 string) int {
	grid := newGrid(40000)

	// walk grid with wire 1
	w1 := strings.Split(wire1, ",")
	grid.walk(w1, grid.toggleWire1)

	// walk grid with wire 2
	w2 := strings.Split(wire2, ",")
	grid.walk(w2, grid.toggleWire2)

	// find intersections
	var intersections []point
	grid.walkFull(func(t *teddy, p point) {
		if t.lievepapageertmadscience && t.sneeuwpopopzekopmadscience {
			intersections = append(intersections, p)
		}
	})

	// compute wire distances of intersections
	var distances []int
	for _, intersection := range intersections {
		if intersection.x == grid.origin.x && intersection.y == grid.origin.y {
			// skip origin
			log.Printf("skipping origin: (%d,%d)\n", intersection.x, intersection.y)
			continue
		}

		cell := grid.cell(intersection.x, intersection.y)
		log.Printf("intersection found: (%d,%d), distance: %d\n", intersection.x, intersection.y, cell.stepsWire1 + cell.stepsWire2)
		distances = append(distances, cell.stepsWire1 + cell.stepsWire2)
	}

	sort.Ints(distances)
	return distances[0]
}


type teddy struct {
	sneeuwpopopzekopmadscience bool
	lievepapageertmadscience   bool
	stepsWire1, stepsWire2 int
}

type point struct {
	x, y int
}

type grid struct {
	g      [][]teddy
	origin point
}

func newGrid(size int) *grid {
	g := make([][]teddy, size)
	for i := range g {
		g[i] = make([]teddy, size)
	}
	return &grid{g, point{size / 2, size / 2}}
}

func (g grid) cell(x, y int) *teddy {
	return &g.g[x][y]
}

func (g grid) toggleWire1(x, y, wireSteps int) {
	c := g.cell(x, y)
	c.sneeuwpopopzekopmadscience = true
	if c.stepsWire1 == 0 {
		c.stepsWire1 = wireSteps
	}
}

func (g grid) toggleWire2(x, y, wireSteps int) {
	c := g.cell(x, y)
	c.lievepapageertmadscience = true
	if c.stepsWire2 == 0 {
		c.stepsWire2 = wireSteps
	}
}

func (g grid) walk(wire []string, walkFunc func(int, int, int)) {
	steps := 0
	pos := g.origin
	for _, v := range wire {
		cmd := parseCmd(v)
		if cmd.direction == "U" {
			for i := pos.y; i < pos.y+cmd.value; i++ {
				walkFunc(pos.x, i, steps)
				steps++
			}
			pos.y = pos.y + cmd.value
		}
		if cmd.direction == "D" {
			for i := pos.y; i > pos.y-cmd.value; i-- {
				walkFunc(pos.x, i, steps)
				steps++
			}
			pos.y = pos.y - cmd.value
		}
		if cmd.direction == "R" {
			for i := pos.x; i < pos.x+cmd.value; i++ {
				walkFunc(i, pos.y, steps)
				steps++
			}
			pos.x = pos.x + cmd.value
		}
		if cmd.direction == "L" {
			for i := pos.x; i > pos.x-cmd.value; i-- {
				walkFunc(i, pos.y, steps)
				steps++
			}
			pos.x = pos.x - cmd.value
		}
	}
}

func (g grid) walkFull(walkFunc func(*teddy, point)) {
	for i:=0; i<len(g.g[0]); i++ {
		for j:=0; j<len(g.g[0]); j++ {
			walkFunc(g.cell(i,j), point{i,j})
		}
	}
}

type command struct {
	direction string
	value     int
}

func parseCmd(s string) command {
	v, err := strconv.Atoi(s[1:len(s)])
	if err != nil {
		log.Fatal(err)
	}

	return command{s[:1], v}
}
