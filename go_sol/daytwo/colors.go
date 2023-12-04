package daytwo

import (
	"bufio"
	_ "embed"
	"fmt"
	"go_sol/utils"
	"log"
	"strings"
)

//go:embed data.txt
var input string

type Game struct {
	Id		int
	Blue  int
	Red   int
	Green int
}

func NewGame() *Game {
	return &Game{
		Id: -1,
		Blue: -1,
		Red: -1,
		Green: -1,
	}
}

func (g *Game) String() string {
	return fmt.Sprintf("Game %d: %d blue, %d red, %d green", g.Id, g.Blue, g.Red, g.Green)
}

func PossibleGames(s string, gameType utils.PartType) (*Game, bool) {
	var id, red, blue, green int
	game := NewGame()

	_, err := fmt.Sscanf(s, "Game %d:", &id)
	if err != nil {
		return NewGame(), false
	}

	game.Id = id

	indexOfColor := strings.Index(s, ":")
	if indexOfColor == -1 {
		return NewGame(), false
	}

	colors := s[indexOfColor+2:]
	var count int

	for _, color := range strings.Split(colors, ";") {
		splitColor := strings.Split(color, ",")
		for _, c := range splitColor {
			_, err := fmt.Sscanf(strings.TrimSpace(c), "%d %s", &count, &color)
			if err != nil {
				return NewGame(), false
			}

			if gameType == utils.PartOne {
				switch color {
				case "blue":
					if count > 14 {
						failedGame := NewGame()
						failedGame.Id = id
						return failedGame, false
					}
					blue += count
				case "red":
					if count > 12 {
						failedGame := NewGame()
						failedGame.Id = id
						return failedGame, false
					}
					red += count
				case "green":
					if count > 13 {
						failedGame := NewGame()
						failedGame.Id = id
						return failedGame, false
					}
					green += count
				}
			} else {
				// part two
				switch color {
				case "blue":
					blue = max(blue, count)
				case "red":
					red = max(red, count)
				case "green":
					green = max(green, count)
				}
			}
		}
	}

	game.Blue = blue
	game.Red = red
	game.Green = green

	return game, true
}

func DayTwoMain() {
	var possibleGames []*Game
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		game, ok := PossibleGames(scanner.Text(), utils.PartTwo)
		if !ok {
			fmt.Println(game, "is not possible")
		} else {
			possibleGames = append(possibleGames, game)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	// part one
	// for _, id := range possibleGames {
	// 	sum += id.Id
	// }
	//
	// fmt.Println("Sum of possible games:", sum)

	// part two
	for i := range possibleGames {
		sum += possibleGames[i].Blue * possibleGames[i].Red * possibleGames[i].Green
	}

	fmt.Println("Sum of possible games:", sum)

	// part one
	// var possibleGames []int
	// scanner := bufio.NewScanner(strings.NewReader(input))
	//
	// for scanner.Scan() {
	// 	game, ok := PossibleGamesPartOne(scanner.Text())
	// 	if !ok {
	// 		fmt.Println(game, "is not possible")
	// 	} else {
	// 		possibleGames = append(possibleGames, game.Id)
	// 	}
 // 	}
	//
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// sum := 0
	//
}
