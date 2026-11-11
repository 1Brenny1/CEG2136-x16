package lexer

func (l *Lexer) peek(offset int) byte {
	if l.Pos+offset < l.Length {
		return l.Source[l.Pos+offset]
	}

	return 0
}

func (l *Lexer) current() byte {
	return l.peek(0)
}

func (l *Lexer) advance() {
	if l.current() == '\n' {
		l.Line++
		l.Column = 1
	} else {
		l.Column++
	}

	l.Pos++
}

func (l *Lexer) addToken(tokenType TokenType, value string) {
	l.Tokens = append(l.Tokens, Token{
		Type:   tokenType,
		Value:  value,
		Line:   l.Line,
		Column: l.Column,
	})
}

func isAlpha(c byte) bool {
	return (c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
