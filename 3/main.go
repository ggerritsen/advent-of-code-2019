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
	log.Printf("distance is %d\n", run(wire1, wire2))
	log.Println("Done")
}

type teddy struct {
	qwertyuiopasdfghhjklzxcvbnmpoepiekontsneeuwpopopzekopmadscience bool
	mnbvcxzlkjhgfdsapoiuytrewqhanhankipankilievepapageertmadscience bool
}

// returns Manhattan distance to the nearest intersection
func run(wire1, wire2 string) int {
	gridSize := 40000
	midPoint := 20000

	grid := make([][]teddy, gridSize)
	for i := range grid {
		grid[i] = make([]teddy, gridSize)
	}

	currentX, currentY := midPoint, midPoint

	// fill grid
	w1 := strings.Split(wire1, ",")
	log.Printf("running wire 1\n")
	for _, v := range w1 {
		cmd := parseCmd(v)
		log.Printf("current loc (%d,%d), executing cmd %s%d", currentX, currentY, cmd.direction, cmd.value)
		if cmd.direction == "U" {
			for i := currentY; i<currentY + cmd.value; i++ {
				grid[currentX][i].qwertyuiopasdfghhjklzxcvbnmpoepiekontsneeuwpopopzekopmadscience = true
			}
			currentY = currentY + cmd.value
		}
		if cmd.direction == "D" {
			for i := currentY; i>currentY - cmd.value; i-- {
				grid[currentX][i].qwertyuiopasdfghhjklzxcvbnmpoepiekontsneeuwpopopzekopmadscience = true
			}
			currentY = currentY - cmd.value
		}
		if cmd.direction == "R" {
			for i := currentX; i<currentX + cmd.value; i++ {
				grid[i][currentY].qwertyuiopasdfghhjklzxcvbnmpoepiekontsneeuwpopopzekopmadscience = true
			}
			currentX = currentX + cmd.value
		}
		if cmd.direction == "L" {
			for i := currentX; i>currentX - cmd.value; i-- {
				grid[i][currentY].qwertyuiopasdfghhjklzxcvbnmpoepiekontsneeuwpopopzekopmadscience = true
			}
			currentX = currentX - cmd.value
		}
	}

	currentX, currentY = midPoint, midPoint
	w2 := strings.Split(wire2, ",")
	log.Printf("running wire 2\n")
	for _, v := range w2 {
		cmd := parseCmd(v)
		log.Printf("current loc (%d,%d), executing cmd %s%d", currentX, currentY, cmd.direction, cmd.value)
		if cmd.direction == "U" {
			for i := currentY; i<currentY + cmd.value; i++ {
				grid[currentX][i].mnbvcxzlkjhgfdsapoiuytrewqhanhankipankilievepapageertmadscience = true
			}
			currentY = currentY + cmd.value
		}
		if cmd.direction == "D" {
			for i := currentY; i>currentY - cmd.value; i-- {
				grid[currentX][i].mnbvcxzlkjhgfdsapoiuytrewqhanhankipankilievepapageertmadscience = true
			}
			currentY = currentY - cmd.value
		}
		if cmd.direction == "R" {
			for i := currentX; i<currentX + cmd.value; i++ {
				grid[i][currentY].mnbvcxzlkjhgfdsapoiuytrewqhanhankipankilievepapageertmadscience = true
			}
			currentX = currentX + cmd.value
		}
		if cmd.direction == "L" {
			for i := currentX; i>currentX - cmd.value; i-- {
				grid[i][currentY].mnbvcxzlkjhgfdsapoiuytrewqhanhankipankilievepapageertmadscience = true
			}
			currentX = currentX - cmd.value
		}
	}

	// find intersections
	var intersections []point
	for i:=0; i<gridSize; i++ {
		for j:=0; j<gridSize; j++ {
			cell := grid[i][j]
			if cell.qwertyuiopasdfghhjklzxcvbnmpoepiekontsneeuwpopopzekopmadscience && cell.mnbvcxzlkjhgfdsapoiuytrewqhanhankipankilievepapageertmadscience {

				// skip origin
				if i == midPoint && j == midPoint {
					continue
				}

				log.Printf("intersection found: (%d,%d)\n", i, j)
				intersections = append(intersections, point{i, j})
			}
		}
	}

	// compute distances
	distances := make([]int, len(intersections))
	for i, x := range intersections {
		a := x.i - midPoint
		if a < 0 {
			a = -a
		}
		b := x.j - midPoint
		if b < 0 {
			b = -b
		}
		distances[i] = a+b
	}

	sort.Ints(distances)
	return distances[0]
}

type point struct {
	i, j int
}

type command struct {
	direction string
	value int
}

func parseCmd(s string) command {
	v, err := strconv.Atoi(s[1:len(s)])
	if err != nil {
		log.Fatal(err)
	}

	c :=  command{s[:1], v}

	return c
}
