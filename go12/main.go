package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type pos struct {
	x int
	y int
}

func checkConnected(val byte, x int, y int, matrix []string, plotsPtr *[][]pos) bool {
	plots := *plotsPtr


	line := matrix[x]
	if y > 0 {
		yleftval := line[y-1]
		if yleftval == val {
			for n,s := range plots {
				if slices.Contains(s, pos{x,y-1}) {
					plots[n] = append(plots[n], pos{x,y})
					return true
				}
			}
		}
	}
	if x > 0 {
		upline := matrix[x-1]
		xupval := upline[y]
		if xupval == val {
			for n,s := range plots {
				if slices.Contains(s, pos{x-1,y}) {
					plots[n] = append(plots[n], pos{x,y})
					return true
				}
			}
		}
	}
	// no point in checking in other directions, because we only check what we've already processed
	return false
}

func main() {
	var matrix []string
	var plots [][]pos

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
	for x,line := range matrix {
		for y,val := range line {
			if !checkConnected(byte(val), x, y, matrix, &plots) {
				// new plot
				p := make([]pos, 0, 1)
				p = append(p, pos{x,y})
				plots = append(plots, p)
			}
		}
	}

	for _, p := range plots {
		fmt.Println("%s: %v\n", string(matrix[p[0].x][p[0].y]), p)
	}
	
}
