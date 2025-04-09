package cards

import htmx "github.com/katallaxie/fiber-htmx"

// CardProps contains the properties for the card component.
type CardProps struct {
	ClassNames htmx.ClassNames
}

// Card is a component for the htmx card extension.
func Card(p CardProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"card":        true,
				"bg-base-100": true,
				"shadow-xl":   true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CardBordered is a component for the htmx card extension.
func CardBordered(p CardProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"card":          true,
				"card-bordered": true,
				"bg-base-100":   true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// BodyProps contains the properties for the card body component.
type BodyProps struct {
	ClassNames htmx.ClassNames
}

// Body is a component for the htmx card extension.
func Body(p BodyProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"card-body": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ActionsProps contains the properties for the card actions component.
type ActionsProps struct {
	ClassNames htmx.ClassNames
}

// Actions is a component for the htmx card extension.
func Actions(p ActionsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"card-actions": true,
				"justify-end":  true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TitleProps contains the properties for the card title component.
type TitleProps struct {
	ClassNames htmx.ClassNames
}

// Title is a component for the htmx card extension.
func Title(p TitleProps, children ...htmx.Node) htmx.Node {
	return htmx.H2(
		htmx.Merge(
			htmx.ClassNames{
				"card-title": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
