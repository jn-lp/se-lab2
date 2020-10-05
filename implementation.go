package lab2

import "regexp"

type astNode struct {
	left *astNode
	right *astNode
	value string
}

func isOperator(char string) bool {
	res, _ := regexp.MatchString("([+\\-*/^])", char)
	return res
}
func isNumber(char string) bool {
	res, _ := regexp.MatchString("([0-9])", char)
	return res
}

func astFromPostfix(expr string) *astNode {
	var stack [] *astNode
	var operand = ""
	for i:=0; i<len(expr); i++	{
		var char = string(expr[i])
		if isNumber(char) {
			operand += char
		} else {
			if len(operand) > 0 {
				stack = append(stack, &astNode{nil, nil, operand})
				operand = ""
			}
			if isOperator(char) && len(stack) >= 2 {
				stack[len(stack)-2] = &astNode {stack[len(stack)-2], stack[len(stack)-1], char}
				stack = stack[:len(stack) - 1]
			}
		}
	}
	if len(stack) >= 1 {
		return stack[len(stack) - 1]
	}
	return &astNode{nil, nil, ""}
}

func operatorPriority(op string) int {
	if op == "+" || op == "-" { return 0 }
	if op == "*" || op == "/" { return 1 }
	if op == "^" { return 2 }
	return -1
}

func (node *astNode) toInfix() string {
	var left string
	var right string
	if node.left != nil {
		left = node.left.toInfix()
		if isOperator(node.value) && isOperator(node.left.value) &&
			operatorPriority(node.value) > operatorPriority(node.left.value) {
			left = "(" + left + ")"
		}
	}
	if node.right != nil {
		right = node.right.toInfix()
	}
	return left + node.value + right
}

// TODO: document this function.
func PostfixToInfix(expr string) string {
	return astFromPostfix(expr).toInfix()
}
