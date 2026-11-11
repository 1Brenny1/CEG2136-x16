package ast

type Node interface {
	node()
}

type Expression interface {
	Node
	expression()
}

type Statement interface {
	Node
	statement()
}

type Declaration interface {
	Node
	declaration()
}

type Identifier struct {
	Name string
}

func (Identifier) node()       {}
func (Identifier) expression() {}

type Program struct {
	Declarations []Declaration
}

func (Program) node() {}
