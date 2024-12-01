package main

import (
  "testing"
)

func  TestParseLineIntoTwoInts(t *testing.T) {
  fileLine := "1   3"
  num1, num2, _ := parseLine(fileLine)
  if num1 != 1 { 
    t.Errorf(`Num1 was %v and should have been 1`, num1)
  }

  if num2 != 3 {
    t.Errorf(`Num2 was %v and should have been 3`, num1)
  }
}

func TestParseLineShouldOnlyHaveTwoNumbers(t *testing.T) {
  fileLine := "1   3 4"
  _, _, err := parseLine(fileLine)
  if err == nil {
    t.Errorf("Expected an error but did not get one")
  }
}
