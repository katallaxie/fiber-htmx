package heros

import htmx "github.com/katallaxie/fiber-htmx"

// Props is the props for the Hero component.
type Props struct {
	htmx.ClassNames
}

// Hero is a component that renders a hero section.
func Hero(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"hero": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Content is a component that renders the content of the hero section.
func Content(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"hero-content": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Overlay is a component that renders an overlay for the hero section.
func Overlay(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"hero-overlay": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
