package htmx

// Filter loops and filters the content.
func Filter(f func(n Node) bool, children ...Node) []Node {
	var nodes []Node
	for i := 0; i < len(children); i++ {
		if f(children[i]) {
			nodes = append(nodes, children[i])
		}
	}

	return nodes
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
