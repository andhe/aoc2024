package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftCol := make([]int, 0, 100)
	rightCol := make([]int, 0, 100)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Line does not consist of 2 values: %s\n", line)
		}

		if len(parts[0]) != 5 {
			log.Fatalf("Left len() != 5 (%d ---> %s)", len(parts[0]), parts[0])
		}
		if len(parts[1]) != 5 {
			log.Fatalf("Right len() != 5 (%d ---> %s)", len(parts[1]), parts[1])
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to convert left value to int")
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed to convert right value to int")
		}

		//fmt.Printf("DEBUG: Left: %s == %d\n", parts[0], left)
		//fmt.Printf("DEBUG: Right: %s == %d\n", parts[1], right)

		leftCol = append(leftCol, left)
		rightCol = append(rightCol, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error!")
	}

	if len(leftCol) != len(rightCol) {
		log.Fatalf("right and left column length not the same!")
	}

	//fmt.Println("### Left")
	//fmt.Println(leftCol)
	//fmt.Println("### Right")
	//fmt.Println(rightCol)

	sort.Ints(leftCol)
	sort.Ints(rightCol)

	/*
	fmt.Println("### Left")
	fmt.Println(leftCol)
	fmt.Println("### Right")
	fmt.Println(rightCol)
	*/

	diffCol := make([]int, 0, len(leftCol))
	distSum := uint64(0)

	simCol := make([]int, len(leftCol))
	simSum := uint64(0)

	for idx, _ := range leftCol {
		var distance int
		//fmt.Println(idx)
		if rightCol[idx] > leftCol[idx] {
			distance = rightCol[idx] - leftCol[idx]
		} else {
			distance = leftCol[idx] - rightCol[idx]
		}
		//fmt.Printf("Left: %d, Right: %d, diff: %d\n", leftCol[idx], rightCol[idx], diff)
		diffCol = append(diffCol, distance)
		distSum += uint64(distance)

		// -----------------------------------------

		// we assume the rightCol is sorted, so we can just iterate until
		// the value is no longer the same as the one we started with to
		// count the number of times a certain number exists in rightCol.
		for i := slices.Index(rightCol, leftCol[idx]); i > 0 && rightCol[i] == leftCol[idx]; i=i+1 {
			simCol[idx] += leftCol[idx]
		}

		/*
		if simCol[idx] != 0 {
			fmt.Printf("DEBUG: similar %d (idx %d) found %d times ==> %d\n",
				leftCol[idx], idx, simCol[idx] / leftCol[idx], simCol[idx])
			if simCol[idx] > 1000000 {
				fmt.Println("BANG!")
			}
		}
		*/
			

		simSum += uint64(simCol[idx])
	}

	if len(leftCol) != len(diffCol) {
		log.Fatalf("Diff column not the same length as left column\n")
	}

	fmt.Printf("Sum of distances: %d\n", distSum)
	fmt.Printf("Sum of similarities: %d\n", simSum)

	/*
	for idx, val := range leftCol {
		if idx+1 < len(leftCol) {
			if (val == leftCol[idx+1]) {
				fmt.Printf("Left[%d] == Left[%d] == %d\n", idx, idx+1, val)
			}
		}
	}
	*/
}
