package parser

import (
	"fmt"
	"strconv"
	"x16/internal/tinyc_compiler/ast"
	"x16/internal/tinyc_compiler/lexer"
)

func (p *Parser) parseBinary(next func() (ast.Expression, error), tokenTypes ...lexer.TokenType) (ast.Expression, error) {
	node, err := next()
	if err != nil {
		return nil, err
	}

	for _, tokType := range tokenTypes {
		if p.current().Type == tokType {
			p.advance()
			right, err := next()
			if err != nil {
				return nil, err
			}

			node = &ast.BinaryExpression{
				Left:     node,
				Operator: tokType,
				Right:    right,
			}
		}
	}

	return node, nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	return p.parse_or()
}

func (p *Parser) parse_or() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_and,
		lexer.TOKEN_LOGICAL_OR,
	)
}

func (p *Parser) parse_and() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_bitwise_or,
		lexer.TOKEN_LOGICAL_AND,
	)
}

func (p *Parser) parse_bitwise_or() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_bitwise_xor,
		lexer.TOKEN_BITWISE_OR,
	)
}

func (p *Parser) parse_bitwise_xor() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_bitwise_and,
		lexer.TOKEN_BITWISE_XOR,
	)
}

func (p *Parser) parse_bitwise_and() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_equality,
		lexer.TOKEN_AMPERSAND,
	)
}

func (p *Parser) parse_equality() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_comparison,
		lexer.TOKEN_EQUAL,
		lexer.TOKEN_NOTEQUAL,
	)
}

func (p *Parser) parse_comparison() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_shift,
		lexer.TOKEN_GREATER,
		lexer.TOKEN_GREATER_EQUAL,
		lexer.TOKEN_LESS,
		lexer.TOKEN_LESS_EQUAL,
	)
}

func (p *Parser) parse_shift() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_term,
		lexer.TOKEN_SHIFT_LEFT,
		lexer.TOKEN_SHIFT_RIGHT,
	)
}

func (p *Parser) parse_term() (ast.Expression, error) {
	return p.parseBinary(
		p.parse_unary, // Skip parse_factor
		lexer.TOKEN_PLUS,
		lexer.TOKEN_MINUS,
	)
}

func (p *Parser) parse_unary() (ast.Expression, error) {
	switch p.current().Type {
	case lexer.TOKEN_LOGICAL_NOT, lexer.TOKEN_BITWISE_NOT, lexer.TOKEN_MINUS, lexer.TOKEN_AMPERSAND, lexer.TOKEN_STAR:
		op := p.advance()

		operand, err := p.parse_unary()
		if err != nil {
			return nil, err
		}

		return &ast.UnaryExpression{
			Operator: op.Type,
			Operand:  operand,
		}, nil
	}

	return p.parse_postfix()
}

func (p *Parser) parse_postfix() (ast.Expression, error) {
	expr, err := p.parse_primary()
	if err != nil {
		return nil, err
	}

	for {
		// Function Call
		if p.match(lexer.TOKEN_LPAREN) {
			args := []ast.Expression{}

			for !p.check(lexer.TOKEN_RPAREN) {
				arg, err := p.parseExpression()
				if err != nil {
					return nil, err
				}

				args = append(args, arg)

				if !p.match(lexer.TOKEN_COMMA) {
					break
				}
			}

			if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
				return nil, err
			}

			expr = &ast.CallExpression{
				Function:  expr,
				Arguments: args,
			}

			continue
		}

		// Array/Pointer indexing
		if p.match(lexer.TOKEN_LBRACKET) {
			index, err := p.parseExpression()
			if err != nil {
				return nil, err
			}

			if _, err := p.expect(lexer.TOKEN_RBRACKET); err != nil {
				return nil, err
			}

			expr = &ast.IndexExpression{
				Object: expr,
				Index:  index,
			}

			continue
		}

		break
	}

	return expr, nil
}

func (p *Parser) parse_primary() (ast.Expression, error) {
	tok := p.current()

	switch tok.Type {
	case lexer.TOKEN_INTEGER:
		p.advance()

		value, err := strconv.ParseUint(tok.Value, 0, 16)
		if err != nil {
			return nil, err
		}

		return &ast.IntegerLiteral{
			Value: uint16(value),
		}, nil

	case lexer.TOKEN_CHARACTER:
		p.advance()

		return &ast.CharacterLiteral{
			Value: tok.Value[0],
		}, nil

	case lexer.TOKEN_STRING:
		p.advance()
		return &ast.StringLiteral{
			Value: tok.Value,
		}, nil

	case lexer.TOKEN_IDENTIFIER:
		p.advance()

		return &ast.Identifier{
			Name: tok.Value,
		}, nil

	case lexer.TOKEN_LPAREN:
		p.advance()

		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
			return nil, err
		}

		return expr, nil
	}

	return nil, fmt.Errorf("Unexpected token %v", tok)
}
