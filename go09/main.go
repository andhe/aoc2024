package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

const (
	FILE = iota
	FREE
)

func defragFs(fs []int) []int {
	y := len(fs) - 1
	fsd := make([]int, len(fs))

	for i := 0; i < len(fs); i += 1 {
		val := fs[i]

		if val >= 0 {
			fsd[i] = val
			continue
		}

		// found first free space, now find last block
		for ; y > i && fs[y] == -1; y -= 1 {
			continue
		}

		// switch places
		fsd[i] = fs[y]
		fs[y] = -1
	}

	return fsd
}

func checksumFs(fs []int) uint64 {
	sum := uint64(0)
	pos := 0
	for _, val := range fs {
		if val < 0 {
		pos += 1
			continue
		}

		sum += uint64(val * pos)
		pos += 1
	}

	return sum
}

func main() {
	var filesystem []int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to read input.txt")
	}
	for i := len(input) - 1; input[i] == '\n'; i -= 1 {
		input = input[:len(input)-1]
	}

	fileid := 0
	state := FILE

	for _, val := range input {
		n, err := strconv.Atoi(string(val))
		if err != nil {
			panic("failed to atoi value")
		}

		switch state {
		case FILE:
			for i := 0; i < n; i += 1 {
				filesystem = append(filesystem, fileid)
			}
			fileid += 1
			state = FREE
		case FREE:
			for i := 0; i < n; i += 1 {
				filesystem = append(filesystem, -1)
			}
			state = FILE
		default:
			panic("not supposed to happen")
		}
	}
	origfs := make([]int, len(filesystem))
	copy(origfs, filesystem)

	log.Printf("DEBUG: filesystem=%v\n", filesystem)
	fsd := defragFs(filesystem)
	log.Printf("DEBUG: fsd=%v\n", fsd)

	copy(filesystem, origfs)

	log.Printf("DEBUG: filesystem=%v\n", filesystem)
	fsd2 := defragFs2(filesystem)

	log.Printf("DEBUG: fsd2=%v\n", fsd2)

	checksum := checksumFs(fsd)
	checksum2 := checksumFs(fsd2)

	fmt.Printf("Checksum: %v\n", checksum)
	fmt.Printf("Checksum: %v\n", checksum2)
}

func defragFs2(fs []int) []int {
	fileid := fs[len(fs)-1]

	for fileid >= 0 {
		//log.Println("DEBUG: fileid:", fileid)

		// find location of file
		y := slices.Index(fs, fileid)

		if y < 0 {
			fileid -= 1
			continue
		}

		// measure file length
		filelen := 0
		for ; y + filelen < len(fs) && fs[y+filelen] == fileid ; filelen += 1 {
			continue
		}

		x := 0

		freelen := 0
		for x < len(fs) {
			//log.Println("DEBUG: searching for free space for", fileid, "at", x)

			// find free space
			xoff := slices.Index(fs[x:], -1)

			if xoff < 0 {
				x = -1
				break
			}
			x += xoff

			// measure free space length
			freelen = 0
			for ; x+freelen < len(fs) && fs[x+freelen] == -1 ; freelen += 1 {
				continue
			}

			if filelen <= freelen {
				break
			}

			x += freelen
		}

		if x > 0 && y > 0 && x < y {
			//log.Println("DEBUG: swapping", fileid, "at", y, "to free space at", x, "(filelen:", filelen,", freelen:", freelen, ")")
			for n := range filelen {
				fs[x+n] = fileid
				fs[y+n] = -1
			}
		}

		fileid -= 1

		//log.Printf("DEBUG: fs=%v\n", fs)
	}

	return  fs
}
