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

func checkCanStartLoop(matrix []string, voffset int, hoffset int, dir int) bool {
	switch dir {
	case UP:
		// check right
	case DOWN:
		// check left
	case LEFT:
		// check up
		if voffset-1 > 0 && matrix[voffset-1][hoffset] == '|' {
			return true
		}
	case RIGHT:
		// check down
	default:
		panic("Unknown direction when checkinf for start loop")
	}

	return false
}

func markObstacleAtNextMove(matrixPtr *[]string, voffset int, hoffset int, dir int) {
	matrix := *matrixPtr

	nextv := voffset
	nexth := hoffset
	nextInMap := moveDir(matrix, &nextv, &nexth, dir)
	if !nextInMap {
		panic("What?!")
	}

	matrix[nextv] = matrix[nextv][:nexth] + "O" + matrix[nextv][nexth+1:]
}

func moveGuard(matrixPtr *[]string, voffset *int, hoffset *int, dir *int) bool {
	matrix := *matrixPtr

	// mark position as visited
	markToken := 'X'

	switch *dir {
	case UP:
		fallthrough
	case DOWN:
		if matrix[*voffset][*hoffset] == '.' {
			markToken = '|'
		} else {
			markToken = '+'
		}
	case LEFT:
		fallthrough
	case RIGHT:
		if matrix[*voffset][*hoffset] == '.' {
			markToken = '-'
		} else {
			markToken = '+'
		}
	default:
		panic("Unknown direction when setting markToken")
	}

	if checkCanStartLoop(matrix, *voffset, *hoffset, *dir) {
		// FIXME: record found loops and check if this is already found
		alreadyFound := false
		if !alreadyFound {
			markObstacleAtNextMove(&matrix, *voffset, *hoffset, *dir)

			return false
		}
	}

	for canMove := !needsTurn(matrix, *voffset, *hoffset, *dir); !canMove;
	canMove = !needsTurn(matrix, *voffset, *hoffset, *dir) {
		// turn right
		*dir += 1
		if *dir >= DirectionMax {
			*dir = 0
		}
		markToken = '+'

	}

	//matrix[*voffset][*hoffset] = markToken
	matrix[*voffset] = matrix[*voffset][:*hoffset] + string(markToken) + matrix[*voffset][*hoffset+1:]

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
	startv := voffset
	starth := hoffset
	startGuardToken := matrix[startv][starth]
	steps := 0
	uniquePlaces := 1

	for moveGuard(&matrix, &voffset, &hoffset, &currentDir) {
		log.Printf("DEBUG: currently at %d/%d (dir %d)\n",
			voffset, hoffset, currentDir)
		steps += 1
		if matrix[voffset][hoffset] == '.' {
			uniquePlaces += 1
		}
	}

	// restore starting position to guard symbol
	matrix[startv] = matrix[startv][:starth] + string(startGuardToken) + matrix[startv][starth+1:]

	fmt.Println("")
	for _, v := range matrix {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("")

	fmt.Printf("Steps: %d, unique: %d\n", steps, uniquePlaces)

}

