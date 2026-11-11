package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"x16/internal/emulator"
)

func main() {
	memDump := flag.String("dump", "", "Optional path to write a memory dump when the program exits")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file-path>\n\nOptions:", os.Args)
		flag.PrintDefaults()
	}

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Error: <file-path> is required")
		flag.Usage()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err.Error())
		os.Exit(1)
	}

	cpu, err := emulator.Emulate(bytes)
	if err != nil {
		fmt.Printf("An error occurred when trying to emulate the rom: %s\n", err.Error())
		os.Exit(1)
	}

	if *memDump != "" {
		bytes := make([]byte, len(cpu.Memory)*2)
		for i, word := range cpu.Memory {
			binary.BigEndian.PutUint16(bytes[i*2:], word)
		}

		err := os.WriteFile(*memDump, bytes, 0644)
		if err != nil {
			fmt.Printf("An error occurred when trying to write to the memory dump: %s\n", err.Error())
			os.Exit(1)
		}
	}
}
