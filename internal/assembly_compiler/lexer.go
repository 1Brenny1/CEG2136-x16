package assembly_compiler

import "fmt"

type TokenType uint8

const (
	TOKEN_KEYWORD TokenType = iota
	TOKEN_DIRECTIVE
	TOKEN_IDENTIFIER
	TOKEN_NUMBER
	TOKEN_COLON
	TOKEN_INDIRECT
)

func isAlpha(c byte) bool {
	return (c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isHexDigit(c byte) bool {
	return isDigit(c) ||
		(c >= 'A' && c <= 'F') ||
		(c >= 'a' && c <= 'f')
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

type Lexer struct {
	I      int
	N      int
	Source string
	Tokens []Token
}

type Token struct {
	Type  TokenType
	Value string
}

func (l *Lexer) peek(offset int) byte {
	if l.I+offset < l.N {
		return l.Source[l.I+offset]
	}

	return 0
}

func (l *Lexer) Tokenize() ([]Token, error) {
	for l.I < l.N {
		c := l.peek(0)

		switch {
		case c == ' ' || c == '\t' || c == '\r' || c == '\n':
			l.I++

		case isAlpha(c) || c == '_':
			if l.isHexOperand() && isHexDigit(c) {
				l.lexHexNumber()
			} else {
				l.lexWord()
			}

		case isDigit(c) || c == '-':
			if isDigit(c) && l.isHexOperand() {
				l.lexHexNumber()
			} else {
				l.lexNumber()
			}

		case c == ':':
			l.Tokens = append(l.Tokens, Token{
				Type:  TOKEN_COLON,
				Value: ":",
			})
			l.I++

		case c == ';' || c == '/':
			l.skipComment()

		default:
			return nil, fmt.Errorf("Illegal character '%c' at position %d", c, l.I)
		}
	}

	return l.Tokens, nil
}

func (l *Lexer) lexWord() {
	start := l.I
	l.I++

	for l.I < l.N {
		c := l.peek(0)
		if !isAlphaNumeric(c) && c != '_' {
			break
		}

		l.I++
	}

	word := l.Source[start:l.I]

	if _, ok := keywords[word]; ok {
		l.Tokens = append(l.Tokens, Token{
			Type:  TOKEN_KEYWORD,
			Value: word,
		})
	} else if _, ok := directives[word]; ok {
		l.Tokens = append(l.Tokens, Token{
			Type:  TOKEN_DIRECTIVE,
			Value: word,
		})
	} else if word == "I" {
		l.Tokens = append(l.Tokens, Token{
			Type:  TOKEN_INDIRECT,
			Value: "I",
		})
	} else {
		l.Tokens = append(l.Tokens, Token{
			Type:  TOKEN_IDENTIFIER,
			Value: word,
		})
	}
}

func (l *Lexer) lexNumber() {
	start := l.I

	if l.peek(0) == '-' {
		l.I++
	}

	for l.I < l.N && isDigit(l.peek(0)) {
		l.I++
	}

	l.Tokens = append(l.Tokens, Token{
		Type:  TOKEN_NUMBER,
		Value: l.Source[start:l.I],
	})
}

func (l *Lexer) isHexOperand() bool {
	if len(l.Tokens) == 0 {
		return false
	}

	last := l.Tokens[len(l.Tokens)-1]
	return last.Type == TOKEN_DIRECTIVE && last.Value == "HEX"
}

func (l *Lexer) lexHexNumber() {
	start := l.I
	for l.I < l.N && isHexDigit(l.peek(0)) {
		l.I++
	}

	l.Tokens = append(l.Tokens, Token{
		Type:  TOKEN_NUMBER,
		Value: l.Source[start:l.I],
	})
}

func (l *Lexer) skipComment() {
	for l.I < l.N && l.peek(0) != '\n' {
		l.I++
	}
}
