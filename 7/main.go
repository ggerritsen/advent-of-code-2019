package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
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
	log.Printf("Max output: %d", findOptimalOutput(input))
	log.Printf("End part 1")

	log.Printf("Start part 2")
	log.Printf("Max output with feedback loop: %d", findOptimalOutputWithFeedbackLoop(input))
	log.Printf("End part 2")
}

func findOptimalOutputWithFeedbackLoop(opCodes []int) int {
	phases := []int{5, 6, 7, 8, 9}
	permutations := generatePermutations(5, phases)

	maxOutput := 0
	for _, p := range permutations {
		pcs := make([]*intCodePC, 5)
		for i := 0; i<5; i++ {
			pcs[i] = newIntCodePC(p[i], opCodes)
		}

		pcs[1].in = pcs[0].out
		pcs[2].in = pcs[1].out
		pcs[3].in = pcs[2].out
		pcs[4].in = pcs[3].out

		var wg sync.WaitGroup
		for i :=0; i<5; i++ {
			wg.Add(1)
			pc := pcs[i]
			go func() {
				defer func() {
					wg.Done()
				}()
				pc.run()
			}()
		}

		// start
		pcs[0].in <- 0

		// feed output of last amplifier back to first amplifier
		wg.Add(1)
		go func() {
			defer wg.Done()
			for o := range pcs[4].out {
				if o > maxOutput {
					maxOutput = o
				}
				pcs[0].in <- o
			}
		}()

		wg.Wait()
	}

	return maxOutput
}

func newIntCodePC(phase int, opCodes []int) *intCodePC {
	return &intCodePC{
		opCodes: opCodes,
		phase:   phase,
		in:      make(chan int, 1),
		out:     make(chan int, 1),
	}
}

type intCodePC struct {
	opCodes []int
	phase   int
	in, out chan int
}

func (pc *intCodePC) run() {
	result := make([]int, len(pc.opCodes))
	for j := 0; j < len(pc.opCodes); j++ {
		result[j] = pc.opCodes[j]
	}

	index, inputCount := 0, 0
	phaseInserted := false
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
			if phaseInserted {
				result[resultLocation] = <-pc.in
			} else {
				result[resultLocation] = pc.phase
				phaseInserted = true
			}
			index = index + 2
			inputCount++
		}
		// output
		if operator == 4 {
			resultLocation := result[index+1]
			pc.out <- result[resultLocation]
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

	close(pc.out)
}

func findOptimalOutput(opCodes []int) int {
	phases := []int{0, 1, 2, 3, 4}
	permutations := generatePermutations(5, phases)

	var outputs []int
	for _, p := range permutations {
		output := 0
		for i := 0; i < 5; i++ {
			_, output = runIntcode(opCodes, []int{p[i], output})
		}
		outputs = append(outputs, output)
	}

	sort.Ints(outputs)
	return outputs[len(outputs)-1]
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
