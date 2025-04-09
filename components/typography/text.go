package typography

import htmx "github.com/katallaxie/fiber-htmx"

// Props contains the properties for the text component.
type Props struct {
	// ClassNames contains the class names for the text component.
	ClassNames htmx.ClassNames
}

// Error is a component that displays an error text.
func Error(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-error": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Info is a component that displays an info text.
func Info(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-info": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Success is a component that displays a success text.
func Success(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-success": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Warning is a component that displays a warning text.
func Warning(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-warning": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Primary is a component that displays a primary text.
func Primary(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-primary": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Secondary is a component that displays a secondary text.
func Secondary(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-secondary": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Accent is a component that displays an accent text.
func Accent(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-accent": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Neutral is a component that displays a neutral text.
func Neutral(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-neutral": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Base100 is a component that displays a base100 text.
func Base100(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-base-100": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Base200 is a component that displays a base200 text.
func Base200(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-base-200": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Base300 is a component that displays a base300 text.
func Base300(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-base-300": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
