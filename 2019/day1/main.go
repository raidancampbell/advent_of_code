package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Main function
func main() {
	str, err := FileToString("2019/day1/input")
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
		f := calcFuel(size)
		if f < 0 {
			f = 0
		}
		acc += f
		return calcFuel2(f, acc)
	}
	return acc
}



func FileToString(fname string) (result string, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}
	err = file.Close()

	result = string(b)
	return
}