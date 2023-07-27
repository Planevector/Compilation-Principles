package main

import (
	"fmt"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MULTIPLY
	LEFT
	RIGHT
	INVALID
)

type Token struct {
	Type  TokenType
	Value string
}

func lex(input string) []Token {
	var tokens []Token
	var currentState TokenType
	var currentToken string

	for i := 0; i < len(input); i++ {
		char := input[i]
		fmt.Println(string(char))

		switch currentState {
		case NUMBER:
			if isDigit(char) {
				currentToken += string(char)
			} else {
				tokens = append(tokens, Token{Type: NUMBER, Value: currentToken})
				currentToken = ""
				currentState = getInitialState(char)
				i--
			}
		case PLUS, MULTIPLY,LEFT,RIGHT:
			tokens = append(tokens, Token{Type: currentState, Value: string(char)})
			currentState = getInitialState(char)

		case INVALID:
			return nil
		default:
			currentState = getInitialState(char)
		}
	}

	// Handle the last token
	if currentState == NUMBER {
		tokens = append(tokens, Token{Type: NUMBER, Value: currentToken})
	} else if currentState == PLUS || currentState == MULTIPLY {
		tokens = append(tokens, Token{Type: currentState, Value: currentToken})
	}

	return tokens
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func getInitialState(char byte) TokenType {
	switch char {
	case '+':
		return PLUS
	case '*':
		return MULTIPLY
	case '(':
		return LEFT
	case ')':
		return RIGHT
	default:
		if isDigit(char) {
			return NUMBER
		}
		return INVALID
	}
}

func main() {
	input := "(2+3)*4"

	tokens := lex(input)

	if tokens == nil {
		fmt.Println("Invalid input")
		return
	}

	for _, token := range tokens {
		fmt.Printf("Token Type: %v, Value: %s\n", token.Type, token.Value)
	}
}