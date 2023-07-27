package main

import (
	"fmt"
	"strconv"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MULTIPLY
	LPAREN
	RPAREN
)

type Token struct {
	Type  TokenType
	Value string
}

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) parse() (int, error) {
	return p.parseExpr()
}

func (p *Parser) parseExpr() (int, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for p.match(PLUS) {
		p.advance()
		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}

		left += right
	}

	return left, nil
}

func (p *Parser) parseTerm() (int, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for p.match(MULTIPLY) {
		p.advance()
		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		left *= right
	}

	return left, nil
}

func (p *Parser) parseFactor() (int, error) {
	if p.match(NUMBER) {
		num, err := strconv.Atoi(p.peek().Value)
		if err != nil {
			return 0, err
		}
		p.advance()
		return num, nil
	}

	if p.match(LPAREN) {
		p.advance()
		expr, err := p.parseExpr()
		if err != nil {
			return 0, err
		}

		if !p.match(RPAREN) {
			return 0, fmt.Errorf("Expected closing parenthesis")
		}
		p.advance()

		return expr, nil
	}

	return 0, fmt.Errorf("Unexpected token: %v", p.peek())
}

func (p *Parser) match(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens)
}

func (p *Parser) peek() Token {
	if p.isAtEnd() {
		return Token{}
	}
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	if p.current <= 0 {
		return Token{}
	}
	return p.tokens[p.current-1]
}

func lex(input string) []Token {
	var tokens []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: "+"})
		case '*':
			tokens = append(tokens, Token{Type: MULTIPLY, Value: "*"})
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
		default:
			if isDigit(input[i]) {
				numEnd := i + 1
				for numEnd < len(input) && isDigit(input[numEnd]) {
					numEnd++
				}
				tokens = append(tokens, Token{Type: NUMBER, Value: input[i:numEnd]})
				i = numEnd - 1
			} else if !isWhitespace(input[i]) {
				fmt.Printf("Unknown token: %c\n", input[i])
				return nil
			}
		}
	}

	fmt.Println(tokens)

	return tokens
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func main() {
	input := "2 + 3 * (4 + 5)"

	tokens := lex(input)
	parser := NewParser(tokens)

	result, err := parser.parse()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Expression: %s\n", input)
	fmt.Printf("Result: %d\n", result)
}
