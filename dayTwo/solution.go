package main

import "fmt"

var test1 = [...]int{
	1, 0, 0, 0, 99,
}

var test2 = [...]int{
	2, 3, 0, 3, 99,
}

var test3 = [...]int{
	2, 4, 4, 5, 99, 0,
}

var test4 = [...]int{
	1, 1, 1, 4, 99, 5, 6, 0, 99,
}

func computer(params ...int) int {

	var input = [...]int{
		1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 19, 5, 23, 1, 6, 23, 27, 1, 27, 5, 31, 2, 31, 10, 35, 2, 35, 6, 39, 1, 39, 5, 43, 2, 43, 9, 47, 1, 47, 6, 51, 1, 13, 51, 55, 2, 9, 55, 59, 1, 59, 13, 63, 1, 6, 63, 67, 2, 67, 10, 71, 1, 9, 71, 75, 2, 75, 6, 79, 1, 79, 5, 83, 1, 83, 5, 87, 2, 9, 87, 91, 2, 9, 91, 95, 1, 95, 10, 99, 1, 9, 99, 103, 2, 103, 6, 107, 2, 9, 107, 111, 1, 111, 5, 115, 2, 6, 115, 119, 1, 5, 119, 123, 1, 123, 2, 127, 1, 127, 9, 0, 99, 2, 0, 14, 0,
	}

	if len(params) == 2 {
		input[1] = params[0]
		input[2] = params[1]
	}

	ic := 0

	for {
		if input[ic] == 99 {
			break
		}
		instruction := input[ic : ic+4]

		switch instruction[0] {
		case 1:
			input[instruction[3]] = input[instruction[1]] + input[instruction[2]]
		case 2:
			input[instruction[3]] = input[instruction[1]] * input[instruction[2]]
		}
		ic += 4
	}
	return input[0]
}

func part2() {
	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			solution := computer(noun, verb)
			if solution == 19690720 {
				fmt.Println("Part 2 Noun and Verb:", noun, verb, "Solution:", 100*noun+verb)
				break
			}
		}
	}
}

func main() {
	// Part 1
	fmt.Println("Part 1 solution:", computer())
	// Part2
	part2()
}
