package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	log.Printf("Start part 1")

	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var input []string
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	log.Printf("orbit map checksum: %d", calculateNumOrbits(input))
	log.Printf("End part 1")

	log.Printf("Start part 2")
	log.Printf("min transfers to Santa: %d", minOrbitTransfers(input))
	log.Printf("End part 2")
}

func calculateNumOrbits(s []string) int {
	orbitMap := map[string]string{}
	for _, o := range s {
		ss := strings.Split(o, ")")
		orbitMap[ss[1]] = ss[0]
	}

	numOrbits := 0
	for _, v := range orbitMap {
		for v != "COM" {
			numOrbits++
			v = orbitMap[v]
		}
		numOrbits++
	}

	return numOrbits
}

func minOrbitTransfers(s []string) int {
	orbitMap := map[string]string{}
	for _, o := range s {
		ss := strings.Split(o, ")")
		orbitMap[ss[1]] = ss[0]
	}

	var transfers []int

	target := "SAN"
	iteration := -1
outer:
	for target != "COM" {
		iteration++
		numTransfers := iteration
		current := orbitMap["YOU"]
		target = orbitMap[target]
		for current != target {
			// do 1 transfer
			current = orbitMap[current]
			numTransfers++

			// arrived at COM - this is not a viable route, try next iteration
			if current == "COM" {
				continue outer
			}
		}
		// arrived at target! Store and continue next iteration
		transfers = append(transfers, numTransfers)
	}

	sort.Ints(transfers)
	return transfers[0]
}
