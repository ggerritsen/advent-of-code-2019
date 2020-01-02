package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

	log.Printf("Start part 1")
	log.Printf("Test boost output: %v", runBoostWithInput(opCodes, 1))
	log.Printf("End part 1")

	log.Printf("Start part 2")
	log.Printf("Executing boost output: %v", runBoostWithInput(opCodes, 2))
	log.Printf("End part 2")
}

func runBoostWithInput(opCodes []int, input int) []int {
	var wg sync.WaitGroup

	pc := newIntCodePC(opCodes)
	wg.Add(1)
	go func() {
		defer wg.Done()
		pc.run()
	}()

	var result []int
	wg.Add(1)
	go func() {
		defer wg.Done()
		for o := range pc.out {
			//log.Printf("out: %d", o)
			result = append(result, o)
		}
	}()

	pc.in <- input

	wg.Wait()

	return result
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
		//log.Printf("Exec %d %v", operator, modes)
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
