package ast

type BlockStatement struct {
	Statements []Statement
}

func (BlockStatement) node()      {}
func (BlockStatement) statement() {}

type ExpressionStatement struct {
	Expression Expression
}

func (ExpressionStatement) node()      {}
func (ExpressionStatement) statement() {}

// TODO: Come up with better names
type AssignmentStatement struct {
	Variable Expression
	Value    Expression
}

func (AssignmentStatement) node()      {}
func (AssignmentStatement) statement() {}

type IfStatement struct {
	Condition Expression
	Then      Statement
	Else      Statement
}

func (IfStatement) node()      {}
func (IfStatement) statement() {}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func (WhileStatement) node()      {}
func (WhileStatement) statement() {}

type ReturnStatement struct {
	Value Expression
}

func (ReturnStatement) node()      {}
func (ReturnStatement) statement() {}

type BreakStatement struct{}

func (BreakStatement) node()      {}
func (BreakStatement) statement() {}

type ContinueStatement struct{}

func (ContinueStatement) node()      {}
func (ContinueStatement) statement() {}
