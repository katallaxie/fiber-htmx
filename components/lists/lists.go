package lists

import htmx "github.com/katallaxie/fiber-htmx"

// ListProps is a struct that contains the properties of the List component.
type ListProps struct {
	htmx.ClassNames
}

// List is a component that renders a list of items.
func List(props ListProps, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"list": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ListHeader is a component that renders a header for a list.
type ListHeaderProps struct {
	htmx.ClassNames
}

// ListHeader is a component that renders a header for a list.
func ListHeader(props ListHeaderProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"p-4":           true,
				"pb-2":          true,
				"text-xs":       true,
				"opacity-60":    true,
				"tracking-wide": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ListRow is a component that renders a row in a list.
type ListRowProps struct {
	htmx.ClassNames
}

// ListRow is a component that renders a row in a list.
func ListRow(props ListRowProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"list-row": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
