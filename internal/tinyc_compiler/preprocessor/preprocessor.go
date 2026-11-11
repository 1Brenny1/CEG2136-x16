package preprocessor

import (
	"fmt"
	"os"
	"unicode"
)

type Preprocessor struct {
	Source string
	Pos    int
	Length int
}

func New(source string) *Preprocessor {
	return &Preprocessor{
		Source: source,
		Pos:    0,
		Length: len(source),
	}
}

func (p *Preprocessor) Process() (string, error) {

	for p.Pos < p.Length {
		if p.current() != '#' {
			p.advance()
			continue
		}

		start := p.Pos

		for p.Pos < p.Length && !unicode.IsSpace(rune(p.current())) {
			p.advance()
		}

		keyword := p.Source[start+1 : p.Pos]

		switch keyword {
		case "include":
			err := p.parse_include(start)
			if err != nil {
				return "", err
			}
		default:
			return "", fmt.Errorf("Unexpected Preprocessor keyword: %s", keyword)
		}
	}

	return p.Source, nil
}

func (p *Preprocessor) parse_include(cmdStart int) error {
	for p.Pos < p.Length && unicode.IsSpace(rune(p.current())) {
		p.advance()
	}

	if err := p.expect('<'); err != nil {
		return err
	}

	start := p.Pos

	for p.Pos < p.Length && p.current() != '>' {
		p.advance()
	}

	if err := p.expect('>'); err != nil {
		return err
	}

	filename := p.Source[start : p.Pos-1]

	src, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Failed to read source file %s: %s", filename, err.Error())
	}
	p.replace(cmdStart, p.Pos, string(src))

	return nil
}
