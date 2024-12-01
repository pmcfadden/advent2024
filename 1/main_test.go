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

func TestDiffNumbersShouldReturnDifference(t *testing.T) {
  difference := diffNumbers(3, 1)
  if difference != 2 {
    t.Errorf("Expected difference of 2, but got %d", difference)
  }
}

func TestDiffNumbersShouldReturnAbsValueDifference(t *testing.T) {
  difference := diffNumbers(1, 3)
  if difference != 2 {
    t.Errorf("Expected difference of 2, but got %d", difference)
  }
}

func TestSumlistsReturnsDifferenceBetweenLists(t *testing.T) {
  list1 := []int{1}
  list2 := []int{3}
  total := sumDifferenceBetweenLists(list1, list2)
  if total != 2 {
    t.Errorf("Expected total of 2 but got %d", total)
  }
}

func TestSumlistsReturnsDifferenceBetweenMultipleItemsInLists(t *testing.T) {
  list1 := []int{1, 5}
  list2 := []int{3, 6}
  total := sumDifferenceBetweenLists(list1, list2)
  if total != 3 {
    t.Errorf("Expected total of 3 but got %d", total)
  }
}
