package stacks

import htmx "github.com/katallaxie/fiber-htmx"

// StackProps is a struct that contains the properties of a stack.
type StackProps struct {
	htmx.ClassNames
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

// StackTop is a function that returns a stack with the top class.
func StackTop(props StackProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stack":     true,
				"stack-top": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// StackStart is a function that returns a stack with the start class.
func StackStart(props StackProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stack":       true,
				"stack-start": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// StackEnd is a function that returns a stack with the end class.
func StackEnd(props StackProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stack":     true,
				"stack-end": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
