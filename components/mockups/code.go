package mockups

import htmx "github.com/katallaxie/fiber-htmx"

// CodeProps is a struct that contains the properties of the code component.
type CodeProps struct {
	// Language is the language of the code.
	Language string

	htmx.ClassNames
}

// Code is a function that returns the code component.
func Code(props CodeProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"mockup-code": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// WindowProps is a struct that contains the properties of the window component.
type WindowProps struct {
	htmx.ClassNames
}

// Window is a function that returns the window component.
func Window(props WindowProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"mockup-window":   true,
				"border":          true,
				"border-base-300": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
