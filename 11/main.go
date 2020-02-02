package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
run()
}

const gridSize = 75
const midPoint = 37

func run() {
	// 0 is black, 1 is white
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(dir + "/11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s, err := csv.NewReader(f).Read()
	if err != nil {
		log.Fatal(err)
	}

	opCodes := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ss, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		opCodes[i] = ss
	}

	pc := newIntCodePC(opCodes)
	loc := point{midPoint, midPoint}
	//orientations := []string{"W", "N", "E", "S"}
	orientation := 1

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		pc.run()
	}()

	painted := map[point]bool{}
	for {
		log.Printf("Loc is %v", loc)
		pc.in <- grid[loc.x][loc.y]

		newColor, ok := <-pc.out
		if !ok {
			break
		}
		grid[loc.x][loc.y] = newColor
		painted[loc] = true

		direction := <-pc.out
		if direction == 0 {
			orientation--
			if orientation == -1 {
				orientation = 3
			}
		}
		if direction == 1 {
			orientation++
			if orientation == 4 {
				orientation = 0
			}
		}

		switch orientation {
		case 0:
			loc = point{loc.x - 1, loc.y}
		case 1:
			loc = point{loc.x, loc.y + 1}
		case 2:
			loc = point{loc.x + 1, loc.y}
		case 3:
			loc = point{loc.x, loc.y - 1}
		}
	}

	wg.Wait()

	log.Printf("Done: %v", grid)
	log.Printf("Painted: %d", len(painted))
}

type point struct {
	x, y int
}

func newIntCodePC(opCodes []int) *intCodePC {
	return &intCodePC{
		opCodes: opCodes,
		in:      make(chan int, 1),
		out:     make(chan int, 1),
	}
}

type intCodePC struct {
	opCodes []int
	in, out chan int
	result  []int
}

func (pc *intCodePC) run() {
	pc.result = make([]int, len(pc.opCodes)+100000)
	for j := 0; j < len(pc.opCodes); j++ {
		pc.result[j] = pc.opCodes[j]
	}

	index, relativeBase := 0, 0
	operator, modes := parseOperator(pc.result[index])
	for operator != 99 {
		if operator == 1 || operator == 2 {
			x, y := pc.result[index+1], pc.result[index+2]
			resultLocation := pc.result[index+3]

			a := x
			if modes[0] == 0 {
				// position mode
				a = pc.result[x]
			}
			if modes[0] == 2 {
				// relative mode
				a = pc.result[x+relativeBase]
			}
			b := y
			if modes[1] == 0 {
				// position mode
				b = pc.result[y]
			}
			if modes[1] == 2 {
				// relative mode
				b = pc.result[y+relativeBase]
			}

			// addition
			if operator == 1 {
				if modes[2] == 0 {
					// position mode
					pc.result[resultLocation] = a + b
				}
				if modes[2] == 2 {
					// relative mode
					pc.result[resultLocation+relativeBase] = a + b
				}
			}
			// multiplication
			if operator == 2 {
				if modes[2] == 0 {
					// position mode
					pc.result[resultLocation] = a * b
				}
				if modes[2] == 2 {
					// relative mode
					pc.result[resultLocation+relativeBase] = a * b
				}
			}

			index = index + 4
		}
		// input
		if operator == 3 {
			in := <-pc.in

			resultLocation := pc.result[index+1]
			if modes[0] == 0 {
				// position mode
				pc.result[resultLocation] = in
			}
			if modes[0] == 2 {
				// relative mode
				pc.result[resultLocation+relativeBase] = in
			}

			index = index + 2
		}
		// output
		if operator == 4 {
			x := pc.result[index+1]

			if modes[0] == 0 {
				// position mode
				pc.out <- pc.result[x]
			}
			if modes[0] == 1 {
				// direct mode
				pc.out <- x
			}
			if modes[0] == 2 {
				// relative mode
				pc.out <- pc.result[x+relativeBase]
			}

			index = index + 2
		}
		// jump-if-true
		if operator == 5 {
			x, y := pc.result[index+1], pc.result[index+2]

			a := x
			if modes[0] == 0 {
				// position mode
				a = pc.result[x]
			}
			if modes[0] == 2 {
				// relative mode
				a = pc.result[x+relativeBase]
			}

			if a != 0 {
				index = y
				if modes[1] == 0 {
					// position mode
					index = pc.result[y]
				}
				if modes[1] == 2 {
					// relative mode
					index = pc.result[y+relativeBase]
				}
			} else {
				index = index + 3
			}
		}
		// jump-if-false
		if operator == 6 {
			x, y := pc.result[index+1], pc.result[index+2]

			a := x
			if modes[0] == 0 {
				// position mode
				a = pc.result[x]
			}
			if modes[0] == 2 {
				// relative mode
				a = pc.result[x+relativeBase]
			}

			if a == 0 {
				index = y
				if modes[1] == 0 {
					// position mode
					index = pc.result[y]
				}
				if modes[1] == 2 {
					// relative mode
					index = pc.result[y+relativeBase]
				}
			} else {
				index = index + 3
			}
		}
		// less-than
		if operator == 7 {
			x, y := pc.result[index+1], pc.result[index+2]
			resultLocation := pc.result[index+3]

			a := x
			if modes[0] == 0 {
				// position mode
				a = pc.result[x]
			}
			if modes[0] == 2 {
				// relative mode
				a = pc.result[x+relativeBase]
			}
			b := y
			if modes[1] == 0 {
				// position mode
				b = pc.result[y]
			}
			if modes[1] == 2 {
				// relative mode
				b = pc.result[y+relativeBase]
			}

			if a < b {
				if modes[2] == 0 {
					// position mode
					pc.result[resultLocation] = 1
				}
				if modes[2] == 2 {
					// relative mode
					pc.result[resultLocation+relativeBase] = 1
				}
			} else {
				if modes[2] == 0 {
					// position mode
					pc.result[resultLocation] = 0
				}
				if modes[2] == 2 {
					// relative mode
					pc.result[resultLocation+relativeBase] = 0
				}
			}

			index = index + 4
		}
		// equals
		if operator == 8 {
			x, y := pc.result[index+1], pc.result[index+2]
			resultLocation := pc.result[index+3]

			a := x
			if modes[0] == 0 {
				// position mode
				a = pc.result[x]
			}
			if modes[0] == 2 {
				// relative mode
				a = pc.result[x+relativeBase]
			}
			b := y
			if modes[1] == 0 {
				// position mode
				b = pc.result[y]
			}
			if modes[1] == 2 {
				// relative mode
				b = pc.result[y+relativeBase]
			}

			if a == b {
				if modes[2] == 0 {
					// position mode
					pc.result[resultLocation] = 1
				}
				if modes[2] == 2 {
					// relative mode
					pc.result[resultLocation+relativeBase] = 1
				}
			} else {
				if modes[2] == 0 {
					// position mode
					pc.result[resultLocation] = 0
				}
				if modes[2] == 2 {
					// relative mode
					pc.result[resultLocation+relativeBase] = 0
				}
			}

			index = index + 4
		}
		// adjust relative base
		if operator == 9 {
			x := pc.result[index+1]

			if modes[0] == 0 {
				// position mode
				relativeBase = relativeBase + pc.result[x]
			}
			if modes[0] == 1 {
				// direct mode
				relativeBase = relativeBase + x
			}
			if modes[0] == 2 {
				// relative mode
				relativeBase = relativeBase + pc.result[x+relativeBase]
			}

			index = index + 2
		}
		if operator != 99 && operator > 9 {
			log.Fatalf("Unknown operator %d (modes %v)", operator, modes)
		}

		operator, modes = parseOperator(pc.result[index])
	}

	close(pc.out)
}

func parseOperator(i int) (int, []int) {
	operand := i % 100
	modes := convertToReversedIntSlice(i/100, nil)

	for i := len(modes); i < 3; i++ {
		modes = append(modes, 0)
	}

	return operand, modes
}

func convertToReversedIntSlice(i int, result []int) []int {
	for i > 0 {
		j := i % 10
		result = append(result, j)
		return convertToReversedIntSlice(i/10, result)
	}
	return result
}
