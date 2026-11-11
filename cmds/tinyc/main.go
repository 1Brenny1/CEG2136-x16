package main

import (
	"flag"
	"fmt"
	"os"
	"x16/internal/tinyc_compiler"
)

func main() {
	output := flag.String("output", "out.s", "Output assembly path")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <source-path>\n\nOptions:\n", os.Args)
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Error: <source-path> is required")
		flag.Usage()
		os.Exit(1)
	}

	sourcePath := flag.Arg(0)

	bytes, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err.Error())
		os.Exit(1)
	}

	asm, err := tinyc_compiler.Compile(string(bytes))
	if err != nil {
		fmt.Printf("Failed to compile file: %s\n", err.Error())
		os.Exit(1)
	}

	f, err := os.OpenFile(*output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to write assembly to %s: %s", *output, err.Error())
		os.Exit(1)
	}

	f.WriteString(asm)

	f.Close()
}
