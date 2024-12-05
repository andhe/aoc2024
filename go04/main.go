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

func checkWordDir(matrix []string, voffset int, hoffset int, word string, dir int) int {
	for idx, val := range word {
		if val != rune(matrix[voffset][hoffset]) {
			return 0
		}

		if idx == len(word)-1 {
			// don't need to move more, we're done. Entire word now checked.
			return 1
		}

		switch dir {
		case Up:
			fallthrough
		case Down:
			fallthrough
		case Left:
			fallthrough
		case Right:
			if !moveDir(matrix, &voffset, &hoffset, dir) {
				return 0
			}
		case SE:
			if !moveDir(matrix, &voffset, &hoffset, Down) || !moveDir(matrix, &voffset, &hoffset, Right) {
				return 0
			}
		case SW:
			if !moveDir(matrix, &voffset, &hoffset, Down) || !moveDir(matrix, &voffset, &hoffset, Left) {
				return 0
			}
		case NW:
			log.Printf("DEBUG: checked NW: %d/%d\n", hoffset, voffset)
			if !moveDir(matrix, &voffset, &hoffset, Up) || !moveDir(matrix, &voffset, &hoffset, Left) {
				return 0
			}
		case NE:
			if !moveDir(matrix, &voffset, &hoffset, Up) || !moveDir(matrix, &voffset, &hoffset, Right) {
				return 0
			}
		default:
			panic("Unhandled direction")
		}
	}

	log.Printf("Found word in dir: %d\n", dir)

	return 1
}

func checkWordAllDirections(matrix []string, voffset int, hoffset int, word string) int {
	found := 0

	found += checkWordDir(matrix, voffset, hoffset, word, Up)
	found += checkWordDir(matrix, voffset, hoffset, word, Down)
	found += checkWordDir(matrix, voffset, hoffset, word, Left)
	found += checkWordDir(matrix, voffset, hoffset, word, Right)
	found += checkWordDir(matrix, voffset, hoffset, word, SE)
	found += checkWordDir(matrix, voffset, hoffset, word, SW)
	found += checkWordDir(matrix, voffset, hoffset, word, NE)
	found += checkWordDir(matrix, voffset, hoffset, word, NW)

	return found
}

func main() {
	word := "XMAS"
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
		for i := strings.IndexByte(input[pos:], word[0]); i >= 0; i = strings.IndexByte(input[pos:], word[0]) {
			pos += i
			found := checkWordAllDirections(matrix, idx, pos, word)

			if found > 0 {
				foundSum += found

				log.Printf("DEBUG: Found %c at %d/%d - word '%s' in all directions: %d\n",
					word[0], idx, pos, word, found)
			}

			pos += 1
		}
	}

	fmt.Printf("Sum: %d\n", foundSum)

}

