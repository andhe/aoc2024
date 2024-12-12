package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type pos struct {
	x int
	y int
}


func findTrailHeads(matrix [][]int) []pos {
	var h []pos

	for y, row := range matrix {
		for x, val := range row {
			if val == 0 {
				h = append(h, pos{x,y})
			}
		}
	}

	return h
}

func read2D(filename string) [][]int {
	var result [][]int

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, val := range line {
			num, err := strconv.Atoi(string(val))
			if err != nil {
				num = -1
			}

			row = append(row, num)
		}
		result = append(result, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func findNextPos(matrix [][]int, p pos) []pos {
	var np []pos

	// check left
	if p.x > 0 && matrix[p.y][p.x-1] == matrix[p.y][p.x]+1 {
		np = append(np, pos{p.x-1, p.y})
	}
	// check right
	if p.x < len(matrix[p.y])-1 && matrix[p.y][p.x+1] == matrix[p.y][p.x]+1 {
		np = append(np, pos{p.x+1, p.y})
	}
	// check up
	if p.y > 0 && matrix[p.y-1][p.x] == matrix[p.y][p.x]+1 {
		np = append(np, pos{p.x, p.y-1})
	}
	// check down
	if p.y < len(matrix)-1 && matrix[p.y+1][p.x] == matrix[p.y][p.x]+1 {
		np = append(np, pos{p.x, p.y+1})
	}

	return np
}


func findRemainingTrails(matrix [][]int, trails [][]pos, p pos, trail []pos) [][]pos {
	var newtrails [][]pos
	t := make([]pos, len(trail))

	//log.Println("DEBUG: current trail:", trail)

	copy(t, trail)
	t = append(t, p)

	//log.Println("DEBUG: trail including current position:", t)

	if matrix[p.y][p.x] == 9 {
		// found complete trail
		//log.Println("DEBUG: Found complete trail:", t)
		newtrails = append(newtrails, t)
		return newtrails
	}


	nextPositions := findNextPos(matrix, p)
	//log.Println("DEBUG: next positions:", nextPositions)
	for _, np := range nextPositions {
		newtrails = append(newtrails, findRemainingTrails(matrix, trails, np, t)...)
	}

	return newtrails
}


func findTrailsFromHeads(matrix [][]int, h []pos) [][]pos{
	var trails [][]pos
	var trail []pos

	for _, p := range h {
		trails = append(trails, findRemainingTrails(matrix, trails, p, trail)...)
	}

	return trails
}


func findTrails(matrix [][]int) [][]pos {
	trailHeads := findTrailHeads(matrix)
	log.Printf("DEBUG: trailHeads=%v\n", trailHeads)

	trails := findTrailsFromHeads(matrix, trailHeads)
	return trails
}

func main() {
	matrix := read2D("input.txt")

	fmt.Println("-------------------------------------")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("-------------------------------------")

	trails := findTrails(matrix)

	for _, t := range trails {
		log.Println(t)
		/*
		for _, p := range t {
			log.Print(matrix[p.y][p.x])
		}
		log.Println()
		*/
	}

	trailsum := calcScore(trails)
	fmt.Println("Sum:", trailsum)
}

func calcScore(trails [][]pos) int {
	headTails := make(map[pos][]pos)
	sum := 0

	// map start positions to finish (removing duplicates)
	for _, t := range trails {
		tails := headTails[t[0]]

		if slices.Index(tails, t[len(t)-1]) < 0 {
			tails = append(tails, t[len(t)-1])
		}

		headTails[t[0]] = tails
	}

	for key, tails := range headTails {
		score := len(tails)
		log.Printf("DEBUG: head %d/%d has %d", key.x, key.y, score)
		sum += score
	}

	return sum
}
