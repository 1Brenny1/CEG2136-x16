package main

import (
	"flag"
	"fmt"
	"os"
	"x16/internal/assembly_compiler"
)

func main() {
	output := flag.String("output", "out.x16", "Output binary path")

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

	romPath := flag.Arg(0)

	bytes, err := os.ReadFile(romPath)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err.Error())
		os.Exit(1)
	}

	source := string(bytes)
	bytes, err = assembly_compiler.Compile(source)
	if err != nil {
		fmt.Printf("Failed to compile file: %s\n", err.Error())
		os.Exit(1)
	}

	f, err := os.OpenFile(*output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to write output: %s\n", err.Error())
		os.Exit(1)
	}
	f.Write(bytes)

	f.Close()
}
