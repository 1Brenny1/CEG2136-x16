package lexer

import (
	"unicode"
)

type Lexer struct {
	Source string
	Pos    int
	Length int

	Line   int
	Column int

	Tokens []Token
}

func New(source string) *Lexer {
	return &Lexer{
		Source: source,

		Pos:    0,
		Length: len(source),

		Line:   1,
		Column: 1,
	}
}

func (l *Lexer) Tokenize() ([]Token, error) {
	for l.Pos < l.Length {
		c := l.current()

		if unicode.IsSpace(rune(c)) {
			l.advance()
			continue
		}

		if isAlpha(c) || c == '_' {
			l.readIdentifier()
			continue
		}

		if isDigit(c) {
			l.readNumber()
			continue
		}

		switch c {
		case '/':
			err := l.skipComment()
			if err != nil {
				return nil, err
			}

		case '\'':
			err := l.readCharacter()
			if err != nil {
				return nil, err
			}

		case '"':
			err := l.readString()
			if err != nil {
				return nil, err
			}

		case '+', '-', '=', '!', '<', '>', '&', '*', '|', '^', '~':
			err := l.readOperator()
			if err != nil {
				return nil, err
			}

		case ';', ',', '(', ')', '{', '}', '[', ']':
			err := l.readPunctuation()
			if err != nil {
				return nil, err
			}

		}

	}

	return l.Tokens, nil
}
