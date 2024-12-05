package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"slices"
	"strings"
	"strconv"
)

type rule struct {
	x int
	y int
}

type update struct {
	pages []int
}

func main() {
	var rules []rule
	var updates []update
	var correctUpdates []update

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	processingRules := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			processingRules = false
			continue
		}

		if processingRules {
			var x int
			var y int

			n,err := fmt.Sscanf(line, "%d|%d", &x, &y)
			if err != nil {
				panic(err)
			}

			if n != 2 {
				log.Fatalf("Rule did not consist of 2 numbers")
			}

			rules = append(rules, rule{x,y})

		} else {
			parts := strings.Split(line, ",")
			var u update

			for _, val := range parts {
				n, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}

				u.pages = append(u.pages, n)
			}

			updates = append(updates, u)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	for idx, u := range updates {
		if !checkUpdateAgainstRules(u, rules) {
			log.Printf("DEBUG: Update %d failed rules check\n", idx)
		} else {
			correctUpdates = append(correctUpdates, u)
		}
	}

	midSum := 0
	for _, u := range correctUpdates {
		mid := u.pages[(len(u.pages))/2]
		log.Printf("DEBUG: mid==%d\n", mid)
		midSum += mid
	}

	log.Printf("Sum: %d\n", midSum)
}


func checkUpdateAgainstRules(u update, rules []rule) bool {
	for idx, r := range rules {
		xoff := slices.Index(u.pages, r.x)
		yoff := slices.Index(u.pages, r.y)

		// check if current rule applies, if not continue.
		if xoff < 0 || yoff < 0 {
			continue
		}

		// verify x comes before y in u.pages
		if xoff >= yoff {
			log.Printf("DEBUG: rule %d (%v) violated in %v\n", idx, r, u)
			return false
		}
	}

	return true
}
