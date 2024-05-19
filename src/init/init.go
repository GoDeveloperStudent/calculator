package init

import mathlogic "GoMod/src/math-logic"

func Initfunc(num1 float32, num2 float32, selected string) (InitAnswer string) {
	switch selected {
	case "+":
		InitAnswer = mathlogic.Sum(num1, num2)
	case "-":
		InitAnswer = mathlogic.Sub(num1, num2)
	case "*":
		InitAnswer = mathlogic.Div(num1, num2)
	case "/":
		InitAnswer = mathlogic.Mul(num1, num2)
	}
	return InitAnswer
}
