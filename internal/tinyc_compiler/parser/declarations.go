package parser

import (
	"fmt"
	"x16/internal/tinyc_compiler/ast"
	"x16/internal/tinyc_compiler/lexer"
)

func (p *Parser) parseDeclaration() (ast.Declaration, error) {
	switch p.current().Type {
	case lexer.TOKEN_FUNC:
		return p.parse_function()
	case lexer.TOKEN_VAR:
		return p.parse_variable()
	default:
		return nil, fmt.Errorf("Expected Declaration")
	}
}

func (p *Parser) parse_function() (ast.Declaration, error) {
	p.advance()

	name, err := p.expect(lexer.TOKEN_IDENTIFIER)
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}

	params := []string{}

	for !p.check(lexer.TOKEN_RPAREN) {
		param, err := p.expect(lexer.TOKEN_IDENTIFIER)
		if err != nil {
			return nil, err
		}

		params = append(params, param.Value)

		if !p.match(lexer.TOKEN_COMMA) {
			break
		}
	}

	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &ast.FunctionDeclaration{
		Name:       name.Value,
		Parameters: params,
		Body:       body,
	}, nil
}

func (p *Parser) parse_variable() (ast.Declaration, error) {
	p.advance()

	name, err := p.expect(lexer.TOKEN_IDENTIFIER)
	if err != nil {
		return nil, err
	}

	if !p.match(lexer.TOKEN_ASSIGN) {
		if _, err := p.expect(lexer.TOKEN_SEMICOLON); err != nil {
			return nil, err
		}

		return &ast.VariableDeclaration{
			Name:  name.Value,
			Value: nil,
		}, nil
	}

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(lexer.TOKEN_SEMICOLON); err != nil {
		return nil, err
	}

	return &ast.VariableDeclaration{
		Name:  name.Value,
		Value: value,
	}, nil
}
