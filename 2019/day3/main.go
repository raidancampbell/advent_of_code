package main

import (
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rmath"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"math"
	"strconv"
	"strings"
)


type coord struct {
	x, y int
}

func main() {
	str, err := rstrings.FileToString("2019/day3/input")
	if err != nil {
		fmt.Printf("Error encountered reading file.  Ensure working directory is repository root: %v\n", err)
	}
	grid := map[coord]int{}

	// map out the first line
	line := strings.Split(str, "\n")[0]
	x, y, steps := 0, 0, 0
	for _, command := range strings.Split(line, ","){
		magnitude, _ := strconv.Atoi(command[1:])
		switch command[:1] {
		case "L":
			for i := 0; i < magnitude; i ++ {
				x--
				steps++
				grid[coord{x, y}] = steps
			}
		case "R":
			for i := 0; i < magnitude; i ++ {
				x++
				steps++
				grid[coord{x, y}] = steps
			}
		case "U":
			for i := 0; i < magnitude; i ++ {
				y++
				steps++
				grid[coord{x, y}] = steps
			}
		case "D":
			for i := 0; i < magnitude; i ++ {
				y--
				steps++
				grid[coord{x, y}] = steps
			}
		}
	}

	// walk the second line, checking for intersections
	line = strings.Split(str, "\n")[1]
	x, y, steps = 0, 0, 0
	minManhattan := coord{
		x: math.MaxInt32,
		y: math.MaxInt32,
	}
	minSteps := math.MaxInt64
	for _, command := range strings.Split(line, ","){
		magnitude, _ := strconv.Atoi(command[1:])
		switch command[:1] {
		case "L":
			for i := 0; i < magnitude; i ++ {
				x--
				steps++
				c := coord{x, y}
				if grid[c] != 0 {
					if c.manhattanCloserThan(minManhattan) {
						minManhattan = c
					}
					minSteps = rmath.Min(minSteps, steps+grid[c])
				}
			}
		case "R":
			for i := 0; i < magnitude; i ++ {
				x++
				steps++
				c := coord{x, y}
				if grid[c] != 0 {
					if c.manhattanCloserThan(minManhattan) {
						minManhattan = c
					}
					minSteps = rmath.Min(minSteps, steps+grid[c])
				}
			}
		case "U":
			for i := 0; i < magnitude; i ++ {
				y++
				steps++
				c := coord{x, y}
				if grid[c] != 0 {
					if c.manhattanCloserThan(minManhattan) {
						minManhattan = c
					}
					minSteps = rmath.Min(minSteps, steps+grid[c])
				}
			}
		case "D":
			for i := 0; i < magnitude; i ++ {
				y--
				steps++
				c := coord{x, y}
				if grid[c] != 0 {
					if c.manhattanCloserThan(minManhattan) {
						minManhattan = c
					}
					minSteps = rmath.Min(minSteps, steps+grid[c])
				}
			}
		}
	}
	fmt.Printf("manhattan distance: %v\n", minManhattan)
	fmt.Printf("step distance: %v\n", minSteps)
}


func (c coord) manhattanCloserThan(other coord) bool {
	return abs(c.x) + abs(c.y) <  abs(other.x) + abs(other.y)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}