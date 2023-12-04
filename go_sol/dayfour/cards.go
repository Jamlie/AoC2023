package dayfour

import (
	_ "embed"
	"fmt"
	"go_sol/utils"
	"math"
	// "slices"
	"strings"
)

//go:embed data.txt
var input string

func CardsPartOne(data string, cardsScore *[]int) int {
	indexOfColon := strings.Index(data, ":")
	if indexOfColon == -1 {
		return 0
	}
	indexOfPipe := strings.Index(data, "|")
	if indexOfPipe == -1 {
		return 0
	}

	allWinningCards := strings.TrimSpace(data[indexOfColon+2:indexOfPipe-1])
	winningCards := utils.StrArrayToIntArray(strings.Split(allWinningCards, " "))

	allUserCards := strings.TrimSpace(data[indexOfPipe+2:])
	userCards := utils.StrArrayToIntArray(strings.Split(allUserCards, " "))

	userWinningCards := userCards.Filter(func(card int) bool {
		return winningCards.Contains(card)
	})

	*cardsScore = append(*cardsScore, len(userWinningCards))

	count := 0

	if len(userWinningCards) != 0 {
		count += int(math.Pow(2, float64(len(userWinningCards)-1)))
	}

	return count
}

func ScratchCardsPart2(cardsScore *[]int) int {
	scores := *cardsScore
	copies := make(utils.Array[int], len(*cardsScore))

	copies.Fill(1)

	for i := range utils.Range(0, len(copies)) {
		for j := 1; j <= scores[i]; j++ {
			copies[i + j] += copies[i]
		}
	}

	sum := 0

	copies.ForEach(func(v int) {
		sum += v
	})

	return sum
}

func DayFourMain() {
	lines := strings.Split(input, "\n")
	var cardsScore []int

	sum := 0
	for _, line := range lines {
		CardsPartOne(line, &cardsScore)
	}

	sum += ScratchCardsPart2(&cardsScore)

	fmt.Println(sum)
}
