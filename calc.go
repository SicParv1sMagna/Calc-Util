package main

import (
	"fmt"
	"os"

	calc "calc/calcfunctions"
)

func main() {
	expr := os.Args[1]

	rpnExpr := calc.ConvertToRPN(expr)
	result, _ := calc.EvalRPN(rpnExpr)

	fmt.Println(result)
}
