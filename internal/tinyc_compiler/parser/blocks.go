package parser

import (
	"x16/internal/tinyc_compiler/ast"
	"x16/internal/tinyc_compiler/lexer"
)

func (p *Parser) parseBlock() (*ast.BlockStatement, error) {
	if _, err := p.expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}

	block := &ast.BlockStatement{}

	for !p.check(lexer.TOKEN_RBRACE) && !p.isAtEnd() {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		block.Statements = append(block.Statements, stmt)
	}

	if _, err := p.expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return block, nil
}

// func (p *Parser) parseBlockItem() (ast.Statement, error) {
// 	switch p.current().Type {
// 	case lexer.TOKEN_VAR:
// 		break
// 	default:
// 		return p.parseStatement()
// 	}
// }
