package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
  fmt.Println(len(list1))
  fmt.Println(len(list2))

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
