package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

func appendPart(old string, part string) string {
	if old == "" {
		return part
	} else {
		return fmt.Sprintf("%s %s", old, part)
	}
}

func trim(in string) string {
	out := strings.TrimLeft(in, "0")
	if out == "" {
		out = "0"
	}

	return out
}

func blink(in string) string {
	out := ""
	parts := strings.Split(in, " ")

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}

		if n == 0 {
			out = appendPart(out, "1")
		} else if len(p)%2 == 0 {
			out = appendPart(out, trim(p[:len(p)/2])) + " " + trim(p[len(p)/2:])
		} else {

			out = appendPart(out, fmt.Sprintf("%d", n*2024))
		}
	}

	return out
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	d := strings.Trim(string(data), "\r\n")

	fmt.Println("Initial data:", d)

	// blink 25 times
	for i:= 0; i < 25; i += 1 {
		d = blink(d)
		//fmt.Println(d)
	}

	parts := strings.Split(d, " ")
	stones := len(parts)
	fmt.Println("Stones:", stones)
}
