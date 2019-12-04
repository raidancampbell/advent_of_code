package main

import (
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"strconv"
	"strings"
)

const(
	OPCODE_ADD = 1
	OPCODE_MULTIPLY = 2
	OPCODE_TERMINATE = 99
)

func main() {
	str, err := rstrings.FileToString("2019/day2/input")
	if err != nil {
		fmt.Printf("Error encountered reading file.  Ensure working directory is repository root: %v\n", err)
	}

	elements := parseAndRepair(str, 12, 2)
	fmt.Printf("result %v\n", loop(elements)) // part 1 answer

	// part 2 brute force
	for noun := 0; noun < 100; noun++{
		for verb := 0; verb < 100; verb++{

			str, err := rstrings.FileToString("2019/day2/input")
			if err != nil {
				fmt.Printf("Error encountered reading file.  Ensure working directory is repository root: %v\n", err)
			}

			elements := parseAndRepair(str, noun, verb)
			result := loop(elements)
			if result == 19690720 {
				fmt.Printf("noun %v verb %v\n", noun, verb)
				return
			}

		}
	}
}

func loop(elements []int) int {
	i := 0
	for {
		// read the op code
		switch elements[i] {
		case OPCODE_ADD:
			doAdd(elements, i)
		case OPCODE_MULTIPLY:
			doMultiply(elements, i)
		case OPCODE_TERMINATE:
			return elements[0]
		}
		i += 4
	}
}

func doAdd(elements []int, i int) {
	i++
	elements[elements[i+2]] = elements[elements[i]] + elements[elements[i+1]]
}

func doMultiply(elements []int, i int) {
	i++
	elements[elements[i+2]] = elements[elements[i]] * elements[elements[i+1]]
}

// before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
func parseAndRepair(input string, noun, verb int) []int {
	elements := strings.Split(input, ",")
	elements[1] = strconv.Itoa(noun)
	elements[2] = strconv.Itoa(verb)

	ret := make([]int, len(elements))
	for i, _ := range elements {
		ret[i], _ = strconv.Atoi(elements[i])
	}
	return ret
}