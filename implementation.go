package lab2

import (
	"errors"
	"strings"
)

func isOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/" || op == "^"
}

func PostfixToPrefix(input string) (string, error) {
	tokens := strings.Fields(input)
	stack := []string{}

	for _, token := range tokens {
		if !isOperator(token) {
			stack = append(stack, token)
		} else {
			if len(stack) < 2 {
				return "", errors.New("invalid postfix expression: insufficient operands")
			}
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			expression := token + " " + operand1 + " " + operand2
			stack = append(stack, expression)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid postfix expression: too many operands")
	}

	return stack[0], nil
}
