package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/adam-lavrik/go-imath/ix"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberOfValidReports := 0
	for scanner.Scan() {
		line := scanner.Text()
		report, err := parseLine(line)
		if err != nil {
			log.Panicf("Error parsing line: %v", err)
		}
		if validateReport(report) {
			numberOfValidReports++
		}
	}
	fmt.Printf("There were %d valid reports.\n", numberOfValidReports)

}

func parseLine(line string) ([]int, error) {
	fields := strings.Fields(line)
	var levels []int
	for _, field := range fields {
		level, err := strconv.Atoi(field)
		if err != nil {
			log.Panicf("could not convert to int: %v", level)
		}
		levels = append(levels, level)
	}
	return levels, nil
}

func validateReport(report []int) bool {
	prevDifference := 0
	for i := 0; i < (len(report) - 1); i++ {
		difference := report[i] - report[i+1]
		if prevDifference != 0 {
			if difference > 0 && prevDifference < 0 {
				return false
			}
			if difference < 0 && prevDifference > 0 {
				return false
			}
		}

		absDifference := ix.Abs(difference)
		if absDifference > 3 || absDifference == 0 {
			return false
		}
		prevDifference = difference
	}
	return true
}
