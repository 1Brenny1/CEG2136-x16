package parser

import (
	"fmt"
	"x16/internal/tinyc_compiler/ast"
	"x16/internal/tinyc_compiler/lexer"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.current().Type {
	case lexer.TOKEN_VAR:
		decl, err := p.parse_variable()
		if err != nil {
			return nil, err
		}

		return &ast.DeclarationStatement{
			Declaration: decl,
		}, nil

	case lexer.TOKEN_RETURN:
		return p.parse_return()
	case lexer.TOKEN_IF:
		return p.parse_if()
	case lexer.TOKEN_WHILE:
		return p.parse_while()
	case lexer.TOKEN_DO: // Do While?
		return nil, fmt.Errorf("Not Implemented!")
	case lexer.TOKEN_FOR:
		return nil, fmt.Errorf("Not Implemented!")
	case lexer.TOKEN_IDENTIFIER:
		return p.parse_assignment()
	case lexer.TOKEN_LBRACE:
		return p.parseBlock()
	default:
		return nil, fmt.Errorf("Unexpected Token: %v", p.current())
	}
}

func (p *Parser) parse_if() (ast.Statement, error) {
	p.advance() // if

	if _, err := p.expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	ifBlock, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	if !p.match(lexer.TOKEN_ELSE) {
		return &ast.IfStatement{
			Condition: expr,
			Then:      ifBlock,
			Else:      nil,
		}, nil
	}

	var elseBlock *ast.BlockStatement

	if p.check(lexer.TOKEN_IF) {
		stmnt, err := p.parse_if()
		if err != nil {
			return nil, err
		}

		elseBlock = &ast.BlockStatement{
			Statements: []ast.Statement{
				stmnt,
			},
		}
	} else {
		elseBlock, err = p.parseBlock()
		if err != nil {
			return nil, err
		}
	}

	return &ast.IfStatement{
		Condition: expr,
		Then:      ifBlock,
		Else:      elseBlock,
	}, nil
}

func (p *Parser) parse_while() (ast.Statement, error) {
	p.advance() // while

	if _, err := p.expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	whileBlock, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &ast.WhileStatement{
		Condition: expr,
		Body:      whileBlock,
	}, nil
}

func (p *Parser) parse_return() (ast.Statement, error) {
	p.advance() // return

	if p.match(lexer.TOKEN_SEMICOLON) {
		return &ast.ReturnStatement{
			Value: nil,
		}, nil
	}

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(lexer.TOKEN_SEMICOLON); err != nil {
		return nil, err
	}

	return &ast.ReturnStatement{
		Value: expr,
	}, nil
}

func (p *Parser) parse_assignment() (ast.Statement, error) {
	varExpr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// TODO: Add support for +=, -=
	switch p.current().Type {
	case lexer.TOKEN_INCREMENT:
		p.advance()

		if _, err := p.expect(lexer.TOKEN_SEMICOLON); err != nil {
			return nil, err
		}

		return &ast.AssignmentStatement{
			Variable: varExpr,
			Value: &ast.BinaryExpression{
				Left:     varExpr,
				Operator: lexer.TOKEN_PLUS,
				Right: &ast.IntegerLiteral{
					Value: uint16(1),
				},
			},
		}, nil

	case lexer.TOKEN_DECREMENT:
		p.advance()

		if _, err := p.expect(lexer.TOKEN_SEMICOLON); err != nil {
			return nil, err
		}

		return &ast.AssignmentStatement{
			Variable: varExpr,
			Value: &ast.BinaryExpression{
				Left:     varExpr,
				Operator: lexer.TOKEN_MINUS,
				Right: &ast.IntegerLiteral{
					Value: uint16(1),
				},
			},
		}, nil

	case lexer.TOKEN_ASSIGN:
		p.advance()

		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(lexer.TOKEN_SEMICOLON); err != nil {
			return nil, err
		}

		return &ast.AssignmentStatement{
			Variable: varExpr,
			Value:    expr,
		}, nil

	default:
		return nil, fmt.Errorf("Unexpected token: %v", p.current())
	}
}
