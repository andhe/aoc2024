package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	reportValid := true // assume until otherwise proven
	direction := 0
	for idx, val := range reportValues {
		if idx == 0 {
			continue
		}

		prevVal, err := strconv.Atoi(reportValues[idx-1])
		if err != nil {
			log.Fatalf("Failed to Atoi prevVal")
		}
		curVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Failed to Atoi curVal")
		}

		if idx == 1 {
			// set which direction we assume we're going in
			if curVal > prevVal {
				direction = 1
			} else if curVal < prevVal {
				direction = -1
			} else {
				log.Println("Value is neither higher or lower, bad report")
				reportValid = false
				break
			}
		} else {
			if curVal < prevVal && direction > 0 {
				log.Println("value is lower than previous, but direction says it should be higher")
				reportValid = false
				break
			} else if curVal > prevVal && direction < 0 {
				log.Println("value is higher than previous, but direction says it should be lower")
				reportValid = false
				break
			}
		}

		if intAbs(prevVal - curVal) > 3 {
			log.Println("Difference is bigger than 3, report is invalid")
			reportValid = false
			break
		}

		if prevVal == curVal {
			log.Println("Must differ by at least one, report is invalid")
			reportValid = false
			break
		}

	}

	return reportValid
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
