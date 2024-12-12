package main

import (
	"fmt"
	"strconv"
	"os"
)

const (
	FILE = iota
	FREE
)

func defragFs(fs []int) []int {
	y := len(fs)-1
	fsd := make([]int, len(fs))

	for i:= 0; i<len(fs); i += 1 {
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

	return  fsd
}

func checksumFs(fs []int) uint64 {
	sum := uint64(0)
	pos := 0
	for _, val := range fs {
		if val < 0 {
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
	for i := len(input)-1 ; input[i] == '\n' ; i -= 1 {
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
			for i := 0 ; i < n; i += 1 {
				filesystem = append(filesystem, fileid)
			}
			fileid+=1
			state = FREE
		case FREE:
			for i := 0 ; i < n; i += 1 {
				filesystem = append(filesystem, -1)
			}
			state = FILE
		default:
			panic("not supposed to happen")
		}
	}
	fsd := defragFs(filesystem)

	checksumFs := checksumFs(fsd)
	fmt.Printf("Checksum: %v\n", checksumFs)
}
