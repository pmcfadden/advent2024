package main

import (
	"reflect"
	"testing"
)

func TestReturnMatchesFindsASingleMatch(t *testing.T) {
	testString := "smul(1,2)"
	got := returnMatches(testString)
	want := []string{"mul(1,2)"}
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("ReturnMatches = %v, want %v", got, want)
	}
}

func TestReturnMatchesFindsMultipleMatches(t *testing.T) {
	testString := "smul(1,2)mul(2,3)"
	got := returnMatches(testString)
	want := []string{"mul(1,2)", "mul(2,3)"}
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("ReturnMatches = %v, want %v", got, want)
	}
}

func TestParseAndMultiplyNumbers(t *testing.T) {
	testString := "mul(13,2)"
	got := parseAndMultiplyNumbers(testString)
	want := 26
	if got != want {
		t.Errorf("parseAndMultiplyNumbers = %v, want %v", got, want)
	}
}
