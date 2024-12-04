package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func findMul(s string) int {
	return strings.Index(s, "mul(")
}

func trimDigit(s string) string {
	len := 0
	for _, char := range s {
		if !unicode.IsDigit(char) {
			break
		}
		len += 1
	}

	return s[:len]
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	mulSum := 0

	idx := findMul(string(data))
	for pos := 0; idx >= 0; idx = findMul(string(data[pos:])) {
		idx += pos

		//log.Printf("DEBUG: idx = %d\n", idx)
		// find closing parenthesis (or first corrupted data)
		end := 4
		for i, val := range string(data[idx+4:]) {
			if unicode.IsDigit(val) || val == ',' {
				end += 1
			} else {
				break
			}

			if i > 7 { // 1-3 digits *2 + ','
				break
			}
		}
		//end := strings.Index(string(data[idx:]),")")

		if string(data[idx+end:idx+end+1]) == ")" {
			end += 1
		} else {
			log.Println("mul(digits,digits did not end with )")
			pos = idx + end
			continue
		}
	
		mulArgs := string(data[idx:idx+end])
		log.Printf("DEBUG: mulArgs=%s\n", mulArgs[4:])

		// split on ,
		parts := strings.Split(mulArgs[4:len(mulArgs)-1], ",")

		if len(parts) != 2 {
			log.Printf("WARNING (ignoring): Not to values in mul operation: %s\n", mulArgs)
			pos = idx + end
			continue
		}

		//parts[0] = trimDigit(parts[0])
		//parts[1] = trimDigit(parts[1])

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("WARNING (ignoring): Failed to convert first value to integer")
			pos = idx + end
			continue
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("WARNING (ignoring): Failed to convert second value to integer")
			pos = idx + end
			continue
		}

		// multiply and summarize
		mulSum += x*y
		log.Printf("Multiplying %d x %d = %d (--> %d)\n", x, y, x*y, mulSum)

		// continue searching from ....
		pos = idx + end
		//log.Printf("DEBUG: pos = %d\n", pos)
	}

	fmt.Printf("Summary: %d\n", mulSum)
}
