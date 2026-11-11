package parser

import (
	"x16/internal/tinyc_compiler/ast"
	"x16/internal/tinyc_compiler/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	Pos    int
	Length int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{
		Tokens: tokens,
		Pos:    0,
		Length: len(tokens),
	}
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := &ast.Program{}

	for !p.isAtEnd() {
		decl, err := p.parseDeclaration()
		if err != nil {
			return nil, err
		}

		program.Declarations = append(program.Declarations, decl)
	}

	return program, nil
}
