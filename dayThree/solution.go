package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
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
	lines := strings.Split(string(dat), "\n")
	fmt.Println("Line 1: ", lines[0])
	move("L319")
}

func move(cmd string) {
	fmt.Println(reflect.TypeOf(cmd[0]))
	switch string(cmd[0]) {
	case "L":
		fmt.Println("Left")
	case "R":
		fmt.Println("Right")
	case "U":
		fmt.Println("Up")
	case "D":
		fmt.Println("Down")
	}

}

func main() {
	// Part 1
	part1()
}
