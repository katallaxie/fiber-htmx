package dividers

import htmx "github.com/katallaxie/fiber-htmx"

// DividerProps is a struct that contains the props of a divider.
type DividerProps struct {
	ClassNames htmx.ClassNames
}

// Divider is a struct that contains the props of a divider.
func Divider(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerNeutral is a struct that contains the props of a neutral divider.
func DividerNeutral(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-neutral": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerPrimary is a struct that contains the props of a primary divider.
func DividerPrimary(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-primary": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerSecondary is a struct that contains the props of a secondary divider.
func DividerSecondary(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":           true,
				"divider-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerSuccess is a struct that contains the props of a success divider.
func DividerSuccess(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-success": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerWarning is a struct that contains the props of a warning divider.
func DividerWarning(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-warning": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerInfo is a struct that contains the props of an info divider.
func DividerInfo(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":      true,
				"divider-info": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DividerError is a struct that contains the props of an error divider.
func DividerError(p DividerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":      true,
				"divider-info": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
