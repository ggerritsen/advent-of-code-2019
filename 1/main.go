package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Start")

	f, err := os.Open("/Users/ggerritsen/dev/personal/advent-of-code-2019/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	part1(f)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	part2(f)

	log.Println("End")
}

func part1(r io.Reader) {
	fuelNeeded := 0
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		mass, err := strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel := mass/3 - 2
		fuelNeeded += fuel
	}

	log.Printf("Fuel needed: %d", fuelNeeded)
}

func part2(r io.Reader) int {
	fuelNeeded := 0

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fuel, err := strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
		}

		for fuel > 0 {
			newFuel := calcFuel(fuel)
			if newFuel > 0 {
				fuelNeeded += newFuel
			}

			fuel = newFuel
		}
	}

	log.Printf("Fuel needed: %d", fuelNeeded)

	return fuelNeeded
}

func calcFuel(i int) int {
	return i/3 - 2
}
