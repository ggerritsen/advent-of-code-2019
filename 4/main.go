package main

import (
	"log"
)

func main() {
	log.Println("Start")

	start, end := 193651, 649729

	var solutions []int
	for i := start; i <= end; i++ {
		x := convertToIntSlice(i, []int{})

		if isDecreasing(x) {
			continue
		}
		if containsAdjacentEquals(x) {
			solutions = append(solutions, i)
		}
	}

	log.Printf("Solutions found: %d", len(solutions))

	log.Println("End")
}

func convertToIntSlice(i int, result []int) []int {
	for i > 0 {
		j := i % 10
		result = append([]int{j}, result...)
		return convertToIntSlice(i/10, result)
	}
	return result
}

func isDecreasing(number []int) bool {
	for i := 0; i < len(number)-1; i++ {
		a := number[i]
		b := number[i+1]
		if b < a {
			return true
		}
	}
	return false
}

func containsAdjacentEquals(number []int) bool {
	for i := 0; i < len(number)-1; i++ {
		if number[i] == number[i+1] {
			return true
		}
	}
	return false
}
