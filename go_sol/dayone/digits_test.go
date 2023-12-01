package dayone_test

import (
	"go_sol/dayone"
	"testing"
)

func TestGetFirstAndLastDigitsInAStringPartOne(t *testing.T) {
	tests := [4]struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		actual, err := dayone.FirstAndLastDigitsPartOne(test.input)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestGetFirstAndLastDigitsInAStringPartTwo(t *testing.T) {
	tests := [7]struct {
		input    string
		expected int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, test := range tests {
		actual, err := dayone.FirstAndLastDigitsPartTwo(test.input)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
