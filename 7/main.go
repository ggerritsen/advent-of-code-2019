package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Printf("Start")

	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/7/input.txt")
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

	_, o := runIntcode(input, []int{3, 0})
	log.Printf("Output: %d", o)

	log.Printf("End part 1")
}

func findOptimalPhases(opCodes []int) []int {
	return []int{0}
}

// source: https://en.wikipedia.org/wiki/Heap%27s_algorithm
//procedure generate(k : integer, A : array of any):
//    if k = 1 then
//        output(A)
//    else
//        // Generate permutations with kth unaltered
//        // Initially k == length(A)
//        generate(k - 1, A)
//
//        // Generate permutations for kth swapped with each k-1 initial
//        for i := 0; i < k-1; i += 1 do
//            // Swap choice dependent on parity of k (even or odd)
//            if k is even then --> gg: removed this if, seems to cause a bug (in a 3-elem list, the 2nd element never gets put in the 3rd place)
//                swap(A[i], A[k-1]) // zero-indexed, the kth is at k-1
//            else
//                swap(A[0], A[k-1])
//            end if
//            generate(k - 1, A)
//
//        end for
//    end if
func generatePermutations(k int, a []int) [][]int {
	if k == 1 {
		return [][]int{a}
	}

	c := generatePermutations(k-1, a)

	aa := make([]int, len(a))
	copy(aa, a)
	for i := 0; i < k-1; i++ {
		tmp := aa[i]
		aa[i] = aa[k-1]
		aa[k-1] = tmp

		aaa := make([]int, len(aa))
		copy(aaa, aa)
		c = append(c, generatePermutations(k-1, aaa)...)
	}

	return c
}

func runAmplifiers(intCodes []int, phases []int) int {
	output := 0
	for i := 0; i < len(phases); i++ {
		_, output = runIntcode(intCodes, []int{phases[i], output})
	}

	return output
}

func runIntcode(i []int, input []int) ([]int, int) {
	result := make([]int, len(i))
	for j := 0; j < len(i); j++ {
		result[j] = i[j]
	}
	output := 0

	index, inputCount := 0, 0
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
			result[resultLocation] = input[inputCount]
			index = index + 2
			inputCount++
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
