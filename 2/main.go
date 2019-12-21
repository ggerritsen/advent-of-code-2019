package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Start")

	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/2/input.txt")
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

	log.Printf("Start part 2")
	var result []int
	outer: for i:= 0; i<100; i++ {
		for j:=0; j<100; j++ {
			original[1] = i
			original[2] = j
			result = run(original)

			if result[0] == 19690720 {
				break outer
			}
		}
	}
	log.Printf("output: %+v\n", result)
	log.Printf("End part 2")

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
		x, y := result[index+1], result[index+2]
		resultLocation := result[index+3]
		if operand == 1 {
			result[resultLocation] = result[x] + result[y]
		}
		if operand == 2 {
			result[resultLocation] = result[x] * result[y]
		}

		index = index + 4
		operand = result[index]
	}

	return result
}
