package calcfunctions

import (
	"errors"
	"strconv"
)

var ErrZeroDivision = errors.New("division by zero")

// Функция, которая вычисляет значение выражения, используя обратную польскую запись
func EvalRPN(expr []string) (float64, error) {
	stack := make([]float64, 0)

	operation := map[string]func(float64, float64) (float64, error){
		"+": func(a, b float64) (float64, error) { return a + b, nil },
		"-": func(a, b float64) (float64, error) { return a - b, nil },
		"*": func(a, b float64) (float64, error) { return a * b, nil },
		"/": func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, ErrZeroDivision
			}
			return a / b, nil
		},
	}

	for _, token := range expr {
		if op, ok := operation[token]; ok {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res, err := op(a, b)
			if err != nil {
				return 0, err
			}
			stack = append(stack, res)
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		}
	}

	return stack[0], nil
}
