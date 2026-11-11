package tinyc_compiler

import (
	"encoding/json"
	"fmt"
	"x16/internal/tinyc_compiler/assembler"
	"x16/internal/tinyc_compiler/lexer"
	"x16/internal/tinyc_compiler/parser"
	"x16/internal/tinyc_compiler/preprocessor"
)

func Compile(source string) (string, error) {

	preprocessor := preprocessor.New(source)

	src, err := preprocessor.Process()
	if err != nil {
		return "", err
	}

	lexer := lexer.New(src)

	tokens, err := lexer.Tokenize()
	if err != nil {
		return "", err
	}

	fmt.Println("Lexer Output:")

	for _, tok := range tokens {
		fmt.Printf("%v ", tok)
	}
	fmt.Println()

	parser := parser.New(tokens)

	program, err := parser.Parse()
	if err != nil {
		return "", err
	}

	fmt.Println("Parser Output:")

	bytes, err := json.Marshal(program)
	if err != nil {
		return "", nil
	}

	fmt.Println(string(bytes))

	assembler := assembler.New(*program)

	assembly, err := assembler.Assemble()
	if err != nil {
		return "", err
	}

	return assembly, nil
}
