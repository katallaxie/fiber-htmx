package html

import (
	htmx "github.com/katallaxie/fiber-htmx"
)

type Parser interface {
	// FromBytes parses the given byte slice and returns a Node.
	FromBytes([]byte) (htmx.Node, error)
}

var _ Parser = (*parser)(nil)

type parser struct{}

// NewParser returns a new instance of the parser.
func NewParser() *parser {
	return &parser{}
}

// FromBytes parses the given byte slice and returns a Node.
func (p *parser) FromBytes(in []byte) (htmx.Node, error) {
	return nil, nil
}
