package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left *Node
	right *Node
}

func findLeafs(t *Node) []*Node {
	if t.left == nil && t.right == nil {
		var tarr []*Node
		tarr = append(tarr, t)
		return tarr
	} else {
		return append(findLeafs(t.left), findLeafs(t.right)...)
	}
}

func printTree(t *Node, indent uint) {
	if t == nil {
		return
	}

	for range indent {
		fmt.Printf("-")
	}
	fmt.Printf("%d\n", t.value)

	printTree(t.left, indent+1)
	printTree(t.right, indent+1)
}

func canAddMultiply(total int, values []int) bool {

	//log.Printf("DEBUG: total:%d, values:%v\n", total, values)

	// build tree
	t := Node{ values[0], nil, nil }

	for _, v := range values[1:] {
		// find leafs and add/multiply
		leafs := findLeafs(&t)

		for _, l := range leafs {
			nl := Node{l.value + v, nil, nil}
			nr := Node{l.value * v, nil, nil}
			l.left = &nl
			l.right = &nr
		}
	}

	//printTree(&t, 0)

	// check tree leafs if total is found
	for _, l := range findLeafs(&t) {
		if l.value == total {
			log.Printf("DEBUG: found total %d\n", total)
			return true
		}
	}

	return false
}

func isValid(input string) (int, error) {
	parts := strings.Split(input, " ")

	if len(parts) < 2 {
		log.Fatalf("Invalid input: %s\n", input)
	}

	if parts[0][len(parts[0])-1] != ':' {
		panic("first value does not end with :")
	}

	total, err := strconv.Atoi(parts[0][:len(parts[0])-1])
	if err != nil {
		panic("Failed to atoi total value")
	}

	values := make([]int, len(parts)-1)
	for idx, p := range parts[1:] {
		i, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("Failed atoi of parts[%d]=%s\n", idx+1, parts[idx+1])
		}

		values[idx] = i
	}

	if canAddMultiply(total, values) {
		return total, nil
	} else {
		return 0, nil
	}

}

func main() {
	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := isValid(line)
		if err == nil {
			sum += i
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("sum: %d\n", sum)
}
