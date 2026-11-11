package parser

import (
	"fmt"
	"x16/internal/tinyc_compiler/lexer"
)

func (p *Parser) current() lexer.Token {
	if p.Pos >= p.Length {
		return lexer.Token{
			Type: lexer.TOKEN_EOF,
		}
	}

	return p.Tokens[p.Pos]
}

func (p *Parser) peek() lexer.Token {
	if p.Pos+1 >= p.Length {
		return lexer.Token{
			Type: lexer.TOKEN_EOF,
		}
	}

	return p.Tokens[p.Pos+1]
}

func (p *Parser) advance() lexer.Token {
	token := p.current()

	if p.Pos < p.Length {
		p.Pos++
	}

	return token
}

func (p *Parser) check(tokenType lexer.TokenType) bool {
	return p.current().Type == tokenType
}

func (p *Parser) match(types ...lexer.TokenType) bool {
	for _, t := range types {
		if p.current().Type == t {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) expect(t lexer.TokenType) (lexer.Token, error) {
	token := p.current()

	if token.Type != t {
		return token, fmt.Errorf("Expected %v, got %v at line %d(%d)", t, token.Type, token.Line, token.Column)
	}

	return p.advance(), nil
}
func (p *Parser) isAtEnd() bool {
	return p.current().Type == lexer.TOKEN_EOF
}
