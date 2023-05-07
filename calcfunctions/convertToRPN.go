package calcfunctions

// Определяем приоритет операций
var operatorPrecedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

// Функция, которая преобразует выражение в обратную польскую запись
func ConvertToRPN(expr string) []string {
	operators := make([]string, 0)
	output := make([]string, 0)

	tokens := symbolise(expr)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(operators) > 0 && operatorPrecedence[operators[len(operators)-1]] >= operatorPrecedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		case "(":
			operators = append(operators, token)
		case ")":
			for operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = operators[:len(operators)-1]
		default:
			output = append(output, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output
}
