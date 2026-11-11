package lexer

func (l *Lexer) readNumber() {
	start := l.Pos

	for l.Pos < l.Length && isDigit(l.current()) {
		l.advance()
	}

	l.addToken(TOKEN_INTEGER, l.Source[start:l.Pos])
}
