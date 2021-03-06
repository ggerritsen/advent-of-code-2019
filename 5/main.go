package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Start")

	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s, err := csv.NewReader(f).Read()
	if err != nil {
		log.Fatal(err)
	}

	input := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ss, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		input[i] = ss
	}

	// keep a copy
	original := make([]int, len(input))
	for i := 0; i < len(s); i++ {
		original[i] = input[i]
	}

	log.Printf("Start part 1")
	log.Printf("input: %+v\n", input)
	log.Printf("output: %+v\n", run(input))
	log.Printf("End part 1")

	log.Println("End")
}

func run(i []int) []int {
	result := make([]int, len(i))
	for j := 0; j < len(i); j++ {
		result[j] = i[j]
	}

	index := 0

	operator, modes := parseOperator(result[index])
	for operator != 99 {
		log.Printf("Executing %d, %v\n", operator, modes)
		if operator == 1 || operator == 2 {
			x, y := result[index+1], result[index+2]
			resultLocation := result[index+3]

			a := x
			if modes[0] == 0 {
				// position mode
				a = result[x]
			}
			b := y
			if modes[1] == 0 {
				// position mode
				b = result[y]
			}

			// addition
			if operator == 1 {
				result[resultLocation] = a + b
			}
			// multiplication
			if operator == 2 {
				result[resultLocation] = a * b
			}

			index = index + 4
		}
		// input
		if operator == 3 {
			resultLocation := result[index+1]

			log.Printf("Give input\n")
			sc := bufio.NewScanner(os.Stdin)
			sc.Scan()
			input := sc.Text()
			if err := sc.Err(); err != nil {
				log.Fatal(err)
			}
			v, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}

			result[resultLocation] = v
			index = index + 2
		}
		// output
		if operator == 4 {
			resultLocation := result[index+1]
			log.Printf("OUTPUT: %d\n", result[resultLocation])
			index = index + 2
		}
		// jump-if-true
		if operator == 5 {
			x, y := result[index+1], result[index+2]

			a := x
			if modes[0] == 0 {
				// position mode
				a = result[x]
			}

			if a != 0 {
				index = y
				if modes[1] == 0 {
					// position mode
					index = result[y]
				}
			} else {
				index = index + 3
			}
		}
		// jump-if-false
		if operator == 6 {
			x, y := result[index+1], result[index+2]

			a := x
			if modes[0] == 0 {
				// position mode
				a = result[x]
			}

			if a == 0 {
				index = y
				if modes[1] == 0 {
					// position mode
					index = result[y]
				}
			} else {
				index = index + 3
			}
		}
		// less-than
		if operator == 7 {
			x, y := result[index+1], result[index+2]
			resultLocation := result[index+3]

			a := x
			if modes[0] == 0 {
				// position mode
				a = result[x]
			}
			b := y
			if modes[1] == 0 {
				// position mode
				b = result[y]
			}

			if a < b {
				result[resultLocation] = 1
			} else {
				result[resultLocation] = 0
			}

			index = index + 4
		}
		// equals
		if operator == 8 {
			x, y := result[index+1], result[index+2]
			resultLocation := result[index+3]

			a := x
			if modes[0] == 0 {
				// position mode
				a = result[x]
			}
			b := y
			if modes[1] == 0 {
				// position mode
				b = result[y]
			}

			if a == b {
				result[resultLocation] = 1
			} else {
				result[resultLocation] = 0
			}

			index = index + 4
		}

		operator, modes = parseOperator(result[index])
	}

	return result
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
