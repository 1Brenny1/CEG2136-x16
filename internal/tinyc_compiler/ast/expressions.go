package ast

import "x16/internal/tinyc_compiler/lexer"

type BinaryExpression struct {
	Left     Expression
	Operator lexer.TokenType
	Right    Expression
}

func (BinaryExpression) node()       {}
func (BinaryExpression) expression() {}

type UnaryExpression struct {
	Operator lexer.TokenType
	Operand  Expression
}

func (UnaryExpression) node()       {}
func (UnaryExpression) expression() {}

type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (CallExpression) node()       {}
func (CallExpression) expression() {}

type IndexExpression struct {
	Object Expression
	Index  Expression
}

func (IndexExpression) node()       {}
func (IndexExpression) expression() {}
