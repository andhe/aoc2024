package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x int
	y int
}

func main() {
	var matrix []string
	var plots []map[rune][]pos

	// populate matrix
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// find plots
	
}
