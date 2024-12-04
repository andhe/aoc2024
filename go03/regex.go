package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mul(s string) int {
	// input: mul(123,456)
	parts := strings.Split(s[4:len(s)-1], ",")
	if len(parts) != 2 {
		log.Fatalf("Not two parts!")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("x failed atoi")
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("y failed atoi")
	}

	return x * y
}

func main() {
	// Input string
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Compile the regex pattern
	mulpattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(mulpattern)

	// Find all matches
	muls := re.FindAllString(string(input), -1)
	mulindices := re.FindAllStringIndex(string(input), -1)

	dopattern := `do\(\)`
	re = regexp.MustCompile(dopattern)
	doindices := re.FindAllStringIndex(string(input), -1)

	dontpattern := `don\'t\(\)`
	re = regexp.MustCompile(dontpattern)
	dontindices := re.FindAllStringIndex(string(input), -1)

	mulEnabled := true
	sum := 0

	for len(mulindices) > 0 {
		donti := -1
		doi := -1

		if len(dontindices) > 0 && dontindices[0][0] < mulindices[0][0] {
			donti = dontindices[0][0]
		}
		if len(doindices) > 0 && doindices[0][0] < mulindices[0][0] {
			doi = doindices[0][0]
		}

		if donti >= 0 && doi >= 0 {
			if donti < doi {
				dontindices = dontindices[1:]
				mulEnabled = false
			} else {
				doindices = doindices[1:]
				mulEnabled = true
			}
		} else if donti < 0 && doi < 0 {
			// ok
		} else {
			if donti < 0 {
				doindices = doindices[1:]
				mulEnabled = true
			} else {
				dontindices = dontindices[1:]
				mulEnabled = false
			}
		}

		if mulEnabled {
			sum += mul(muls[0])
		}

		muls = muls[1:]
		mulindices = mulindices[1:]
	}

	fmt.Printf("Sum: %d\n", sum)
}

