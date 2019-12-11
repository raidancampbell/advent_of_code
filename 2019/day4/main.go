package main

import (
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"math"
	"strconv"
	"strings"
)

/*
  Two adjacent digits are the same (like 22 in 122345).
  Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
 */
func main() {
	str, err := rstrings.FileToString("2019/day4/input")
	if err != nil {
		fmt.Printf("Error encountered reading file.  Ensure working directory is repository root: %v\n", err)
	}
	bounds := strings.Split(str, "-")
	lb, _ := strconv.Atoi(bounds[0])
	ub, _ := strconv.Atoi(bounds[1])


	// 356261-846303
	matches := 0
	p := password{lb}
	for p.val = lb; p.val < ub; {
		inc := p.check2()
		if inc == 0 {
			println(p.val)
			matches++
			inc = 1
		}

		p.val += inc
	}
	fmt.Printf("found %v matches\n", matches)

}

type password struct {
	val int
}

func (p password) check2() int {
	s :=  []byte(fmt.Sprintf("%06d", p.val))

	dubs := byte(0)
	for idx := range s {
		if idx > 0 && s[idx] < s[idx-1] {
			return 1
		}
		// first digit is occluded
		if idx == 0 {
			if s[idx+1] == s[idx] && s[idx+2] != s[idx] {
				dubs = s[idx]
			}
		} else if idx + 1 < 5 { // none are occluded
			if s[idx-1] != s[idx] && s[idx+1] == s[idx] && s[idx+2] != s[idx] {
				dubs = s[idx]
			}
		} else if idx + 1 < 6{ // last digit is occluded
			if s[idx-1] != s[idx] && s[idx+1] == s[idx] {
				dubs = s[idx]
			}
		}
	}
	if dubs == 0 {
		return 1
	}
	return 0
}

func (p password) check() int {
	s := strconv.Itoa(p.val)

	hasDubs := false
	last := int32(math.MinInt32)
	for _, c := range s {
		if c < last {
			return 1
		}
		if c == last{
			hasDubs = true
		}
		last = c
	}
	if !hasDubs{
		return 1
	}
	return 0
}