package sourcecode

import (
	"fmt"
	"io"
	"strings"
)

type Parser struct{}

func (p Parser) ParseFile(file interface {
	io.Reader
	Name() string
}) (File, error) {
	source, err := io.ReadAll(file)
	if err != nil {
		return File{}, fmt.Errorf("failed to read source file %q: %w", file.Name(), err)
	}

	parts, err := p.extractParts(string(source))
	if err != nil {
		return File{}, fmt.Errorf("failed to extract parts: %w", err)
	}

	template, err := p.parseTemplate(parts.Template)
	if err != nil {
		return File{}, fmt.Errorf("failed to parse template: %w", err)
	}

	return File{
		Path:     file.Name(),
		Preamble: parts.Preamble,
		Template: template,
	}, nil
}

type fileParts struct {
	Preamble, Template string
}

func (p Parser) extractParts(source string) (fileParts, error) {
	const preambleDelimiter = "---\n"

	// Ditch the first preamble delimiter.
	source = strings.TrimLeft(source, preambleDelimiter)

	// Separate the preamble from the template source.
	parts := strings.Split(source, preambleDelimiter)

	// Validate the parts.
	if n := len(parts); n != 2 {
		return fileParts{}, fmt.Errorf("source file must have exactly 2 parts (preamble and template), but got: %d", n)
	}

	return fileParts{
		Preamble: strings.TrimSpace(parts[0]),
		Template: strings.TrimSpace(parts[1]),
	}, nil
}

func (p Parser) parseTemplate(template string) (Template, error) {
	// Clean up the template for easier parsing.
	template = strings.TrimSpace(template)

	// Separate the markup from the CSS.
	parts := strings.Split(template, "<style>")
	return Template{
		Markup: strings.TrimSpace(parts[0]),
		CSS:    strings.TrimSpace(strings.TrimSuffix(parts[1], "</style>")),
	}, nil
}
