package lexer

import (
	"go_sol/daythree/tokentype"
	"strconv"
	"strings"
)

type Token struct {
	Type  tokentype.TokenType
	Value string
}

func createToken(tokenType tokentype.TokenType, value string) Token {
	return Token{tokenType, value}
}

func isWhitespace(src string) bool {
	return src == " " || src == "\t" || src == "\n" || src == "\r"
}

func isInt(src string) bool {
	_, err := strconv.Atoi(src)
	return err == nil
}

func Tokenize(sourceCode string) []int {
	tokens := []Token{}
	src := strings.Split(sourceCode, "")

	for len(src) > 0 {
		if isWhitespace(src[0]) {
			src = src[1:]
		} else if src[0] == "." {
			tokens = append(tokens, createToken(tokentype.Dot, src[0]))
			src = src[1:]
		} else {
			num := ""
			isPart := false

			for len(src) > 0 && (src[0] != "." && !isWhitespace(src[0])) {
				num += src[0]
				src = src[1:]
			}

			if strings.HasPrefix(num, "*") || strings.HasSuffix(num, "*") ||
					strings.HasPrefix(num, "#") || strings.HasSuffix(num, "#") ||
					strings.HasPrefix(num, "@") || strings.HasSuffix(num, "@") ||
					strings.HasPrefix(num, "$") || strings.HasSuffix(num, "$") ||
					strings.HasPrefix(num, "%") || strings.HasSuffix(num, "%") ||
					strings.HasPrefix(num, "^") || strings.HasSuffix(num, "^") ||
					strings.HasPrefix(num, "&") || strings.HasSuffix(num, "&") ||
					strings.HasPrefix(num, "+") || strings.HasSuffix(num, "+") ||
					strings.HasPrefix(num, "-") || strings.HasSuffix(num, "-") ||
					strings.HasPrefix(num, "=") || strings.HasSuffix(num, "=") ||
					strings.HasPrefix(num, "/") || strings.HasSuffix(num, "/") {
				isPart = true
			}

			if !isPart {
				tokens = append(tokens, createToken(tokentype.WholeNumber, num))
			} else {
				if strings.HasPrefix(num, "*") || strings.HasPrefix(num, "#") ||
						strings.HasPrefix(num, "@") || strings.HasPrefix(num, "$") ||
						strings.HasPrefix(num, "%") || strings.HasPrefix(num, "^") ||
						strings.HasPrefix(num, "&") || strings.HasPrefix(num, "+") ||
						strings.HasPrefix(num, "-") || strings.HasPrefix(num, "=") ||
						strings.HasPrefix(num, "/") {
					num = num[1:]
				}
				if strings.HasSuffix(num, "*") || strings.HasSuffix(num, "#") ||
						strings.HasSuffix(num, "@") || strings.HasSuffix(num, "$") ||
						strings.HasSuffix(num, "%") || strings.HasSuffix(num, "^") ||
						strings.HasSuffix(num, "&") || strings.HasSuffix(num, "+") ||
						strings.HasSuffix(num, "-") || strings.HasSuffix(num, "=") ||
						strings.HasSuffix(num, "/") {
					num = num[:len(num)-1]
				}
				
				tokens = append(tokens, createToken(tokentype.PartNumber, num))
			}

		}
	}

		partNumbers := []int{}

		for _, token := range tokens {
			if token.Type == tokentype.PartNumber {
				num, _ := strconv.Atoi(token.Value)
				partNumbers = append(partNumbers, num)
			}
		}

		return partNumbers
	}
