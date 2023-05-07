package calcfunctions

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func symbolise(expr string) []string {
	expr = strings.ReplaceAll(expr, " ", "")
	symbol := make([]string, 0)

	for i := 0; i < len(expr); {
		r, size := utf8.DecodeRuneInString(expr[i:])
		switch {
		case r == '+', r == '-', r == '*', r == '/', r == '(', r == ')':
			symbol = append(symbol, string(r))
			i += size
		case unicode.IsDigit(r), r == '.':
			j := i
			for j < len(expr) {
				r, size = utf8.DecodeRuneInString(expr[j:])
				if !unicode.IsDigit(r) && r != '.' {
					break
				}
				j += size
			}
			symbol = append(symbol, expr[i:j])
			i = j
		default:
			i += size
		}
	}

	return symbol
}
