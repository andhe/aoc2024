package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	v int
	h int
}




func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func getPairs(ps []pos) [][2]pos {
	var pairs [][2]pos
	for i := 0; i < len(ps); i += 1 {
		for j := 0; j < len(ps); j += 1 {
			if i == j {
				continue
			}

			pairs = append(pairs, [2]pos{ps[i], ps[j]})
		}
	}

	return pairs
}

func getAntiForPositions(ps []pos) []pos {
	// for each pair of positions, get antinodes
	for _,p := range getPairs(ps) {
		log.Printf("DEBUG: getting antinodes for pos %v\n", p)
	}

	return []pos{}
}

func main() {
	m := make(map[rune][]pos)
	antip := make(map[rune][]pos)

	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for voff, line := range lines {
		log.Printf("DEBUG: %d: %s\n", voff+1, line)

		for hoff, c := range line {
			if c != '.' {
				p := pos{voff, hoff}

				log.Printf("DEBUG: found %c at %v\n", c, p)

				if val, ok := m[c]; ok {
					m[c] = append(val, p)
				} else {
					m[c] = []pos{p}
				}
			}
		}
	}

	for key := range m {
		log.Printf("DEBUG: getting antipos for key: %v\n", key)
		antip[key] = getAntiForPositions(m[key])
	}
}
