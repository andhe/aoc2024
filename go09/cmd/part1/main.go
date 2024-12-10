package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

const (
	FILE = iota
	FREE
)

func defrag(line string) string {
	y := len(line)-1

	data := line

	for idx, val := range line {
		if val != '.' {
			continue
		}

		// found first free space

		// now find last file block
		for ; line[y] == '.' && y >= 0 ; y -= 1 {
			continue
		}

		if y < idx {
			break
		}

		// move block
		/*
		fmt.Println("---------------------------------------")
		fmt.Printf("DEBUG: %s\n", data[:idx])
		fmt.Printf("DEBUG: %s\n", string(line[y]))
		fmt.Printf("DEBUG: %s\n", data[idx+1:y])
		fmt.Printf("DEBUG: %s\n", ".")
		fmt.Printf("DEBUG: %s\n", data[y+1:])
		fmt.Println("---------------------------------------")
		*/

		data = data[:idx] + string(line[y]) + data[idx+1:y] + "." + data[y+1:]
		y -= 1

		//fmt.Printf("%s\n", data)

	}

	return data
}

func checksum(data string) uint64 {
	sum := uint64(0)
	fileid := 0
	for _, val := range data {
		if val == '.' {
			continue
		}

		n, err := strconv.Atoi(string(val))
		if err != nil {
			panic("failed to atoi value")
		}

		sum += uint64(fileid * n)
		fileid += 1
	}

	return sum

}

func main() {
//	input := "2333133121414131402"
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to read input.txt")
	}
	for i := len(input)-1 ; input[i] == '\n' ; i -= 1 {
		input = input[:len(input)-1]
	}

	fileid := 0
	state := FILE

	output := ""

	for _, val := range input {
		//fmt.Printf("DEBUG: val=%c (%d)\n", val, val)
		n, err := strconv.Atoi(string(val))
		if err != nil {
			panic("failed to atoi value")
		}

		switch state {
		case FILE:
			output += fmt.Sprintf("%s", strings.Repeat(strconv.Itoa(fileid), n))
			fileid+=1
			state = FREE
		case FREE:
			output += fmt.Sprintf("%s", strings.Repeat(".", n))
			state = FILE
		default:
			panic("not supposed to happen")
		}
	}
	fmt.Printf("%s\n", output)
	d := defrag(output)
	fmt.Println(d)

	checksum := checksum(d)
	fmt.Printf("Checksum: %v\n", checksum)
}
