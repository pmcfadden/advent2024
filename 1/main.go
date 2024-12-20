package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
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

  list1 := []int{}
  list2 := []int{}

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    item1, item2, err := parseLine(scanner.Text())
    if err != nil{
      panic(err)
    }
    list1 = append(list1, item1)
    list2 = append(list2, item2)
  }

  sort.Ints(list1)
  sort.Ints(list2)

  fmt.Println(sumDifferenceBetweenLists(list1, list2))
  fmt.Printf("Total Similarity Rating: %d", findTotalWithRepeats(list1, list2))

}

func parseLine(fileLine string) (int, int, error) {
  fields := strings.Fields(fileLine)
  if len(fields) != 2 {
    return -1, -1, errors.New("Incorrect number of fields")
  }
  num1, err := strconv.Atoi(fields[0])
  if err != nil {
    return -1, -1, err
  }
  num2, err := strconv.Atoi(fields[1])
  if err != nil {
    return -1, -1, err
  }
  return num1, num2, nil
}

func diffNumbers(num1 int, num2 int) int {
  return ix.Abs(num1 - num2)
}

func sumDifferenceBetweenLists(list1 []int, list2 []int) int {
  total := 0
  for i := 0 ; i < len(list1) ; i++ {
    total += diffNumbers(list1[i], list2[i])
  }
  return total
}

func findTotalWithRepeats(list1 []int, list2 []int) int {
  countInList2 := make(map[int]int)
  for _, num := range list2 {
    countInList2[num] = countInList2[num] +1
  }

  total := 0
  for _, num := range list1 {
    total += (num * countInList2[num])
  }
  return total
}
