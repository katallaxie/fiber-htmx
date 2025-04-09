package joins

import htmx "github.com/katallaxie/fiber-htmx"

// JoinProps is a struct that contains the properties of a join.
type JoinProps struct {
	// ClassNames is a map of class names.
	ClassNames htmx.ClassNames
}

// Join is a function that returns a join.
func Join(props JoinProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// JoinVertical is a function that returns a vertical join.
func JoinVertical(props JoinProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join":          true,
				"join-vertical": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// JoinItem is a function that returns a join item.
func JoinItem(props JoinProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join-item": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
