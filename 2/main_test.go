package main

import (
	"reflect"
	"testing"
)

func TestParseLineShouldReturnArrayOfInts(t *testing.T) {
	line := "1 2 3 4 5 6"
	got, err := parseLine(line)
	if err != nil {
		t.Errorf(err.Error())
	}
	wanted := []int{1, 2, 3, 4, 5, 6}
	if reflect.DeepEqual(wanted, got) == false {
		t.Errorf("arrays were not equal wanted: %v, got: %v", wanted, got)
	}
}

func TestValidateReportShouldAllowIncrease(t *testing.T) {
	report := []int{1, 2}
	got := validateReportWithoutSafety(report)
	if !got {
		t.Errorf("Report should have been valid and was not")
	}
}

func TestValidateReportShouldNotAllowLargeIncrease(t *testing.T) {
	report := []int{1, 5}
	got := validateReportWithoutSafety(report)
	if got {
		t.Errorf("Report should not have been valid and was")
	}
}

func TestValidateReportShouldNotAllowLargeDecrease(t *testing.T) {
	report := []int{5, 1}
	got := validateReportWithoutSafety(report)
	if got {
		t.Errorf("Report should not have been valid and was")
	}
}

func TestValidateReportShouldNotAllowNoChange(t *testing.T) {
	report := []int{5, 5}
	got := validateReportWithoutSafety(report)
	if got {
		t.Errorf("Report should not have been valid and was")
	}
}

func TestValidateReportShouldNotAllowIncreaseThenDecrease(t *testing.T) {
	report := []int{1, 3, 2}
	got := validateReportWithoutSafety(report)
	if got {
		t.Errorf("Report should not have been valid and was")
	}
}

func TestValidateReportShouldNotAllowDecreaseThenIncrease(t *testing.T) {
	report := []int{3, 1, 2}
	got := validateReportWithoutSafety(report)
	if got {
		t.Errorf("Report should not have been valid and was")
	}
}

func TestValidateReportShouldAllowASingleBadLevel(t *testing.T) {
	report := []int{1, 3, 2}
	got := validateReportWithSafety(report)
	if !got {
		t.Errorf("Report should have been valid and was not")
	}
}

func TestValidateReportShouldNotAllowMultipleBadLevels(t *testing.T) {
	report := []int{1, 3, 2, 2}
	got := validateReportWithSafety(report)
	if got {
		t.Errorf("Report should have been valid and was not")
	}
}
