package daythree

import (
	_ "embed"
	"fmt"
	"go_sol/utils"
	"log"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

func EnginePartNumbers(data *string, partType utils.PartType) (int, int) {
	matrix := map[utils.Coordinates2D]rune{}
	for idx, line := range strings.Fields(string(*data)) {
		for jdx, ch := range line {
			if ch != '.' && !utils.IsDigit(ch) {
				matrix[utils.Coordinates2D{X: jdx, Y: idx}] = ch
			}
		}
	}

	parts := map[utils.Coordinates2D][]int{}
	lines := strings.Fields(string(*data))

	for idx, line := range lines {
		startIdx := -1
		for i, ch := range line {
			if utils.IsDigit(ch) {
				if startIdx == -1 {
					startIdx = i
				}
			} else {
				if startIdx != -1 {
					num, err := strconv.Atoi(line[startIdx:i])
					if err != nil {
						log.Fatal(err)
					}

					bounds := make(map[utils.Coordinates2D]struct{})
					for x := startIdx; x < i; x++ {
						for _, coord := range []utils.Coordinates2D{
							{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: -1},
							{X: 0, Y: 1}, {X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1},
						} {
							bounds[utils.Coordinates2D{X: x, Y: idx}.Add(coord)] = struct{}{}
						}
					}

					for bound := range bounds {
						if _, ok := matrix[bound]; ok {
							parts[bound] = append(parts[bound], num)
						}
					}
					startIdx = -1
				}
			}
		}

		if startIdx != -1 {
			num, err := strconv.Atoi(line[startIdx:])
			if err != nil {
				log.Fatal(err)
			}

			bounds := make(map[utils.Coordinates2D]struct{})
			for x := startIdx; x < len(line); x++ {
				for _, coord := range []utils.Coordinates2D{
					{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: -1},
					{X: 0, Y: 1}, {X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1},
				} {
					bounds[utils.Coordinates2D{X: x, Y: idx}.Add(coord)] = struct{}{}
				}
			}

			for bound := range bounds {
				if _, ok := matrix[bound]; ok {
					parts[bound] = append(parts[bound], num)
				}
			}
		}
	}

	part1, part2 := 0, 0
	for coord, nums := range parts {
		product := 1
		for _, num := range nums {
			if partType == utils.PartOne {
				part1 += num
			} else {
				product *= num
			}
		}
		if partType == utils.PartTwo {
			if matrix[coord] == '*' && len(nums) == 2 {
				part2 += product
			}
		}
	}

	if partType == utils.PartOne {
		return part1, 0
	}

	return 0, part2
}

func DayThreeMain() {
	part1, part2 := EnginePartNumbers(&input, utils.PartTwo)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
