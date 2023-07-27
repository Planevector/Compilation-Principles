package main

import (
	"fmt"
	"strings"
)

// 中缀转后缀
func infixToPostfix(infix_expr string) (string,error) {
	// 操作符栈
	// operatorStack := []rune{}
	// rune == int32
	operator_stack := make([]rune, 0, 10)  
	// 后缀表达式构建器
	postfix_expr := strings.Builder{}
	// 运算符优先级映射
	precedence := map[rune]int {
		'+':1,
		'-':1,
		'*':2,
		'/':2,
		'%':2,
	}

	for _,char := range infix_expr {
		switch char {
		case ' ':
			continue // 忽略空格
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':// 如果是数字则直接输出到后缀表达式
			postfix_expr.WriteRune(char)
		case '(':
			operator_stack = append(operator_stack, char)// 如果是左括号 '('，将其入栈
		case ')':
			// 处理右括号 ')'
			for len(operator_stack) > 0 && operator_stack[len(operator_stack)-1] != '(' {
			// 弹出栈顶的操作符直到遇到左括号 '('，并将这些操作符输出到后缀表达式
				postfix_expr.WriteRune(' ')
				postfix_expr.WriteRune(operator_stack[len(operator_stack)-1])
				operator_stack = operator_stack[:len(operator_stack)-1]
			}
			// 如果栈为空或栈顶不是左括号 '('，表示括号不匹配
			if len(operator_stack) == 0 || operator_stack[len(operator_stack)-1] != '(' {
				return "",fmt.Errorf("Unbalanced parentheses")
			}
			operator_stack = operator_stack[:len(operator_stack) - 1] // 弹出左括号 '('

		default:
			for len(operator_stack) > 0 && operator_stack[len(operator_stack) -1] != '(' && precedence[operator_stack[len(operator_stack)-1]] >= precedence[char] {
				// 比较栈顶操作符和当前操作符的优先级，如果栈顶操作符优先级较高，则将其输出到后缀表达式
				postfix_expr.WriteRune(' ')
				postfix_expr.WriteRune(operator_stack[len(operator_stack) -1 ])
				operator_stack = operator_stack[:len(operator_stack) -1 ]
			} 
			postfix_expr.WriteRune(' ')
			operator_stack = append(operator_stack, char)// 当前操作符入栈
		}
	}

	for len(operator_stack) > 0 {
		if operator_stack[len(operator_stack) -1 ] == '(' {
			return "", fmt.Errorf("Unbalanced parentheses") // 如果栈中还有左括号 '('，表示括号不匹配
		}
		postfix_expr.WriteRune(' ')
		postfix_expr.WriteRune(operator_stack[len(operator_stack) -1 ])
		operator_stack = operator_stack[:len(operator_stack) -1 ]
	}

	return postfix_expr.String(),nil
}

func main(){
	infix_expr := "5 + (7 - 2) * 3"
	postfix_expr,err := infixToPostfix(infix_expr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("中缀: %s\n", infix_expr)
	fmt.Printf("后缀: %s\n", postfix_expr)
}