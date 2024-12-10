package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
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

func wrapVert(vertval *int, maxPos pos) bool {
	if *vertval < 0 {
		*vertval = maxPos.v + *vertval
		return true
	} else if *vertval >= maxPos.v {
		*vertval -= maxPos.v
		return true
	}

	return  false
}


func wrapHor(horval *int, maxPos pos) bool {
	if *horval < 0 {
		*horval = maxPos.h + *horval
		return true
	} else if *horval >= maxPos.v {
		*horval -= maxPos.h
		return true
	}

	return false
}

func getAntiForPositions(ps []pos, lines []string) []pos {
	var antinodes []pos
	// for each pair of positions, get antinodes
	for _,p := range getPairs(ps) {
		log.Printf("DEBUG: getting antinodes for pos %v\n", p)

		a := p[0]
		b := p[1]

		vdiff := a.v - b.v
		hdiff := a.h - b.h

		if vdiff < 0 {
			vdiff *= -1
		}
		if hdiff < 0 {
			hdiff *= -1
		}

		var a1v, b1v, a1h, b1h int
		if a.v < b.v {
			a1v = a.v-vdiff
			b1v = b.v+vdiff
		} else {
			a1v = a.v+vdiff
			b1v = b.v-vdiff
		}
		if a.h < b.h {
			a1h = a.h - hdiff
			b1h = b.h + hdiff
		} else {
			a1h = a.h + hdiff
			b1h = b.h - hdiff
		}

		maxH := len(lines[0])
		maxV := len(lines)
		maxPos := pos{ v: maxV, h: maxH }

		apos := pos{ v: a1v, h: a1h }
		bpos := pos{ v:b1v, h: b1h }
		if !wrapVert(&a1v, maxPos) && !wrapHor(&a1h, maxPos) {
			antinodes = append(antinodes, apos)
		}
		if !wrapVert(&b1v, maxPos) && !wrapHor(&b1h, maxPos) {
			antinodes = append(antinodes, bpos)
		}

		log.Printf("DEBUG: antinodes at %v and %v\n",
			apos, bpos)

	}

	return antinodes
}

func drawMap(m map[rune][]pos, antip map[rune][]pos, maxPos pos) []string {
	matrix := make([]string, maxPos.v)

	// fill out the map with blanks
	for i := range maxPos.v {
		matrix[i] = strings.Repeat(".", maxPos.h)
	}

	// fill in antipositions
	// needs to be drawn first, because can be covered by markers
	for key := range antip {
		for i := range antip[key] {
			p := antip[key][i]
			matrix[p.v] = matrix[p.v][:p.h] + "#" + matrix[p.v][p.h+1:]
		}
	}


	// fill in map markers
	for key := range m {
		for i := range m[key] {
			p := m[key][i]
			matrix[p.v] = matrix[p.v][:p.h] + string(key) + matrix[p.v][p.h+1:]
		}
	}



	return matrix
}

func uniquePositions(antip map[rune][]pos) int {
	var pm []pos
	for key := range antip {
		for i := range antip[key] {
			p := antip[key][i]
			if !slices.Contains(pm, p) {
				pm = append(pm, p)
			//} else {
			//	log.Printf("DEBUG: duplicate antipos: %v\n", p)
			}
		}
	}
	log.Printf("DEBUG: Unique antinode pos: %v\n", pm)

	return len(pm)
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
		log.Printf("DEBUG: %d: %s\n", voff, line)

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
		antip[key] = getAntiForPositions(m[key], lines)
	}

	out := drawMap(m, antip, pos{len(lines), len(lines[0]) })

	for _, l := range out {
		fmt.Println(l)
	}

	antiSum := uniquePositions(antip)
	fmt.Printf("Sum antinodes: %d\n", antiSum)
}
