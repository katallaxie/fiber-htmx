package breadcrumbs

import htmx "github.com/katallaxie/fiber-htmx"

// Props represents the properties for a breadcrumb element.
type Props struct {
	htmx.ClassNames
}

// BreadCrumbs generates a breadcrumb element based on the provided properties.
func Breadcrumbs(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"breadcrumbs": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Items generates a list of breadcrumb items.
func Items(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Item generates a single breadcrumb item.
func Item(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
