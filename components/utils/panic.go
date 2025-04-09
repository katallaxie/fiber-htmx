package utils

import htmx "github.com/katallaxie/fiber-htmx"

// Panic is a function that panics with the given error.
func Panic(err error) htmx.Node {
	panic(err)
}
