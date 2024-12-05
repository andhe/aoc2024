package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
)

const (
	Up = iota
	Down
	Left
	Right
	SE
	SW
	NE
	NW
)

func moveDir(matrix []string, voffset *int, hoffset *int, dir int) bool {
	switch dir {
	case Up:
		if *voffset == 0 {
			return false
		}
		*voffset -= 1
	case Down:
		if *voffset >= len(matrix)-1 {
			return false
		}
		*voffset += 1
	case Left:
		if *hoffset <= 0 {
			return false
		}
		*hoffset -= 1
	case Right:
		if *hoffset >= len(matrix[*voffset])-1 {
			return false
		}
		*hoffset += 1
	
	default:
		log.Println("move in unknown direction")
		return false
	}

	return true
}

func checkMAS(matrix []string, voffset int, hoffset int) bool {
	if voffset-1 < 0 || hoffset-1 < 0 {
		return false
	}
	if voffset+1 > len(matrix)-1 || hoffset+1 > len(matrix[0])-1 {
		return false
	}

	NW := matrix[voffset-1][hoffset-1]
	SE := matrix[voffset+1][hoffset+1]
	NE := matrix[voffset-1][hoffset+1]
	SW := matrix[voffset+1][hoffset-1]

	if NW == 'M' && SE == 'S' || NW == 'S' && SE == 'M' {
		if NE == 'M' && SW == 'S' || NE == 'S' && SW == 'M' {
			return true
		}
	}
	return false
}

func main() {
	var matrix []string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	foundSum := 0

	for idx,input := range matrix {
		pos := 0
		for i := strings.IndexByte(input[pos:], 'A'); i >= 0; i = strings.IndexByte(input[pos:], 'A') {
			pos += i
			found := checkMAS(matrix, idx, pos)

			if found {
				foundSum += 1

				log.Printf("DEBUG: Found X-MAS with %c at %d/%d\n",
					"A", idx, pos)
			}

			pos += 1
		}
	}

	fmt.Printf("Sum: %d\n", foundSum)

}

