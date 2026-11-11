package preprocessor

import "fmt"

func (p *Preprocessor) advance() {
	p.Pos++
}

func (p *Preprocessor) current() byte {
	return p.Source[p.Pos]
}

func (p *Preprocessor) expect(char byte) error {
	if p.current() == char {
		p.advance()
		return nil
	}

	return fmt.Errorf("Expected %c but got %c", char, p.current())
}

func (p *Preprocessor) replace(startPos int, endPos int, content string) {
	start := p.Source[:startPos]
	end := p.Source[endPos:]
	p.Source = start + content + end
	p.Pos = startPos
	p.Length = len(p.Source)
}
