package links

import (
	htmx "github.com/katallaxie/fiber-htmx"
)

// Props represents the properties for a link element.
type Props struct {
	// Rel is the relationship between the current document and the linked document.
	Rel string
	// Href is the URL of the linked document.
	Href string
	// Active indicates whether the link is active or not.
	Active bool

	htmx.ClassNames
}

// Link generates a link element based on the provided properties.
func Link(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":   true,
				"active": props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Primary generates a primary link element based on the provided properties.
func Primary(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":         true,
				"link-primary": true,
				"active":       props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Secondary generates a secondary link element based on the provided properties.
func Secondary(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":           true,
				"link-secondary": true,
				"active":         props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Accent generates an accent link element based on the provided properties.
func Accent(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":        true,
				"link-accent": true,
				"active":      props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Neutral generates a neutral link element based on the provided properties.
func Neutral(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":         true,
				"link-neutral": true,
				"active":       props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Info generates an info link element based on the provided properties.
func Info(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":      true,
				"link-info": true,
				"active":    props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Warning generates a warning link element based on the provided properties.
func Warning(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":         true,
				"link-warning": true,
				"active":       props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Error generates an error link element based on the provided properties.
func Error(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":       true,
				"link-error": true,
				"active":     props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Underline generates an underline link element based on the provided properties.
func Underline(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"link":       true,
				"link-hover": true,
				"active":     props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}

// Button generate a link that looks like a button.
func Button(props Props, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"btn":    true,
				"active": props.Active,
			},
			props.ClassNames,
		),
		htmx.Href(props.Href),
		htmx.Rel(props.Rel),
		htmx.Group(children...),
	)
}
