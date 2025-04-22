package forms

import htmx "github.com/katallaxie/fiber-htmx"

// LabelProps is a struct that contains the properties of the Label component.
type LabelProps struct {
	htmx.ClassNames
}

// Label is a component that renders a label element.
func Label(props LabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"input": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// LabelText is a component that renders a label element with text.
func LabelText(props LabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"label": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FloatingLabel is a component that renders a label element with a floating effect.
func FloatingLabel(props LabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"floating-label": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// LabelSelect is a component that renders a label element for a select input.
func LabelSelect(props LabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"select": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// LabelInput is a component that renders a label element for an input element.
func LabelInput(props LabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"input": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
