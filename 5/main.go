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

	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/5/sample_io.txt")
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

	operand := result[index]
	for operand != 99 {

		// addition
		if operand == 1 {
			x, y := result[index+1], result[index+2]
			resultLocation := result[index+3]
			result[resultLocation] = result[x] + result[y]
			index = index + 4
		}
		// multiplication
		if operand == 2 {
			x, y := result[index+1], result[index+2]
			resultLocation := result[index+3]
			result[resultLocation] = result[x] * result[y]
			index = index + 4
		}
		// input
		if operand == 3 {
			resultLocation := result[index+1]

			log.Printf("Give input\n")
			sc := bufio.NewScanner(os.Stdin)
			sc.Scan()
			input := sc.Text()
			v, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}

			result[resultLocation] = v
			index = index + 2
		}
		// output
		if operand == 4 {
			resultLocation := result[index+1]
			log.Printf("OUTPUT: %d\n", result[resultLocation])
			index = index + 2
		}

		operand = result[index]
	}

	return result
}
