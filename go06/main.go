package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
	DirectionMax
)

func isGuard(matrix []string, voffset int, hoffset int) bool {
	return checkCurrentDirection(matrix, voffset, hoffset) >= 0
}

func findGuard(matrix []string) (int, int) {
	for voffset := 0; voffset < len(matrix)-1; voffset += 1 {
		for hoffset, v := range matrix[voffset] {
			if isGuard(matrix, voffset, hoffset) {
				log.Printf("DEBUG: Found guard '%c' at %d/%d\n",
					v, voffset, hoffset)
				return voffset, hoffset
			}
		}
	}

	panic("No guard found! There should be one!")
}

func checkCurrentDirection(matrix []string, voffset int, hoffset int) int {
	switch matrix[voffset][hoffset] {
	case '^':
		return UP
	case '>':
		return RIGHT
	case '<':
		return LEFT
	case 'v':
		return DOWN
	default:
		return -1
	}
}

func moveDir(matrix []string, voffset *int, hoffset *int, dir int) bool {
	switch dir {
	case UP:
		if *voffset == 0 {
			return false
		}
		*voffset -= 1
	case DOWN:
		if *voffset >= len(matrix)-1 {
			return false
		}
		*voffset += 1
	case LEFT:
		if *hoffset <= 0 {
			return false
		}
		*hoffset -= 1
	case RIGHT:
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

func needsTurn(matrix []string, voffset int, hoffset int, dir int) bool {
	// check which would be next position
	nextv := voffset
	nexth := hoffset
	nextInMap := moveDir(matrix, &nextv, &nexth, dir)

	if !nextInMap {
		// we're done
		return false
	}

	// turn if next position is an obstacle
	// else move there.
	if matrix[nextv][nexth] == '#' {
		return true
	} else {
		return false
	}


}

func moveGuard(matrixPtr *[]string, voffset *int, hoffset *int, dir *int) bool {
	matrix := *matrixPtr

	// mark position as visited
	//matrix[*voffset][*hoffset] = 'X'
	matrix[*voffset] = matrix[*voffset][:*hoffset] + "X" + matrix[*voffset][*hoffset+1:]

	for canMove := !needsTurn(matrix, *voffset, *hoffset, *dir); !canMove;
	canMove = !needsTurn(matrix, *voffset, *hoffset, *dir) {
		// turn right
		*dir += 1
		if *dir >= DirectionMax {
			*dir = 0
		}

	}

	inMap := moveDir(matrix, voffset, hoffset, *dir)
	if !inMap {
		return false
	}
	return true

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

	// walk and calculate steps.
	voffset, hoffset := findGuard(matrix)
	currentDir := checkCurrentDirection(matrix, voffset, hoffset)
	steps := 0
	uniquePlaces := 1

	for moveGuard(&matrix, &voffset, &hoffset, &currentDir) {
		log.Printf("DEBUG: currently at %d/%d (dir %d)\n",
			voffset, hoffset, currentDir)
		steps += 1
		if matrix[voffset][hoffset] != 'X' {
			uniquePlaces += 1
		}
	}

	fmt.Printf("Steps: %d, unique: %d\n", steps, uniquePlaces)

}

