package main

import (
	"fmt"
	"github.com/raidancampbell/libraidan/pkg/rstrings"
	"regexp"
	"strconv"
	"strings"
)

var p = regexp.MustCompile("(?P<min>\\d+)-(?P<max>\\d+) (?P<char>[a-z]): (?P<val>.+)")

func main() {
	str, err := rstrings.FileToString("day2/input")
	if err != nil {
		fmt.Printf("Error encountered reading file.  Ensure working directory is repository root: %v\n", err)
	}

	valid := 0
	valid2 := 0
	for _, line := range strings.Split(str, "\n") {
		params := getParams(line, p)
		min, _ := strconv.Atoi(params["min"])
		max, _ := strconv.Atoi(params["max"])
		if isValidPass1(min, max, params["char"], params["val"]) {
			valid++
		}
		if isValidPass2(min, max, params["char"][0], params["val"]) {
			valid2++
		}

	}
	fmt.Println(valid)
	fmt.Println(valid2)
}

func isValidPass1(min, max int, char, val string) bool {
	c := strings.Count(val, char)
	return c >= min && c <= max
}

func isValidPass2(loc1, loc2 int, char uint8, val string) bool {
	loc1--
	loc2--
	if val[loc1] == char {
		return val[loc2] != char
	}
	return val[loc2] == char
}


func getParams(haystack string, pattern *regexp.Regexp) (paramsMap map[string]string) {

	match := pattern.FindStringSubmatch(haystack)

	paramsMap = make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}