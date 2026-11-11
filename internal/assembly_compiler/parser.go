package assembly_compiler

import "fmt"

type Node interface {
	isNode()
}

type LabelNode struct {
	Name string
}

func (LabelNode) isNode() {}

type InstructionNode struct {
	Opcode   string
	Operand  string
	Indirect bool
}

func (InstructionNode) isNode() {}

type DirectiveNode struct {
	Directive string
	Value     string
}

func (DirectiveNode) isNode() {}

type Parser struct {
	Tokens []Token
	I      int
	N      int
}

func (p *Parser) peek(offset int) Token {
	if p.I+offset < p.N {
		return p.Tokens[p.I+offset]
	}

	return Token{}
}

func (p *Parser) consume() Token {
	tok := p.Tokens[p.I]
	p.I++
	return tok
}

func (p *Parser) Parse() ([]Node, error) {
	var nodes []Node

loop:
	for p.I < p.N {
		tok := p.peek(0)

		switch tok.Type {
		case TOKEN_DIRECTIVE:
			nodes = append(nodes, p.parseDirective())
			continue loop

		case TOKEN_KEYWORD:
			nodes = append(nodes, p.parseInstruction())
			continue loop

		case TOKEN_IDENTIFIER:
			if p.peek(1).Type == TOKEN_COLON {
				nodes = append(nodes, p.parseLabel())
				continue loop
			}
		}

		return nil, fmt.Errorf("Unexpected token at %d: %v", p.I, tok)
	}

	return nodes, nil
}

func (p *Parser) parseLabel() LabelNode {
	name := p.consume().Value // Identifier
	p.consume()               // Colon

	return LabelNode{
		Name: name,
	}
}

func (p *Parser) parseInstruction() InstructionNode {
	opcode := p.consume().Value

	node := InstructionNode{
		Opcode: opcode,
	}

	if _, ok := memoryInstructions[opcode]; ok &&
		(p.peek(0).Type == TOKEN_IDENTIFIER || p.peek(0).Type == TOKEN_NUMBER) {
		node.Operand = p.consume().Value

		if p.peek(0).Type == TOKEN_INDIRECT {
			node.Indirect = true
			p.consume()
		}
	}

	return node
}

func (p *Parser) parseDirective() DirectiveNode {
	directive := p.consume().Value

	node := DirectiveNode{
		Directive: directive,
	}

	if p.peek(0).Type == TOKEN_IDENTIFIER || p.peek(0).Type == TOKEN_NUMBER {
		node.Value = p.consume().Value
	}

	return node
}
