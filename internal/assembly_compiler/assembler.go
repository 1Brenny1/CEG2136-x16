package assembly_compiler

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type Assembler struct {
	Nodes   []Node
	Symbols map[string]uint16
	Memory  [4096]uint16
}

func (a *Assembler) Assemble() ([]byte, error) {
	err := a.BuildSymbols()
	if err != nil {
		return nil, err
	}

	err = a.Generate()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, len(a.Memory)*2)
	for i, word := range a.Memory {
		binary.BigEndian.PutUint16(bytes[i*2:], word)
	}

	return bytes, nil
}

func (a *Assembler) BuildSymbols() error {
	address := uint16(0)

	for _, node := range a.Nodes {
		switch n := node.(type) {

		case DirectiveNode:
			switch n.Directive {
			case "ORG":
				val, err := strconv.ParseInt(n.Value, 10, 16)
				if err != nil {
					return err
				}
				address = uint16(val)

			case "DEC", "HEX", "ADR":
				address++
			}

		case InstructionNode:
			address++

		case LabelNode:
			a.Symbols[n.Name] = address
		}
	}

	return nil
}

func (a *Assembler) Generate() error {
	address := uint16(0)

	for _, node := range a.Nodes {
		switch n := node.(type) {
		case DirectiveNode:
			switch n.Directive {
			case "ORG":
				val, err := strconv.ParseInt(n.Value, 10, 16)
				if err != nil {
					return err
				}
				address = uint16(val)

			case "DEC":
				val, err := strconv.ParseInt(n.Value, 10, 16)
				if err != nil {
					return err
				}
				a.Memory[address] = uint16(val)
				address++

			case "HEX":
				val, err := strconv.ParseUint(n.Value, 16, 16)
				if err != nil {
					return err
				}
				a.Memory[address] = uint16(val)
				address++

			case "ADR":
				addr, ok := a.Symbols[n.Value]
				if !ok {
					return fmt.Errorf("Undefined symbol: %s", n.Value)
				}
				a.Memory[address] = addr
				address++
			}

		case InstructionNode:
			if opcode, ok := registerInstructions[n.Opcode]; ok {
				a.Memory[address] = opcode
				address++
				break
			}

			if opcode, ok := ioInstructions[n.Opcode]; ok {
				a.Memory[address] = opcode
				address++
				break
			}

			opcode, ok := memoryInstructions[n.Opcode]
			if !ok {
				return fmt.Errorf("Unknown Instruction: %s", n.Opcode)
			}

			addr, ok := a.Symbols[n.Operand]
			if !ok {
				return fmt.Errorf("Undefined symbol: %s", n.Operand)
			}

			word := (opcode << 12) | addr

			if n.Indirect {
				word |= 0x8000
			}

			a.Memory[address] = word
			address++
		}
	}

	return nil
}
