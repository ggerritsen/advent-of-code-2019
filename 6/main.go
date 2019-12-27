package main

import (
	"bufio"
	"log"
	"os"
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

	log.Printf("orbit map checksum: %d", run(input))

	log.Printf("End part 1")

}

func run(s []string) int {
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
