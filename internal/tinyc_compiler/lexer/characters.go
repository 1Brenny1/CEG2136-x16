package lexer

import "fmt"

func (l *Lexer) readString() error {
	l.advance()
	start := l.Pos

	for l.Pos < l.Length {
		switch l.current() {
		case '"':
			str := l.Source[start:l.Pos]
			l.advance()

			l.addToken(TOKEN_STRING, str)
			return nil
		case '\\':
			l.advance()
		case '\n':
			return fmt.Errorf("Unterminated String Literal")
		}

		l.advance()
	}

	return nil
}

func (l *Lexer) readCharacter() error {
	l.advance()

	if l.Pos >= l.Length {
		return fmt.Errorf("Unterminated Character Literal")
	}

	char := l.current()
	if char == '\\' {
		l.advance()

		if l.Pos >= l.Length {
			return fmt.Errorf("Unterminated Character Literal")
		}

		char = l.current()
	}

	l.advance()

	if l.Pos >= l.Length || l.peek(0) != '\'' {
		return fmt.Errorf("Character literal must contain exactly one character")
	}

	l.advance()

	l.addToken(TOKEN_CHARACTER, string(char))

	return nil
}

func (l *Lexer) skipComment() error {
	l.advance()

	switch l.current() {
	case '/':
		for l.Pos < l.Length && l.current() != '\n' {
			l.advance()
		}

		return nil
	case '*':
		l.advance()
		for l.Pos < l.Length && l.current() != '*' && l.peek(1) != '/' {
			l.advance()
		}

		l.advance()
		l.advance()

		return nil
	}

	return fmt.Errorf("Unexpected character after / : %c", l.current())
}
