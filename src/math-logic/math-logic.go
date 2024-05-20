package math

import (
	"fmt"
)

func Sum(num1 float32, num2 float32) (answer string) {
	answer = fmt.Sprintf("%f", num1+num2)
	return answer
}
func Sub(num1 float32, num2 float32) (answer string) {
	answer = fmt.Sprintf("%f", num1-num2)
	return answer
}
func Div(num1 float32, num2 float32) (answer string) {
	answer = fmt.Sprintf("%f", num1*num2)
	return answer
}
func Mul(num1 float32, num2 float32) (answer string) {
	answer = fmt.Sprintf("%f", num1/num2)
	return answer
}
