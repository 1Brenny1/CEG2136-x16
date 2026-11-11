package lexer

func (l *Lexer) readIdentifier() {
	start := l.Pos

	for l.Pos < l.Length {
		c := l.current()

		if !isAlpha(c) && !isDigit(c) && c != '_' {
			break
		}

		l.advance()
	}

	value := l.Source[start:l.Pos]

	tokenType := TOKEN_IDENTIFIER

	if keyword, ok := Keywords[value]; ok {
		tokenType = keyword
	}

	l.addToken(tokenType, value)
}
