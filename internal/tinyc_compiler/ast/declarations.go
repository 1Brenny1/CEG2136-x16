package ast

type VariableDeclaration struct {
	Name  string
	Value Expression
}

func (VariableDeclaration) node()        {}
func (VariableDeclaration) declaration() {}

type FunctionDeclaration struct {
	Name       string
	Parameters []string // Array For Future Use
	Body       *BlockStatement
}

func (FunctionDeclaration) node()        {}
func (FunctionDeclaration) declaration() {}

type DeclarationStatement struct {
	Declaration Declaration
}

func (DeclarationStatement) node()      {}
func (DeclarationStatement) statement() {}
