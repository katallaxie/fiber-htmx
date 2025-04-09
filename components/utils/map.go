package utils

import htmx "github.com/katallaxie/fiber-htmx"

// Transformer is a function that transforms a node.
type Transformer[T any] func(el T) htmx.Node

// Map is a function that maps a transformer to a list of htmx.Node.
func Map[T any](transformer Transformer[T], elems ...T) htmx.Node {
	nodes := make([]htmx.Node, 0, len(elems))

	for _, n := range elems {
		nodes = append(nodes, transformer(n))
	}

	return htmx.Group(nodes...)
}
