package main

import (
	"log"
)

func main() {
	
}

func runIntcode(i []int, input int) ([]int, int) {
	result := make([]int, len(i))
	for j := 0; j < len(i); j++ {
		result[j] = i[j]
	}
	output := 0

	index := 0
	operator, modes := parseOperator(result[index])
	for operator != 99 {
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
			result[resultLocation] = input
			index = index + 2
		}
		// output
		if operator == 4 {
			resultLocation := result[index+1]
			output = result[resultLocation]
			log.Printf("OUTPUT: %d\n", output)
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

	return result, output
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
