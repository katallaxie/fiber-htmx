package typography

import htmx "github.com/katallaxie/fiber-htmx"

// H1 is a component that displays a h1 text.
func H1(props Props, children ...htmx.Node) htmx.Node {
	return htmx.H1(
		htmx.Merge(
			htmx.ClassNames{
				"text-4xl":  true,
				"font-bold": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// H2 is a component that displays a h2 text.
func H2(props Props, children ...htmx.Node) htmx.Node {
	return htmx.H2(
		htmx.Merge(
			htmx.ClassNames{
				"text-3xl":  true,
				"font-bold": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// H3 is a component that displays a h3 text.
func H3(props Props, children ...htmx.Node) htmx.Node {
	return htmx.H3(
		htmx.Merge(
			htmx.ClassNames{
				"text-2xl":  true,
				"font-bold": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// H4 is a component that displays a h4 text.
func H4(props Props, children ...htmx.Node) htmx.Node {
	return htmx.H4(
		htmx.Merge(
			htmx.ClassNames{
				"text-xl":   true,
				"font-bold": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// H5 is a component that displays a h5 text.
func H5(props Props, children ...htmx.Node) htmx.Node {
	return htmx.H5(
		htmx.Merge(
			htmx.ClassNames{
				"text-lg":   true,
				"font-bold": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// H6 is a component that displays a h6 text.
func H6(props Props, children ...htmx.Node) htmx.Node {
	return htmx.H6(
		htmx.Merge(
			htmx.ClassNames{
				"text-base": true,
				"font-bold": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
