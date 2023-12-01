package dayone

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FirstAndLastDigitsPartOne(s string) (int, error) {
	var firstDigit, lastDigit byte

	for _, char := range s {
		if char >= '0' && char <= '9' {
			if firstDigit == 0 {
				firstDigit = byte(char)
			}
			lastDigit = byte(char)
		}
	}

	if firstDigit != 0 && lastDigit == 0 {
		lastDigit = firstDigit
	}

	if lastDigit != 0 && firstDigit == 0 {
		firstDigit = lastDigit
	}

	if firstDigit == 0 || lastDigit == 0 {
		return 0, fmt.Errorf("No digits found in string")
	}

	return strconv.Atoi(string([]byte{firstDigit, lastDigit}))
}

func FirstAndLastDigitsPartTwo(s string) (int, error) {
	numbers := map[string]byte{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
		"0":     '0',
		"1":     '1',
		"2":     '2',
		"3":     '3',
		"4":     '4',
		"5":     '5',
		"6":     '6',
		"7":     '7',
		"8":     '8',
		"9":     '9',
	}

	result := 0

	var firstChar, lastChar string

	for i := range s {
		char := ""
		if value, ok := numbers[string(s[i])]; ok {
			char = string(value)
		}

		if char == "" {
			for word := range numbers {
				if strings.HasPrefix(s[i:], word) {
					char = string(numbers[word])
				}
			}
		}

		if firstChar == "" && char != "" {
			firstChar = char
		}

		if char != "" {
			lastChar = char
		}
	}

	result += parseAndSum(firstChar, lastChar)

	return result, nil
}

func parseAndSum(firstChar, lastChar string) int {
	firstDigit, _ := strconv.Atoi(firstChar)
	lastDigit, _ := strconv.Atoi(lastChar)
	return firstDigit*10 + lastDigit
}

func DayOneMain() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var firstAndLastDigits []int
	var line string

	for {
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}
		digit, err := FirstAndLastDigitsPartTwo(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}
		firstAndLastDigits = append(firstAndLastDigits, digit)
	}

	var sum int
	for _, digits := range firstAndLastDigits {
		sum += digits
	}

	fmt.Println(sum)
}
