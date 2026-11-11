package ast

type IntegerLiteral struct {
	Value uint16
}

func (IntegerLiteral) node()       {}
func (IntegerLiteral) expression() {}

type CharacterLiteral struct {
	Value byte
}

func (CharacterLiteral) node()       {}
func (CharacterLiteral) expression() {}

type StringLiteral struct {
	Value string
}

func (StringLiteral) node()       {}
func (StringLiteral) expression() {}
