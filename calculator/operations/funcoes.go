package operations

import "errors"

func Add(number1 int, number2 int) int {
	return number1 + number2
}

func Subtract(minuend int, subtrahend int) int {
	return minuend - subtrahend
}

func Multiply(multiplicand int, muiltiplier int) int {
	return multiplicand * muiltiplier
}

func Divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("não é possível realizar divisão por zero")
	}

	quotient := dividend / divisor

	return quotient, nil
}
