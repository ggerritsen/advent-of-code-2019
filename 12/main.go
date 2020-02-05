package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

}

func parse(s []string) []moon {
	result := make([]moon, len(s))
	for i, ss := range s {
		result[i] = parseLine(ss)
	}

	return result
}

func parseLine(s string) moon {
	// <x=8, y=0, z=8>
	s2 := strings.Trim(s, "<>")
	ss := strings.Split(s2, ",")

	xIdx := strings.Index(ss[0], "x=")
	xStr := ss[0][xIdx+2:]
	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Fatal(err)
	}

	yIdx := strings.Index(ss[1], "y=")
	yStr := ss[1][yIdx+2:]
	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Fatal(err)
	}

	zIdx := strings.Index(ss[2], "z=")
	zStr := ss[2][zIdx+2:]
	z, err := strconv.Atoi(zStr)
	if err != nil {
		log.Fatal(err)
	}

	return moon{pos: position{x, y, z}}
}

func iterate(moons []moon) []moon {
	f := func(i int) int {
		if i == 0 {
			return 0
		}
		if i < 0 {
			return -1
		}
		return 1
	}

	// apply gravity
	for i := 0; i < len(moons); i++ {
		this := &moons[i]
		for j := 0; j < len(moons); j++ {
			if i == j {
				continue
			}
			other := moons[j]
			this.vel.x += f(other.pos.x - this.pos.x)
			this.vel.y += f(other.pos.y - this.pos.y)
			this.vel.z += f(other.pos.z - this.pos.z)
		}
	}

	// apply velocity
	for i := 0; i < len(moons); i++ {
		moons[i].pos.x += moons[i].vel.x
		moons[i].pos.y += moons[i].vel.y
		moons[i].pos.z += moons[i].vel.z
	}

	return moons
}

func newMoon(posX, posY, posZ, velX, velY, velZ int) moon {
	return moon{
		pos: position{x: posX, y: posY, z: posZ},
		vel: velocity{x: velX, y: velY, z: velZ},
	}
}

type moon struct {
	pos position
	vel velocity
}

func (m *moon) String() string {
	return fmt.Sprintf("%v", *m)
}

type position struct {
	x, y, z int
}

type velocity struct {
	x, y, z int
}

func run(m []*moon) {

}
