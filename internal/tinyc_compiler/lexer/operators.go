package lexer

import "fmt"

func (l *Lexer) readOperator() error {
	c := l.current()
	l.advance()
	nc := l.current()

	if op, ok := Operators[fmt.Sprintf("%c%c", c, nc)]; ok {
		l.advance() // Consume 2nd Character
		l.addToken(op, fmt.Sprintf("%c%c", c, nc))
		return nil
	}

	if op, ok := Operators[string(c)]; ok {
		l.addToken(op, string(c))
		return nil
	}

	return fmt.Errorf("No token mapping for operator: %c", c)
}

func (l *Lexer) readPunctuation() error {
	c := l.current()
	tokenType, ok := Punctuation[c]
	if !ok {
		return fmt.Errorf("No token mapping for punctuation: %c", c)
	}

	l.advance()
	l.addToken(tokenType, string(c))

	return nil
}
