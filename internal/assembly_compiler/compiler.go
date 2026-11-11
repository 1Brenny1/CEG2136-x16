package assembly_compiler

func Compile(source string) ([]byte, error) {
	lexer := Lexer{
		I:      0,
		N:      len(source),
		Source: source,
		Tokens: make([]Token, 0),
	}

	tokens, err := lexer.Tokenize()
	if err != nil {
		return nil, err
	}

	// for _, tok := range tokens {
	// 	fmt.Printf("%v ", tok)
	// }

	// fmt.Println()

	parser := Parser{
		I:      0,
		N:      len(tokens),
		Tokens: tokens,
	}

	nodes, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	// for _, node := range nodes {
	// 	fmt.Printf("%v ", node)
	// }
	// fmt.Println()

	assembler := Assembler{
		Nodes:   nodes,
		Symbols: make(map[string]uint16),
	}

	bytes, err := assembler.Assemble()
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
