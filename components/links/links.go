package links

import (
	htmx "github.com/katallaxie/fiber-htmx"
)

// LinkProps represents the properties for a link element.
type LinkProps struct {
	Rel        string          // The relationship between the current document and the linked document.
	Href       string          // The URL of the linked document.
	ClassNames map[string]bool // The class names for the link element.
	Active     bool            // Whether the link is active.
}

// Link generates a link element based on the provided properties.
func Link(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Primary(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Secondary(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Accent(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Neutral(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Info(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Warning(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Error(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Underline(props LinkProps, children ...htmx.Node) htmx.Node {
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
func Button(props LinkProps, children ...htmx.Node) htmx.Node {
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
