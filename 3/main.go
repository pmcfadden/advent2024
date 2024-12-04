package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var filename string = "./input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()
	fmt.Printf("Opened file: %v\n", filename)
	scanner := bufio.NewScanner(file)
	allMatches := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Scanning text line length: %v\n", len(line))
		matches := returnMatches(line)
		fmt.Printf("Number of matches found: %v\n", len(matches))
		allMatches = append(allMatches, matches...)
	}
	total := 0
	for _, match := range allMatches {
		result := parseAndMultiplyNumbers(match)
		fmt.Println(result)
		total += result
	}
	fmt.Printf("Total result: %v\n", total)
}

func returnMatches(line string) []string {
	r, _ := regexp.Compile("mul\\((\\d){1,3},(\\d){1,3}\\)")
	allString := r.FindAllString(line, -1)
	fmt.Printf("Length: %v\n", len(allString))
	return allString
}

func parseAndMultiplyNumbers(match string) int {
	result := strings.Split(match, ",")
	num1, _ := strconv.Atoi(result[0][4:])
	num2, _ := strconv.Atoi(result[1][:len(result[1])-1])
	return num1 * num2
}
