package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
		for hoffset, _ := range matrix[voffset] {
			if isGuard(matrix, voffset, hoffset) {
				//log.Printf("DEBUG: Found guard '%c' at %d/%d\n",
				//	matrix[voffset][hoffset], voffset, hoffset)
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

func directionHasObstacle(matrix []string, voffset int, hoffset int, dir int) bool {
	for moveDir(matrix, &voffset, &hoffset, dir) {
		if matrix[voffset][hoffset] == '#' {
			return true
		}
	}
	// we went outside map
	return false
}

func checkCanStartLoop(matrix []string, voffset int, hoffset int, dir int) bool {
	switch dir {
	case UP:
		// check right
		if hoffset+1 < len(matrix[voffset])-1 &&
			(matrix[voffset][hoffset+1] == '-' ||
			matrix[voffset][hoffset+1] == '+') &&
			directionHasObstacle(matrix, voffset, hoffset, RIGHT) {
			return true
		}
	case DOWN:
		// check left
		if hoffset-1 > 0 &&
			(matrix[voffset][hoffset-1] == '-' ||
			matrix[voffset][hoffset-1] == '+') &&
			directionHasObstacle(matrix, voffset, hoffset, LEFT) {
			return true
		}
	case LEFT:
		// check up
		if voffset-1 > 0 &&
			(matrix[voffset-1][hoffset] == '|' ||
			matrix[voffset-1][hoffset] == '+') &&
			directionHasObstacle(matrix, voffset, hoffset, UP) {
			return true
		}
	case RIGHT:
		// check down
		if voffset+1 < len(matrix)-1 &&
			(matrix[voffset+1][hoffset] == '|' ||
			matrix[voffset+1][hoffset] == '+') &&
			directionHasObstacle(matrix, voffset, hoffset, DOWN) {
			return true
		}
	default:
		panic("Unknown direction when checkinf for start loop")
	}

	return false
}

func getNextPos(matrix []string, voffset int, hoffset int, dir int) (int,int) {
	nextv := voffset
	nexth := hoffset
	nextInMap := moveDir(matrix, &nextv, &nexth, dir)
	if !nextInMap {
		panic("What?!")
	}

	return nextv, nexth
}

func markObstacleAtNextMove(matrixPtr *[]string, voffset int, hoffset int, dir int) {
	matrix := *matrixPtr

	nextv, nexth := getNextPos(matrix, voffset, hoffset, dir)

	matrix[nextv] = matrix[nextv][:nexth] + "O" + matrix[nextv][nexth+1:]
}

func moveGuard(matrixPtr *[]string, voffset *int, hoffset *int, dir *int, possibleObstaclePositionsPtr *[]pos) bool {
	matrix := *matrixPtr
	possibleObstaclePositions := *possibleObstaclePositionsPtr

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
		nextv, nexth := getNextPos(matrix, *voffset, *hoffset, *dir)
		if !slices.Contains(possibleObstaclePositions, pos{nextv, nexth}) {
			markObstacleAtNextMove(&matrix, *voffset, *hoffset, *dir)

			*possibleObstaclePositionsPtr = append(possibleObstaclePositions, pos{nextv, nexth})

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

func printMatrix(matrix []string) {
	fmt.Println("")
	for _, v := range matrix {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("")
}

type pos struct {
	vert int
	hor int
}

func main() {
	var origmatrix []string
	var possibleObstaclePositions []pos

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		origmatrix = append(origmatrix, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	// walk and calculate steps.
	steps := 0
	uniquePlaces := 1

	for {
		numObstacles := len(possibleObstaclePositions)

		matrix := make([]string, len(origmatrix))
		copy(matrix, origmatrix)

		voffset, hoffset := findGuard(origmatrix)
		currentDir := checkCurrentDirection(origmatrix, voffset, hoffset)

		startv := voffset
		starth := hoffset
		startGuardToken := origmatrix[startv][starth]

		for moveGuard(&matrix, &voffset, &hoffset, &currentDir, &possibleObstaclePositions) {
			//log.Printf("DEBUG: currently at %d/%d (dir %d)\n",
			//	voffset, hoffset, currentDir)
			steps += 1
			if matrix[voffset][hoffset] == '.' {
				uniquePlaces += 1
			}
		}

		// restore starting position to guard symbol
		matrix[startv] = matrix[startv][:starth] + string(startGuardToken) + matrix[startv][starth+1:]

		printMatrix(matrix)

		if len(possibleObstaclePositions) == numObstacles {
			log.Printf("DEBUG: numObstacles:%d\n", numObstacles)
			break
		}

	}

	fmt.Printf("Steps: %d, unique: %d\n", steps, uniquePlaces)

}

