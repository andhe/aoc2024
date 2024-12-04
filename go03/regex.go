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
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(string(input), -1)

	sum := 0
	// Print matches
	for _, match := range matches {
		sum += mul(match)
	}

	fmt.Printf("Sum: %d\n", sum)
}

