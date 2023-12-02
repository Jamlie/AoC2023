package daytwo_test

import (
	_ "embed"
	"go_sol/daytwo"
	"testing"
)

var input string

func TestPossibleGamesPartOne(t *testing.T) {
 	tests := []struct {
		input    string
		expected bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	sum := 0
	for _, test := range tests {
		game, ok := daytwo.PossibleGames(test.input, daytwo.PartOne)
		if ok != test.expected {
			t.Errorf("Expected %t, got %t", test.expected, ok)
		}
		if ok {
			sum += game.Id
		}
	}

	if sum != 8 {
		t.Errorf("Expected sum to be 8, got %d", sum)
	}
}

func TestPossibleGamesPartTwo(t *testing.T) {
 	tests := []struct {
		input    string
		expected bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", true},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", true},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	sum := 0
	for _, test := range tests {
		game, ok := daytwo.PossibleGames(test.input, daytwo.PartTwo)
		if ok {
			sum += game.Blue * game.Red * game.Green
		}
	}

	if sum != 2286 {
		t.Errorf("Expected sum to be 2286, got %d", sum)
	}
}
