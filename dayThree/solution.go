package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	geo "github.com/paulmach/go.geo"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() {
	absPath, pathErr := filepath.Abs("dayThree/input.txt")
	check(pathErr)
	dat, datErr := ioutil.ReadFile(absPath)
	check(datErr)

	paths := make([]*geo.Path, 0)

	wires := strings.Split(string(dat), "\n")
	for _, wire := range wires {
		instructions := strings.Split(string(wire), ",")
		paths = append(paths, createCoordinates(instructions))
	}

	intersectionPoints, _ := paths[0].IntersectionPath(paths[1])
	fmt.Println("Intersection points", intersectionPoints)
	shortestDistance := 9999999.99
	var shortestDistancePoint *geo.Point

	for _, intersectionPoint := range intersectionPoints {
		distance := math.Abs(intersectionPoint[0]) + math.Abs(intersectionPoint[1])
		if distance < shortestDistance {
			shortestDistance = distance
			shortestDistancePoint = intersectionPoint
		}
	}
	fmt.Println("Intersection", shortestDistancePoint, "at distance", shortestDistance)
}

func createCoordinates(commands []string) *geo.Path {

	path := geo.NewPath()
	x := 0
	y := 0

	for _, command := range commands {
		direction := string(command[0])
		steps, err := strconv.Atoi(string(command[1:]))
		check(err)

		switch direction {
		case "L":
			x = x - steps
		case "R":
			x = x + steps
		case "U":
			y = y + steps
		case "D":
			y = y - steps
		}

		path.Push(geo.NewPoint(float64(x), float64(y)))
	}
	return path
}

func main() {
	// Part 1
	part1()
}
