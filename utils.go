package htmx

import (
	"encoding/json"

	"github.com/katallaxie/pkg/slices"
)

// Filter loops and filters the content.
func Filter(f func(n Node) bool, children ...Node) []Node {
	return slices.Filter(func(n Node) bool { return f(n) }, children...)
}

// Map is using a map to transform into htmx nodes.
func Map[T1 comparable, T2 any](m map[T1]T2, f func(T1, T2) Node) Nodes {
	nodes := make([]Node, 0, len(m))

	for k, v := range m {
		nodes = append(nodes, f(k, v))
	}

	return nodes
}

// Reduce reduces the content to a single node.
func Reduce(f func(prev Node, next Node) Node, children ...Node) Node {
	if len(children) == 0 {
		return nil
	}

	node := children[0]
	for _, n := range children[1:] {
		node = f(node, n)
	}

	return node
}

// ForEach loops over the content.
func ForEach[S ~[]E, E comparable](s S, f func(E, int) Node) Nodes {
	nodes := make([]Node, 0, len(s))

	for i, e := range s {
		nodes = append(nodes, f(e, i))
	}

	return nodes
}

// Merge returns a new ClassNames object that is the result of merging the provided ClassNames objects.
func Merge(classNames ...ClassNames) ClassNames {
	merged := ClassNames{}

	for i := 0; i < len(classNames); i++ {
		for k, v := range classNames[i] {
			merged[k] = v
		}
	}

	return merged
}

// JsonSerializeOrEmpty returns a JSON serialized string of the provided data or an empty string if the serialization fails.
func JsonSerializeOrEmpty(data any) string {
	if data == nil {
		return ""
	}

	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(serialized)
}
