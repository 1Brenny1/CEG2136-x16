package assembler

import (
	"x16/internal/tinyc_compiler/assembler/snippets"
	"x16/internal/tinyc_compiler/ast"
)

type Assembler struct {
	Program  ast.Program
	Assembly string
}

func New(program ast.Program) *Assembler {
	return &Assembler{
		Program:  program,
		Assembly: "",
	}
}

func (a *Assembler) Assemble() (string, error) {

	return snippets.EntryPoint + a.Assembly, nil
}
