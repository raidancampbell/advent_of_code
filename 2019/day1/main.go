package main

import (
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rmath"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"strconv"
	"strings"
)

// Main function
func main() {
	str, err := rstrings.FileToString("2019/day1/input")
	if err != nil {
		fmt.Printf("Error encountered reading file.  Ensure working directory is repository root: %v\n", err)
	}

	sum := 0
	sum2 := 0
	// for each line, ignoring the trailing newline
	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		// module size
		m, _ := strconv.Atoi(line)
		// fuel requirement
		sum += calcFuel(m)
		sum2 += calcFuel2(m, 0)
	}
	fmt.Printf("Part 1: %v\n", sum)
	fmt.Printf("Part 2: %v\n", sum2)
}

func calcFuel(size int) int {
	return (size / 3) - 2
}

func calcFuel2(size, acc int) int {
	for size > 0 {
		f := rmath.Max(calcFuel(size), 0)
		acc += f
		return calcFuel2(f, acc)
	}
	return acc
}
