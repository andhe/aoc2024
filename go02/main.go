package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func intAbs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func checkReportValid(reportValues []string) bool {
	reportProblems := 0
	direction := 0
	checkDirection := true
	var checkedReportValues []int

	for idx, val := range reportValues {
		curVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Failed to Atoi curVal")
		}

		if idx == 0 {
			checkedReportValues = append(checkedReportValues, curVal)
			continue
		}

		prevVal, err := strconv.Atoi(reportValues[idx-1])
		if err != nil {
			log.Fatalf("Failed to Atoi prevVal")
		}

		if checkDirection {
			// set which direction we assume we're going in
			if curVal > prevVal {
				direction = 1
				checkDirection = false
			} else if curVal < prevVal {
				direction = -1
				checkDirection = false
			} else {
				log.Println("Value is neither higher or lower, postponing direction check")
			}
		} else {
			if curVal < prevVal && direction > 0 {
				log.Println("value is lower than previous, but direction says it should be higher")
				reportProblems += 1
			} else if curVal > prevVal && direction < 0 {
				log.Println("value is higher than previous, but direction says it should be lower")
				reportProblems += 1
			}
		}

		if intAbs(prevVal - curVal) > 3 {
			log.Println("Difference is bigger than 3, report is invalid")
			reportProblems += 1
		}

		if prevVal == curVal {
			log.Println("Must differ by at least one, report is invalid")
			reportProblems += 1
		}

		if slices.Contains(checkedReportValues[:len(checkedReportValues)-1], curVal) {
			log.Printf("Current value already in checked values: %d\n", curVal)
			reportProblems += 1
		}

		checkedReportValues = append(checkedReportValues, curVal)

	}

	if reportProblems == len(reportValues) - 2 {
		log.Println("DEBUG: possible direction problem in autodetection between first and second value!")
	}

	if reportProblems == 0 {
		return true
	} else if reportProblems == 1 {
		log.Println("Problem Dampener initiated, only one problem spotted.")
		return true
	}
	return false

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumValidReports := 0
	for scanner.Scan() {
		line := scanner.Text()

		if checkReportValid(strings.Fields(line)) {
			log.Printf("Valid report: %s\n", line)
			sumValidReports += 1
		} else {
			log.Printf("Invalid report: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error!")
	}

	fmt.Printf("Summary of valid reports: %d\n", sumValidReports)
}
