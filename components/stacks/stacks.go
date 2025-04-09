package stacks

import htmx "github.com/katallaxie/fiber-htmx"

// StackProps is a struct that contains the properties of a stack.
type StackProps struct {
	// ClassNames is a map of class names.
	ClassNames htmx.ClassNames
}

// Stack is a function that returns a stack.
func Stack(props StackProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stack": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
